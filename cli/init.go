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

//"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-Modified.xml.gz",
//"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-Recent.xml.gz",
var feedsComplete = []string{
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2002.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2003.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2004.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2005.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2006.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2007.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2008.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2009.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2010.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2011.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2012.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2013.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2014.xml.gz",
	"https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2015.xml.gz",
}

// BuildInitCommand returns a command for initialising vulnerability databases from the NVD CPE feeds.
func BuildInitCommand() *cobra.Command {
	var dbFile string

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initializes the database with the NVD CPE feeds.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Initializing database. Downloading complete NVD CPE feed history. This may take a while.")

			err := os.MkdirAll(defaultBaseDir(), 0777)
			check(err)

			index, err := bleve.Open(dbFile)
			check(err)

			for _, f := range feedsComplete {
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
				for _, entry := range result.Entries {
					index.Index(entry.CveID, entry)
					count++
				}

				fmt.Printf("finished processing %d entries\n", count)
			}
		},
	}

	cmd.Flags().StringVarP(&dbFile, "db-file", "d", defaultDbFile(), "vulnerability db file to build")

	return cmd
}
