package hcl

import (
	"reflect"
	"testing"
)

func TestDataSourceConfigParsing(t *testing.T) {
	expected := &Config{
		DataSources: []DataSourceConfig{
			DataSourceConfig{
				Name: "prometheus",
				Type: "prometheus",
			},
		},
	}

	config, err := ParseConfig(testDataSource)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(config, expected) {
		t.Error("Data Source structure differed from expectation")
	}
}

const testDataSource = `data_source "prometheus" {
    type = "prometheus"
}`
