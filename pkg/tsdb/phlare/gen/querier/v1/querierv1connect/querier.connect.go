// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: querier/v1/querier.proto

package querierv1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect_go "github.com/bufbuild/connect-go"

	v1 "github.com/grafana/grafana/pkg/tsdb/phlare/gen/querier/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// QuerierServiceName is the fully-qualified name of the QuerierService service.
	QuerierServiceName = "querier.v1.QuerierService"
)

// QuerierServiceClient is a client for the querier.v1.QuerierService service.
type QuerierServiceClient interface {
	ProfileTypes(context.Context, *connect_go.Request[v1.ProfileTypesRequest]) (*connect_go.Response[v1.ProfileTypesResponse], error)
	LabelValues(context.Context, *connect_go.Request[v1.LabelValuesRequest]) (*connect_go.Response[v1.LabelValuesResponse], error)
	LabelNames(context.Context, *connect_go.Request[v1.LabelNamesRequest]) (*connect_go.Response[v1.LabelNamesResponse], error)
	Series(context.Context, *connect_go.Request[v1.SeriesRequest]) (*connect_go.Response[v1.SeriesResponse], error)
	SelectMergeStacktraces(context.Context, *connect_go.Request[v1.SelectMergeStacktracesRequest]) (*connect_go.Response[v1.SelectMergeStacktracesResponse], error)
	SelectSeries(context.Context, *connect_go.Request[v1.SelectSeriesRequest]) (*connect_go.Response[v1.SelectSeriesResponse], error)
}

// NewQuerierServiceClient constructs a client for the querier.v1.QuerierService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewQuerierServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) QuerierServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &querierServiceClient{
		profileTypes: connect_go.NewClient[v1.ProfileTypesRequest, v1.ProfileTypesResponse](
			httpClient,
			baseURL+"/querier.v1.QuerierService/ProfileTypes",
			opts...,
		),
		labelValues: connect_go.NewClient[v1.LabelValuesRequest, v1.LabelValuesResponse](
			httpClient,
			baseURL+"/querier.v1.QuerierService/LabelValues",
			opts...,
		),
		labelNames: connect_go.NewClient[v1.LabelNamesRequest, v1.LabelNamesResponse](
			httpClient,
			baseURL+"/querier.v1.QuerierService/LabelNames",
			opts...,
		),
		series: connect_go.NewClient[v1.SeriesRequest, v1.SeriesResponse](
			httpClient,
			baseURL+"/querier.v1.QuerierService/Series",
			opts...,
		),
		selectMergeStacktraces: connect_go.NewClient[v1.SelectMergeStacktracesRequest, v1.SelectMergeStacktracesResponse](
			httpClient,
			baseURL+"/querier.v1.QuerierService/SelectMergeStacktraces",
			opts...,
		),
		selectSeries: connect_go.NewClient[v1.SelectSeriesRequest, v1.SelectSeriesResponse](
			httpClient,
			baseURL+"/querier.v1.QuerierService/SelectSeries",
			opts...,
		),
	}
}

// querierServiceClient implements QuerierServiceClient.
type querierServiceClient struct {
	profileTypes           *connect_go.Client[v1.ProfileTypesRequest, v1.ProfileTypesResponse]
	labelValues            *connect_go.Client[v1.LabelValuesRequest, v1.LabelValuesResponse]
	labelNames             *connect_go.Client[v1.LabelNamesRequest, v1.LabelNamesResponse]
	series                 *connect_go.Client[v1.SeriesRequest, v1.SeriesResponse]
	selectMergeStacktraces *connect_go.Client[v1.SelectMergeStacktracesRequest, v1.SelectMergeStacktracesResponse]
	selectSeries           *connect_go.Client[v1.SelectSeriesRequest, v1.SelectSeriesResponse]
}

// ProfileTypes calls querier.v1.QuerierService.ProfileTypes.
func (c *querierServiceClient) ProfileTypes(ctx context.Context, req *connect_go.Request[v1.ProfileTypesRequest]) (*connect_go.Response[v1.ProfileTypesResponse], error) {
	return c.profileTypes.CallUnary(ctx, req)
}

// LabelValues calls querier.v1.QuerierService.LabelValues.
func (c *querierServiceClient) LabelValues(ctx context.Context, req *connect_go.Request[v1.LabelValuesRequest]) (*connect_go.Response[v1.LabelValuesResponse], error) {
	return c.labelValues.CallUnary(ctx, req)
}

// LabelNames calls querier.v1.QuerierService.LabelNames.
func (c *querierServiceClient) LabelNames(ctx context.Context, req *connect_go.Request[v1.LabelNamesRequest]) (*connect_go.Response[v1.LabelNamesResponse], error) {
	return c.labelNames.CallUnary(ctx, req)
}

// Series calls querier.v1.QuerierService.Series.
func (c *querierServiceClient) Series(ctx context.Context, req *connect_go.Request[v1.SeriesRequest]) (*connect_go.Response[v1.SeriesResponse], error) {
	return c.series.CallUnary(ctx, req)
}

// SelectMergeStacktraces calls querier.v1.QuerierService.SelectMergeStacktraces.
func (c *querierServiceClient) SelectMergeStacktraces(ctx context.Context, req *connect_go.Request[v1.SelectMergeStacktracesRequest]) (*connect_go.Response[v1.SelectMergeStacktracesResponse], error) {
	return c.selectMergeStacktraces.CallUnary(ctx, req)
}

// SelectSeries calls querier.v1.QuerierService.SelectSeries.
func (c *querierServiceClient) SelectSeries(ctx context.Context, req *connect_go.Request[v1.SelectSeriesRequest]) (*connect_go.Response[v1.SelectSeriesResponse], error) {
	return c.selectSeries.CallUnary(ctx, req)
}

// QuerierServiceHandler is an implementation of the querier.v1.QuerierService service.
type QuerierServiceHandler interface {
	ProfileTypes(context.Context, *connect_go.Request[v1.ProfileTypesRequest]) (*connect_go.Response[v1.ProfileTypesResponse], error)
	LabelValues(context.Context, *connect_go.Request[v1.LabelValuesRequest]) (*connect_go.Response[v1.LabelValuesResponse], error)
	LabelNames(context.Context, *connect_go.Request[v1.LabelNamesRequest]) (*connect_go.Response[v1.LabelNamesResponse], error)
	Series(context.Context, *connect_go.Request[v1.SeriesRequest]) (*connect_go.Response[v1.SeriesResponse], error)
	SelectMergeStacktraces(context.Context, *connect_go.Request[v1.SelectMergeStacktracesRequest]) (*connect_go.Response[v1.SelectMergeStacktracesResponse], error)
	SelectSeries(context.Context, *connect_go.Request[v1.SelectSeriesRequest]) (*connect_go.Response[v1.SelectSeriesResponse], error)
}

// NewQuerierServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQuerierServiceHandler(svc QuerierServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/querier.v1.QuerierService/ProfileTypes", connect_go.NewUnaryHandler(
		"/querier.v1.QuerierService/ProfileTypes",
		svc.ProfileTypes,
		opts...,
	))
	mux.Handle("/querier.v1.QuerierService/LabelValues", connect_go.NewUnaryHandler(
		"/querier.v1.QuerierService/LabelValues",
		svc.LabelValues,
		opts...,
	))
	mux.Handle("/querier.v1.QuerierService/LabelNames", connect_go.NewUnaryHandler(
		"/querier.v1.QuerierService/LabelNames",
		svc.LabelNames,
		opts...,
	))
	mux.Handle("/querier.v1.QuerierService/Series", connect_go.NewUnaryHandler(
		"/querier.v1.QuerierService/Series",
		svc.Series,
		opts...,
	))
	mux.Handle("/querier.v1.QuerierService/SelectMergeStacktraces", connect_go.NewUnaryHandler(
		"/querier.v1.QuerierService/SelectMergeStacktraces",
		svc.SelectMergeStacktraces,
		opts...,
	))
	mux.Handle("/querier.v1.QuerierService/SelectSeries", connect_go.NewUnaryHandler(
		"/querier.v1.QuerierService/SelectSeries",
		svc.SelectSeries,
		opts...,
	))
	return "/querier.v1.QuerierService/", mux
}

// UnimplementedQuerierServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedQuerierServiceHandler struct{}

func (UnimplementedQuerierServiceHandler) ProfileTypes(context.Context, *connect_go.Request[v1.ProfileTypesRequest]) (*connect_go.Response[v1.ProfileTypesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("querier.v1.QuerierService.ProfileTypes is not implemented"))
}

func (UnimplementedQuerierServiceHandler) LabelValues(context.Context, *connect_go.Request[v1.LabelValuesRequest]) (*connect_go.Response[v1.LabelValuesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("querier.v1.QuerierService.LabelValues is not implemented"))
}

func (UnimplementedQuerierServiceHandler) LabelNames(context.Context, *connect_go.Request[v1.LabelNamesRequest]) (*connect_go.Response[v1.LabelNamesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("querier.v1.QuerierService.LabelNames is not implemented"))
}

func (UnimplementedQuerierServiceHandler) Series(context.Context, *connect_go.Request[v1.SeriesRequest]) (*connect_go.Response[v1.SeriesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("querier.v1.QuerierService.Series is not implemented"))
}

func (UnimplementedQuerierServiceHandler) SelectMergeStacktraces(context.Context, *connect_go.Request[v1.SelectMergeStacktracesRequest]) (*connect_go.Response[v1.SelectMergeStacktracesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("querier.v1.QuerierService.SelectMergeStacktraces is not implemented"))
}

func (UnimplementedQuerierServiceHandler) SelectSeries(context.Context, *connect_go.Request[v1.SelectSeriesRequest]) (*connect_go.Response[v1.SelectSeriesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("querier.v1.QuerierService.SelectSeries is not implemented"))
}
