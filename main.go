package main

import (
	"github.com/dutchcoders/vulndb/cli"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "vulndb",
		Short: "Vulndb is a command line tool for searching the NIST Vulnerability Database.",
	}

	initCmd := cli.BuildInitCommand()
	searchCmd := cli.BuildSearchCommand()
	buildCmd := cli.BuildBuildCommand()
	updateCmd := cli.BuildUpdateCommand()

	rootCmd.AddCommand(initCmd, searchCmd, buildCmd, updateCmd)
	rootCmd.Execute()
}
