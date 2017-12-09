package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/Zevran/grafana-hcl/hcl"
	"github.com/mitchellh/cli"
)

type GenerateCommand struct {
	Ui cli.Ui
}

func (c *GenerateCommand) Run(args []string) int {
	fileName := args[0]

	wdPath, err := Getwd()
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	c.Ui.Info(fmt.Sprintf("Generating JSON from %s\n", fileName))

	filePath := fmt.Sprintf("%s/%s", wdPath, fileName)
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	conf, err := hcl.ParseConfig(string(dat))
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	hasError := false
	for index := 0; index < len(conf.DataSources); index++ {
		err := hcl.Validate(&conf.DataSources[index])
		if err != nil {
			hasError = true
			c.Ui.Error(err.Error())
		}

	}

	if hasError {
		return 1
	}

	for index := 0; index < len(conf.DataSources); index++ {
		content, err := hcl.Output(&conf.DataSources[index])
		if err != nil {
			c.Ui.Error(err.Error())
		}

		fmt.Println(string(content))
	}

	return 0
}

func (c *GenerateCommand) Help() string {
	return `Generate JSON file from an HCL configuration file.`
}

func (c *GenerateCommand) Synopsis() string {
	return `Generate JSON file from an HCL configuration file.`
}
