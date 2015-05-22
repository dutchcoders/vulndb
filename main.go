package main

import (
	"fmt"

	"github.com/dutchcoders/vulndb/cli"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "vulndb",
		Short: "Vulndb is a command line tool for searching the NIST Vulnerability Database.",
	}

	searchCmd := cli.BuildSearchCommand()

	var buildCmd = &cobra.Command{
		Use:   "build [<file>...]",
		Short: "Build the vulnerability DB from a set of NVD CPE files.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	rootCmd.AddCommand(searchCmd, buildCmd)
	rootCmd.Execute()

	// mapping := bleve.NewIndexMapping()

	// index, _ := bleve.Open("test.bleve")

	// var result cpe.Nvd

	// data, _ := ioutil.ReadFile("./nvdcve-2.0-2015.xml")
	//
	// xml.Unmarshal(data, &result)
	//
	// for _, entry := range result.Entries {
	// 	index.Index(entry.CveID, entry)
	// }
	//
}
