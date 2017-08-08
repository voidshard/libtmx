package v1

import "encoding/xml"

// Block of tile data inside a tilelayer
//
type dataBlock struct {
	XMLName xml.Name `xml:"data"`

	// attrs
	Encoding    string `xml:"encoding,attr,optional,omitempty"`
	Compression string `xml:"compression,attr,optional,omitempty"`

	// subsections
	//tile tile `xml:"tile,optional,omitempty"` // NB: Deliberately removed: creates circular struct

	// value
	Value string `xml:",chardata"`
}
