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

// BuildUpdateCommand returns a command for updating vulnerability databases from a list of files.
func BuildUpdateCommand() *cobra.Command {
	var dbFile string

	cmd := &cobra.Command{
		Use:   "update [<file>...]",
		Short: "Update the existing vulnerability DB from a set of NVD CPE files.",
		Run: func(cmd *cobra.Command, args []string) {

			err := os.MkdirAll(defaultBaseDir(), 0777)
			check(err)

			index, err := bleve.Open(dbFile)
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

	return cmd
}
