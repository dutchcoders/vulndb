package cpe

// Nvd represents the parsing result of a single NVD Document
type Nvd struct {
	Entries []struct {
		ID       string     `xml:"id,attr"`
		Products []*Product `xml:"vulnerable-software-list>product"`
		CveID    string     `xml:"cve-id"`
		Summary  string     `xml:"summary"`
	} `xml:"entry"`
}
