package hcl

import (
	"reflect"
	"testing"
)

func TestDataSourceConfigParsing(t *testing.T) {
	expected := &Config{
		DataSources: []DataSourceConfig{
			DataSourceConfig{
				Name:      "prometheus",
				Type:      "prometheus",
				Access:    "proxy",
				URL:       "http://mydatasource.com",
				Database:  "ohmytsdb",
				User:      "johndoe",
				Password:  "cookieisstrong",
				IsDefault: true,
				BasicAuth: DataSourceBasicAuthConfig{
					User:     "johndoe",
					Password: "cookieisstrong",
				},
				JSONData: `{
    "authType": "keys",
    "defaultRegion": "us-west-1"
}
`,
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
	access = "proxy",
	url = "http://mydatasource.com"
    database = "ohmytsdb"
    user = "johndoe"
	password = "cookieisstrong"
    is_default = true

    basic_auth {
        user = "johndoe",
        password = "cookieisstrong"
    }

    json_data = <<EOF
{
    "authType": "keys",
    "defaultRegion": "us-west-1"
}
EOF
}`
