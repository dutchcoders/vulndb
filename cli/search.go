package cli

import (
	"fmt"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/spf13/cobra"
)

const requiredLength = 7

var queryName = []string{"Part", "Vendor", "Product", "Version", "Update", "Edition", "Language"}

// BuildSearchCommand returns a new reference to a search command.
func BuildSearchCommand() *cobra.Command {
	var dbFile string
	var query bleve.Query

	cmd := &cobra.Command{
		Use:   "search <CPE URI>",
		Short: "Search for vulnerabilities based on the CPE URI.",
		Run: func(cmd *cobra.Command, args []string) {
			index, _ := bleve.Open(dbFile)
			term := args[0]

			if isCpeURI(term) {
				query = buildCpeSearch(term)
			} else {
				query = bleve.NewMatchQuery(term)
			}

			search := bleve.NewSearchRequest(query)
			search.Fields = []string{"Summary"}

			searchResults, _ := index.Search(search)

			for _, hit := range searchResults.Hits {
				fmt.Println(hit.ID)
				fmt.Println(hit.Fields["Summary"])
				fmt.Println("")
			}

			fmt.Printf("Found: %d. Completed in %s\n", searchResults.Total, searchResults.Took)
		},
	}

	cmd.Flags().StringVarP(&dbFile, "db-file", "d", defaultDbFile(), "vulnerability db file to use")

	return cmd
}

func isCpeURI(s string) bool {
	return strings.Index(s, "cpe:/") == 0
}

func buildCpeSearch(term string) bleve.Query {
	cpeStr := strings.TrimLeft(term, "cpe:/")
	cpeList := make([]string, requiredLength)
	copy(cpeList, strings.Split(cpeStr, ":"))

	var queries []bleve.Query

	for i, term := range cpeList {
		if term != "" && term != "*" && term != "-" {
			query := bleve.NewPhraseQuery([]string{term}, "Products."+queryName[i])
			queries = append(queries, query)
		}
	}

	conjunctionQuery := bleve.NewConjunctionQuery(queries)

	return conjunctionQuery
}
