package main

import (
	"fmt"
	"os"

	"github.com/Zevran/grafana-hcl/commands"
	"github.com/mitchellh/cli"
)

func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("grafana-hcl", "0.1.0")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"generate": func() (cli.Command, error) {
			return &commands.GenerateCommand{Ui: ui}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	os.Exit(exitStatus)
}
