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

	searchCmd := cli.BuildSearchCommand()
	buildCmd := cli.BuildBuildCommand()

	rootCmd.AddCommand(searchCmd, buildCmd)
	rootCmd.Execute()
}
