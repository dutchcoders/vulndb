package cli

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/dutchcoders/vulndb/cpe"
	"github.com/spf13/cobra"
)

// BuildBuildCommand returns a command for building vulnerability databases from a list of files.
func BuildBuildCommand() *cobra.Command {
	var dbFile string
	var force bool

	cmd := &cobra.Command{
		Use:   "build [<file>...]",
		Short: "Build the vulnerability DB from a set of NVD CPE files.",
		Run: func(cmd *cobra.Command, args []string) {

			if force == true {
				fmt.Printf("Force removing %s\n", dbFile)
				os.RemoveAll(dbFile)
			}

			err := os.MkdirAll(defaultBaseDir(), 0777)
			check(err)

			mapping := bleve.NewIndexMapping()
			mapping.DefaultAnalyzer = "keyword"

			summaryMapping := bleve.NewDocumentMapping()
			mapping.AddDocumentMapping("summary", summaryMapping)

			index, err := bleve.New(dbFile, mapping)
			check(err)

			for _, f := range args {
				fmt.Printf("Processing %s... ", f)
				file, err := ioutil.ReadFile(f)
				check(err)

				var result cpe.Nvd

				err = xml.Unmarshal(file, &result)
				check(err)

				count := 0
				for _, entry := range result.Entries {
					index.Index(entry.CveID, entry)
					count++
				}

				fmt.Printf("Finished %s, processed %d entries\n", f, count)
			}
		},
	}

	cmd.Flags().StringVarP(&dbFile, "db-file", "d", defaultDbFile(), "vulnerability db file to build")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "overwrite existing vulnerability db")

	return cmd
}
