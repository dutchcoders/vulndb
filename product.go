package main

import "strings"

// Product represents a product unmarshalled from a CPE URL
type Product struct {
	Part      string
	Vendor    string
	Product   string
	Version   string
	Update    string
	Edition   string
	SWEdition string
	TargetSW  string
	TargetHW  string
	Other     string
	Language  string
}

const requiredLength = 7
const requiredEditionLength = 6

// UnmarshalText satisfies the text.Unmarshaler interface for XML parsing
func (p *Product) UnmarshalText(text []byte) error {
	cpeStr := strings.TrimLeft(string(text), "cpe:/")
	cpeList := make([]string, requiredLength)
	copy(cpeList, strings.Split(cpeStr, ":"))

	p.Part = cpeList[0]
	p.Vendor = cpeList[1]
	p.Product = cpeList[2]
	p.Version = cpeList[3]
	p.Update = cpeList[4]

	edition := make([]string, requiredEditionLength)
	copy(edition, strings.Split(cpeList[5], "~"))

	p.Edition = edition[1]
	p.SWEdition = edition[2]
	p.TargetSW = edition[3]
	p.TargetHW = edition[4]
	p.Other = edition[5]

	p.Language = cpeList[6]
	return nil
}
