package main

import (
	"reflect"
	"testing"
)

func TestSimpleParseProduct(t *testing.T) {
	cpeURL := []byte("cpe:/o:canonical:ubuntu_linux:14.10")
	var p Product
	result := &Product{
		Part:    "o",
		Vendor:  "canonical",
		Product: "ubuntu_linux",
		Version: "14.10",
	}

	p.UnmarshalText(cpeURL)

	if !reflect.DeepEqual(result, &p) {
		t.Fatalf("Expected %s, got %s", result, p)
	}
}

func TestSimpleProductWithEmptyFields(t *testing.T) {
	cpeURL := []byte("cpe:/o:canonical:ubuntu_linux:14.10:::")
	var p Product
	result := &Product{
		Part:    "o",
		Vendor:  "canonical",
		Product: "ubuntu_linux",
		Version: "14.10",
		Update:  "",
	}

	p.UnmarshalText(cpeURL)

	if !reflect.DeepEqual(result, &p) {
		t.Fatalf("Expected %s, got %s", result, p)
	}
}

func TestFullProduct(t *testing.T) {
	cpeURL := []byte("cpe:/o:vendor:product:14.10:update:~edition~sw_edition~target_sw~target_hw~other:language")
	var p Product
	result := &Product{
		Part:      "o",
		Vendor:    "vendor",
		Product:   "product",
		Version:   "14.10",
		Update:    "update",
		Edition:   "edition",
		SWEdition: "sw_edition",
		TargetSW:  "target_sw",
		TargetHW:  "target_hw",
		Other:     "other",
		Language:  "language",
	}

	p.UnmarshalText(cpeURL)

	if !reflect.DeepEqual(result, &p) {
		t.Fatalf("Expected %s, got %s", result, p)
	}
}

func TestPartialEdition(t *testing.T) {
	cpeURL := []byte("cpe:/o:microsoft:windows_server_2012:r2::~~~x64~~")
	var p Product
	result := &Product{
		Part:     "o",
		Vendor:   "microsoft",
		Product:  "windows_server_2012",
		Version:  "r2",
		TargetSW: "x64",
	}

	p.UnmarshalText(cpeURL)

	if !reflect.DeepEqual(result, &p) {
		t.Fatalf("Expected %s, got %s", result, p)
	}
}
