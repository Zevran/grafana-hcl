package hcl

import (
	"encoding/json"

	grafana "github.com/grafana-tools/sdk"
)

type DataSourceConfig struct {
	Name      string `hcl:",key"`
	Type      string `hcl:"type"`
	Access    string `hcl:"access"`
	URL       string `hcl:"url"`
	Database  string `hcl:"database"`
	User      string `hcl:"user"`
	Password  string `hcl:"password"`
	IsDefault bool   `hcl:"is_default"`

	JSONData  string                    `hcl:"json_data"`
	BasicAuth DataSourceBasicAuthConfig `hcl:"basic_auth"`
}

type DataSourceBasicAuthConfig struct {
	User     string `hcl:"user"`
	Password string `hcl:"password"`
}

func (d *DataSourceConfig) GenerateJSON() ([]byte, error) {

	ds := &grafana.Datasource{
		Name:      d.Name,
		Type:      d.Type,
		Access:    d.Access,
		URL:       d.URL,
		IsDefault: d.IsDefault,
	}

	if d.Database != "" && d.User != "" && d.Password != "" {
		ds.Database = &d.Database
		ds.User = &d.User
		ds.Password = &d.Password
	}

	if d.BasicAuth.User != "" && d.BasicAuth.Password != "" {
		enable := true
		ds.BasicAuth = &enable
		ds.BasicAuthUser = &d.BasicAuth.User
		ds.BasicAuthPassword = &d.BasicAuth.Password
	}

	if d.JSONData != "" {
		ds.JSONData = d.JSONData
	}

	byteContent, err := json.Marshal(&ds)

	return byteContent, err
}
