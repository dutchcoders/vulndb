package cpe

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

const numEntries = 8

func TestParseNvd(t *testing.T) {
	var result Nvd

	data, _ := ioutil.ReadFile("../data/nvd-test.xml")

	xml.Unmarshal(data, &result)

	if len(result.Entries) != numEntries {
		t.Fatalf("Expected %d entries, got: %d", numEntries, len(result.Entries))
	}
}

func TestIncludesSeverity(t *testing.T) {
	var result Nvd
	data, _ := ioutil.ReadFile("../data/nvd-test.xml")

	xml.Unmarshal(data, &result)

	if result.Entries[0].Severity != "1.9" {
		t.Fatalf("Expected %s severity, got: %s", "1.9", result.Entries[0].Severity)
	}
}

func TestIncludesReferences(t *testing.T) {
	var result Nvd
	data, _ := ioutil.ReadFile("../data/nvd-test.xml")

	xml.Unmarshal(data, &result)

	reference := result.Entries[0].References[0]

	if len(result.Entries[0].References) != 1 {
		t.Fatalf("Expected %d references, got: %d", 1, result.Entries[0].References)
	}

	if reference.URL != "http://technet.microsoft.com/security/bulletin/MS15-006" {
		t.Fatalf("Expected url to be: %s, but got: %s", "http://technet.microsoft.com/security/bulletin/MS15-006", reference.URL)
	}
}
