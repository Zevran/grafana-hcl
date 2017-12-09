package hcl

import (
	multierror "github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl"
)

type Resource interface {
	ValidateHCL() error
	GenerateJSON() ([]byte, error)
}

type Config struct {
	DataSources []DataSourceConfig `hcl:"data_source"`
}

func ParseConfig(hclText string) (*Config, error) {
	result := &Config{}
	var errors *multierror.Error

	hclParseTree, err := hcl.Parse(hclText)
	if err != nil {
		return nil, err
	}

	if err := hcl.DecodeObject(&result, hclParseTree); err != nil {
		return nil, err
	}

	return result, errors.ErrorOrNil()
}

func Output(r Resource) (string, error) {
	jsonContent, err := r.GenerateJSON()

	return string(jsonContent), err
}

func Validate(r Resource) error {
	return r.ValidateHCL()
}
