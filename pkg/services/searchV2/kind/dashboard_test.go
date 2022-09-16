package kind

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/grafana/grafana-plugin-sdk-go/experimental"
	"github.com/grafana/grafana/pkg/services/searchV2/dslookup"
	"github.com/stretchr/testify/require"
)

func dsLookup() dslookup.DatasourceLookup {
	return dslookup.CreateDatasourceLookup([]*dslookup.DatasourceQueryResult{
		{
			UID:       "P8045C56BDA891CB2",
			Type:      "cloudwatch",
			Name:      "cloudwatch-name",
			IsDefault: false,
		},
		{
			UID:       "PD8C576611E62080A",
			Type:      "testdata",
			Name:      "gdev-testdata",
			IsDefault: false,
		},
		{
			UID:       "dgd92lq7k",
			Type:      "frser-sqlite-datasource",
			Name:      "frser-sqlite-datasource-name",
			IsDefault: false,
		},
		{
			UID:       "sqlite-1",
			Type:      "sqlite-datasource",
			Name:      "SQLite Grafana",
			IsDefault: false,
		},
		{
			UID:       "sqlite-2",
			Type:      "sqlite-datasource",
			Name:      "SQLite Grafana2",
			IsDefault: false,
		},
		{
			UID:       "default.uid",
			Type:      "default.type",
			Name:      "default.name",
			IsDefault: true,
		},
	})
}

func TestReadDashboard(t *testing.T) {
	// inputs := []string{
	// 	"check-string-datasource-id",
	// 	"all-panels",
	// 	"panel-graph/graph-shared-tooltips",
	// 	"datasource-variable",
	// 	"default-datasource-variable",
	// 	"empty-datasource-variable",
	// 	"repeated-datasource-variables",
	// 	"string-datasource-variable",
	// 	"datasource-variable-no-curly-braces",
	// 	"all-selected-multi-datasource-variable",
	// 	"all-selected-single-datasource-variable",
	// 	"repeated-datasource-variables-with-default",
	// 	"mixed-datasource-with-variable",
	// 	"special-datasource-types",
	// 	"panels-without-datasources",
	// }

	devdash := "../../../../devenv/dev-dashboards/"

	indexer := dashboardIndexer{lookup: dsLookup()}

	builder := toIndexBuilder(indexer.GetIndex())
	err := filepath.Walk(devdash,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(path, ".json") {
				file, err := os.Open(path)
				if err != nil {
					return err
				}

				uid := path[len("../../../../"):]
				row, err := indexer.Index(uid, file)
				if err != nil {
					return err
				}

				err = builder.append(row)
				if err != nil {
					return err
				}
			}
			return nil
		})
	require.NoError(t, err)

	for _, frame := range builder.frames() {
		experimental.CheckGoldenJSONFrame(t, "testdata", frame.Name, frame, true)
	}

	require.Nil(t, indexer)
}
