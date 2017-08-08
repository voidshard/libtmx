package v1

import "encoding/xml"

type tileOffset struct {
	XMLName xml.Name `xml:"tileoffset"`

	// attrs
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}
