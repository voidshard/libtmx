package v1

import "encoding/xml"

// Represents an arbitrary named group of things
//
type group struct {
	XMLName xml.Name `xml:"group"`

	// attrs
	Name    string  `xml:"name"`
	X       int     `xml:"x,attr,optional,omitempty"`
	Y       int     `xml:"y,attr,optional,omitempty"`
	Visible int     `xml:"visible,attr"`
	OffsetX int     `xml:"offsetx,attr"`
	OffsetY int     `xml:"offsety,attr"`
	Opacity float32 `xml:"opacity,attr"`

	// subsections
	Properties   properties    `xml:"properties,optional,omitempty"`
	TileLayers   []tileLayer   `xml:"layer,optional,omitempty"`
	ImageLayers  []imageLayer  `xml:"imagelayer,optional,omitempty"`
	ObjectGroups []objectGroup `xml:"objectgroup,optional,omitempty"`
	Groups       []group       `xml:"group,optional,omitempty"`
}
