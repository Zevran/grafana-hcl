package hcl

import (
	"log"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl"
)

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

	log.Printf("%v\n", result)

	return result, errors.ErrorOrNil()
}
