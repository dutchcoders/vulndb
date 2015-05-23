package cli

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/spf13/cobra"
)

const requiredLength = 7

var queryName = []string{"Part", "Vendor", "Product", "Version", "Update", "Edition", "Language"}

// BuildSearchCommand returns a new reference to a search command.
func BuildSearchCommand() *cobra.Command {
	var dbFile string
	var limit, offset int
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
			search.Size = limit
			search.From = offset
			search.Highlight = bleve.NewHighlightWithStyle("ansi")
			search.Fields = []string{"*"}

			searchResults, _ := index.Search(search)

			for _, hit := range searchResults.Hits {
				fmt.Println(hit.ID)

				severity := hit.Fields["Severity"]
				if fragment, ok := hit.Fragments["Severity"]; ok {
					severity = fragment[0]
				}
				fmt.Printf("Severity: %s\n", severity)

				vendor := hit.Fields["Products.Vendor"]
				if fragment, ok := hit.Fragments["Products.Vendor"]; ok {
					vendor = fragment[0]
				}
				fmt.Printf("Vendor: %s\n", vendor)

				summary := hit.Fields["Summary"]
				if fragment, ok := hit.Fragments["Summary"]; ok {
					summary = fragment[0]
				}
				fmt.Printf("Summary: %s\n", summary)

				fmt.Printf("References:\n")
				fmt.Println(buildReferences(hit.Fields["References.URL"]))
				fmt.Println("")
			}

			fmt.Printf("Showing entries %d-%d of %d. Completed in %s\n", offset+1, offset+limit, searchResults.Total, searchResults.Took)
		},
	}

	cmd.Flags().StringVarP(&dbFile, "db-file", "d", defaultDbFile(), "vulnerability db file to use")
	cmd.Flags().IntVarP(&limit, "limit", "l", 10, "number of results to display")
	cmd.Flags().IntVarP(&offset, "from", "f", 0, "start results from number")

	return cmd
}

func buildReferences(refs interface{}) string {
	switch reflect.TypeOf(refs).Name() {
	case "string":
		return refs.(string)
	default:
		var strs = make([]string, len(refs.([]interface{})))

		for i, ref := range refs.([]interface{}) {
			strs[i] = ref.(string)
		}

		return strings.Join(strs, "\n")
	}
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
