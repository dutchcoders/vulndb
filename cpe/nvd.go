package cpe

// Nvd represents the parsing result of a single NVD Document
type Nvd struct {
	Entries []struct {
		ID         string      `xml:"id,attr"`
		Products   []*Product  `xml:"vulnerable-software-list>product"`
		CveID      string      `xml:"cve-id"`
		Summary    string      `xml:"summary"`
		Severity   string      `xml:"cvss>base_metrics>score"`
		References []Reference `xml:"references>reference"`
	} `xml:"entry"`
}

// Reference represents a reference link from an NVD Entry
type Reference struct {
	Name string `xml:",chardata"`
	URL  string `xml:"href,attr"`
}
