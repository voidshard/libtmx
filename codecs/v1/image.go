package v1

import (
	"encoding/xml"
	"image/color"
)

// Image data block that belongs to a tile or image layer.
// Technically the spec allows images to be embedded, but Tiled itself doesn't support this.
//
type imageData struct {
	XMLName xml.Name `xml:"image"`

	// attrs
	Format      string `xml:"format,attr,optional,omitempty"`
	Id          int    `xml:"id,optional,omitempty"`
	Source      string `xml:"source,attr,optional,omitempty"`
	TransColour string `xml:"trans,attr,optional,omitempty"`
	Width       int    `xml:"width,attr,optional,omitempty"`
	Height      int    `xml:"height,attr,optional,omitempty"`

	// subsections
	Data dataBlock `xml:"data,optional,omitempty"`
}

// Get the color set to be transparent for this block
//
func (i *imageData) TransparentColour() (*color.RGBA, error) {
	return decodeHexColour(i.TransColour)
}

// Set the transparent colour for this block
//
func (i *imageData) SetTransparentColour(rgba *color.RGBA) {
	i.TransColour = encodeHexColour(rgba)
}
