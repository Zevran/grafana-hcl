package commands

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type GenerateCommand struct {
	Ui cli.Ui
}

func (c *GenerateCommand) Run(args []string) int {
	c.Ui.Output("Plop")
	c.Ui.Output(fmt.Sprintf("%+v", args))
	return 0
}

func (c *GenerateCommand) Help() string {
	return `Generate JSON file from an HCL configuration file.`
}

func (c *GenerateCommand) Synopsis() string {
	return `Generate JSON file from an HCL configuration file.`
}
