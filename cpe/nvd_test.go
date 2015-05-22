package cpe

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

const numEntries = 7

func TestParseNvd(t *testing.T) {
	var result Nvd

	data, _ := ioutil.ReadFile("../data/nvd-test.xml")

	xml.Unmarshal(data, &result)

	if len(result.Entries) != numEntries {
		t.Fatalf("Expected %d entries, got: %d", numEntries, len(result.Entries))
	}
}
