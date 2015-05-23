package cli

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/dutchcoders/vulndb/cpe"
	"github.com/spf13/cobra"
)

var feedsModified = []string{
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-Modified.xml.gz",
}

// BuildUpdateCommand returns a command for updating vulnerability databases from a list of files.
func BuildUpdateCommand() *cobra.Command {
	var dbFile string
	var full bool

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the existing vulnerability DB using the modified NVD CPE feed.",
		Run: func(cmd *cobra.Command, args []string) {

			err := os.MkdirAll(defaultBaseDir(), 0777)
			check(err)

			index, err := bleve.Open(dbFile)
			if err == bleve.ErrorIndexPathDoesNotExist {
				mapping := bleve.NewIndexMapping()
				mapping.DefaultAnalyzer = "keyword"

				summaryMapping := bleve.NewDocumentMapping()
				mapping.AddDocumentMapping("summary", summaryMapping)

				index, err = bleve.New(dbFile, mapping)
			}

			check(err)

			feeds := feedsModified

			if full {
				feeds = feedsComplete
			}

			for _, f := range feeds {
				fmt.Printf("Processing %s... ", f)
				resp, err := http.Get(f)
				check(err)

				reader := resp.Body
				defer reader.Close()

				switch resp.Header.Get("Content-Type") {
				case "application/x-gzip":
					reader, err = gzip.NewReader(reader)
				default:
				}

				check(err)

				var result cpe.Nvd
				err = xml.NewDecoder(reader).Decode(&result)
				check(err)

				count := 0

				batch := index.NewBatch()

				for _, entry := range result.Entries {
					batch.Index(entry.CveID, entry)
					count++

					if count%1000 == 0 {
						index.Batch(batch)
						batch = index.NewBatch()
					}

					if count%(len(result.Entries)/10) == 0 {
						fmt.Printf("#")
					}
				}

				index.Batch(batch)

				fmt.Printf("\b finished processing %d entries\n", count)
			}
		},
	}

	cmd.Flags().BoolVarP(&full, "full", "f", false, "full update")
	cmd.Flags().StringVarP(&dbFile, "db-file", "d", defaultDbFile(), "vulnerability db file to build")

	return cmd
}
