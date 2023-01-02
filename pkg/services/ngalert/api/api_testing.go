package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/services/datasources"
	"github.com/grafana/grafana/pkg/services/featuremgmt"
	apimodels "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
	"github.com/grafana/grafana/pkg/services/ngalert/backtesting"
	"github.com/grafana/grafana/pkg/services/ngalert/eval"
	ngmodels "github.com/grafana/grafana/pkg/services/ngalert/models"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/util"
)

type TestingApiSrv struct {
	*AlertingProxy
	DatasourceCache datasources.CacheService
	log             log.Logger
	accessControl   accesscontrol.AccessControl
	evaluator       eval.EvaluatorFactory
	cfg             *setting.UnifiedAlertingSettings
	backtesting     *backtesting.Engine
	featureManager  featuremgmt.FeatureToggles
}

func (srv TestingApiSrv) RouteTestGrafanaRuleConfig(c *models.ReqContext, body apimodels.TestRulePayload) response.Response {
	if body.Type() != apimodels.GrafanaBackend || body.GrafanaManagedCondition == nil {
		return errorToResponse(backendTypeDoesNotMatchPayloadTypeError(apimodels.GrafanaBackend, body.Type().String()))
	}

	if !authorizeDatasourceAccessForRule(&ngmodels.AlertRule{Data: body.GrafanaManagedCondition.Data}, func(evaluator accesscontrol.Evaluator) bool {
		return accesscontrol.HasAccess(srv.accessControl, c)(accesscontrol.ReqSignedIn, evaluator)
	}) {
		return errorToResponse(fmt.Errorf("%w to query one or many data sources used by the rule", ErrAuthorization))
	}

	evalCond := ngmodels.Condition{
		Condition: body.GrafanaManagedCondition.Condition,
		Data:      body.GrafanaManagedCondition.Data,
	}
	ctx := eval.Context(c.Req.Context(), c.SignedInUser)

	conditionEval, err := srv.evaluator.Create(ctx, evalCond)
	if err != nil {
		return ErrResp(http.StatusBadRequest, err, "invalid condition")
	}

	now := body.GrafanaManagedCondition.Now
	if now.IsZero() {
		now = timeNow()
	}

	evalResults, err := conditionEval.Evaluate(c.Req.Context(), now)
	if err != nil {
		return ErrResp(500, err, "Failed to evaluate the rule")
	}

	frame := evalResults.AsDataFrame()
	return response.JSONStreaming(http.StatusOK, util.DynMap{
		"instances": []*data.Frame{&frame},
	})
}

func (srv TestingApiSrv) RouteTestRuleConfig(c *models.ReqContext, body apimodels.TestRulePayload, datasourceUID string) response.Response {
	if body.Type() != apimodels.LoTexRulerBackend {
		return errorToResponse(backendTypeDoesNotMatchPayloadTypeError(apimodels.LoTexRulerBackend, body.Type().String()))
	}
	ds, err := getDatasourceByUID(c, srv.DatasourceCache, apimodels.LoTexRulerBackend)
	if err != nil {
		return errorToResponse(err)
	}
	var path string
	switch ds.Type {
	case "loki":
		path = "loki/api/v1/query"
	case "prometheus":
		path = "api/v1/query"
	default:
		// this should not happen because getDatasourceByUID would not return the data source
		return errorToResponse(unexpectedDatasourceTypeError(ds.Type, "loki, prometheus"))
	}

	t := timeNow()
	queryURL, err := url.Parse(path)
	if err != nil {
		return ErrResp(http.StatusInternalServerError, err, "failed to parse url")
	}
	params := queryURL.Query()
	params.Set("query", body.Expr)
	params.Set("time", strconv.FormatInt(t.Unix(), 10))
	queryURL.RawQuery = params.Encode()
	return srv.withReq(
		c,
		http.MethodGet,
		queryURL,
		nil,
		instantQueryResultsExtractor,
		nil,
	)
}

func (srv TestingApiSrv) RouteEvalQueries(c *models.ReqContext, cmd apimodels.EvalQueriesPayload) response.Response {
	if !authorizeDatasourceAccessForRule(&ngmodels.AlertRule{Data: cmd.Data}, func(evaluator accesscontrol.Evaluator) bool {
		return accesscontrol.HasAccess(srv.accessControl, c)(accesscontrol.ReqSignedIn, evaluator)
	}) {
		return ErrResp(http.StatusUnauthorized, fmt.Errorf("%w to query one or many data sources used by the rule", ErrAuthorization), "")
	}

	cond := ngmodels.Condition{
		Condition: "",
		Data:      cmd.Data,
	}
	if len(cmd.Data) > 0 {
		cond.Condition = cmd.Data[0].RefID
	}
	evaluator, err := srv.evaluator.Create(eval.Context(c.Req.Context(), c.SignedInUser), cond)

	if err != nil {
		return ErrResp(http.StatusBadRequest, err, "Failed to build evaluator for queries and expressions")
	}

	now := cmd.Now
	if now.IsZero() {
		now = timeNow()
	}

	evalResults, err := evaluator.EvaluateRaw(c.Req.Context(), now)

	if err != nil {
		return ErrResp(http.StatusInternalServerError, err, "Failed to evaluate queries and expressions")
	}

	return response.JSONStreaming(http.StatusOK, evalResults)
}

func (srv TestingApiSrv) BacktestAlertRule(c *models.ReqContext, cmd apimodels.BacktestConfig) response.Response {
	if !srv.featureManager.IsEnabled(featuremgmt.FlagAlertingBacktesting) {
		return ErrResp(http.StatusNotFound, nil, "Backgtesting API is not enabled")
	}

	if cmd.From.After(cmd.To) {
		return ErrResp(400, nil, "From cannot be greater than To")
	}

	noDataState, err := ngmodels.NoDataStateFromString(string(cmd.NoDataState))

	if err != nil {
		return ErrResp(400, err, "")
	}
	forInterval := time.Duration(cmd.For)
	if forInterval < 0 {
		return ErrResp(400, nil, "Bad For interval")
	}

	intervalSeconds, err := validateInterval(srv.cfg, time.Duration(cmd.Interval))
	if err != nil {
		return ErrResp(400, err, "")
	}

	if !authorizeDatasourceAccessForRule(&ngmodels.AlertRule{Data: cmd.Data}, func(evaluator accesscontrol.Evaluator) bool {
		return accesscontrol.HasAccess(srv.accessControl, c)(accesscontrol.ReqSignedIn, evaluator)
	}) {
		return errorToResponse(fmt.Errorf("%w to query one or many data sources used by the rule", ErrAuthorization))
	}

	rule := &ngmodels.AlertRule{
		// ID:             0,
		// Updated:        time.Time{},
		// Version:        0,
		// NamespaceUID:   "",
		// DashboardUID:   nil,
		// PanelID:        nil,
		// RuleGroup:      "",
		// RuleGroupIndex: 0,
		// ExecErrState:   "",
		Title: cmd.Title,
		// prefix backtesting- is to distinguish between executions of regular rule and backtesting in logs (like expression engine, evaluator, state manager etc)
		UID:             "backtesting-" + util.GenerateShortUID(),
		OrgID:           c.OrgID,
		Condition:       cmd.Condition,
		Data:            cmd.Data,
		IntervalSeconds: intervalSeconds,
		NoDataState:     noDataState,
		For:             forInterval,
		Annotations:     cmd.Annotations,
		Labels:          cmd.Labels,
	}

	result, err := srv.backtesting.Test(c.Req.Context(), c.SignedInUser, rule, cmd.From, cmd.To)
	if err != nil {
		if errors.Is(err, backtesting.ErrInvalidInputData) {
			return ErrResp(400, err, "Failed to evaluate")
		}
		return ErrResp(500, err, "Failed to evaluate")
	}

	body, err := data.FrameToJSON(result, data.IncludeAll)
	if err != nil {
		return ErrResp(500, err, "Failed to convert frame to JSON")
	}
	return response.JSON(http.StatusOK, body)
}
