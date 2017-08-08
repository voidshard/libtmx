package v1

import (
	"encoding/xml"
	"image"
	"image/color"
	"github.com/voidshard/libtmx/common"
)

// More properly called an 'object layer' - represents a layer of arbitrary object(s)
//
type objectGroup struct {
	XMLName xml.Name `xml:"objectgroup"`

	// attrs
	Name   string `xml:"name,attr"`
	Colour string `xml:"color,attr"`

	// attrs optional
	X         int     `xml:"x,attr,optional,omitempty"`
	Y         int     `xml:"y,attr,optional,omitempty"`
	Width     int     `xml:"width,attr,optional,omitempty"`
	Height    int     `xml:"height,attr,optional,omitempty"`
	Visible   int     `xml:"visible,attr,optional,omitempty"`
	OffsetX   int     `xml:"offsetx,attr,optional,omitempty"`
	OffsetY   int     `xml:"offsety,attr,optional,omitempty"`
	Opacity   float64 `xml:"opacity,attr,optional,omitempty"`
	DrawOrder string  `xml:"draworder,attr,optional,omitempty"`

	// subsections
	Objects []object `xml:"object"`

	// subsections optional
	Properties properties `xml:"properties,optional,omitempty"`
}

// Get the colour for this group
//
func (o *objectGroup) GroupColour() (*color.RGBA, error) {
	return decodeHexColour(o.Colour)
}

// Set the colour for this group
//
func (o *objectGroup) SetGroupColour(rgba *color.RGBA) {
	o.Colour = encodeHexColour(rgba)
}

// Represents an object read in from xml.
// This can be either an Ellipse, Polygon, Polyline or Text ..
//
type object struct {
	XMLName xml.Name `xml:"object"`

	// attrs
	Name     string `xml:"name,attr,optional,omitempty"`
	Type     string `xml:"type,attr,optional,omitempty"`
	Id       int    `xml:"id,attr,optional,omitempty"`
	Gid      int    `xml:"gid,attr,optional,omitempty"`
	X        int    `xml:"x,attr,optional,omitempty"`
	Y        int    `xml:"y,attr,optional,omitempty"`
	Width    int    `xml:"width,attr,optional,omitempty"`
	Height   int    `xml:"height,attr,optional,omitempty"`
	Visible  int    `xml:"visible,attr,optional,omitempty"`
	Rotation int    `xml:"rotation,attr,optional,omitempty"`

	// subsections
	Properties properties `xml:"properties,optional,omitempty"`
	Ellipse    ellipse    `xml:"ellipse,optional,omitempty"`
	Polygon    polygon    `xml:"polygon,optional,omitempty"`
	Polyline   polyline   `xml:"polyline,optional,omitempty"`
	Text       text       `xml:"text,optional,omitempty"`
}

type ellipse struct {
	XMLName xml.Name `xml:"ellipse"`

	// attrs
	X      int `xml:"x,attr"`
	Y      int `xml:"y,attr"`
	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

type polygon struct {
	XMLName xml.Name `xml:"polygon"`

	// attrs
	RawPoints string `xml:"points,attr"`
}

func (p *polygon) Points() ([]image.Point, error) {
	return decodePoints(p.RawPoints)
}

func (p *polygon) SetPoints(in []image.Point) {
	p.RawPoints = encodePoints(in)
}

type polyline struct {
	XMLName xml.Name `xml:"polyline"`

	// attrs
	RawPoints string `xml:"points,attr"`
}

func (p *polyline) Points() ([]image.Point, error) {
	return decodePoints(p.RawPoints)
}

func (p *polyline) SetPoints(in []image.Point) {
	p.RawPoints = encodePoints(in)
}

type text struct {
	XMLName xml.Name `xml:"text"`

	// attrs
	FontFamily string `xml:"fontfamily,attr,optional,omitempty"`
	PixelSize  int    `xml:"pixelsize,attr,optional,omitempty"`
	Colour     string `xml:"color,attr,optional,omitempty"`

	AlignH     string `xml:"halign,attr,optional,omitempty"`
	AlignV     string `xml:"valign,attr,optional,omitempty"`

	Bold       int    `xml:"bold,attr,optional,omitempty"`
	Italic     int    `xml:"italic,attr,optional,omitempty"`
	Underline  int    `xml:"underline,attr,optional,omitempty"`
	Kerning    int    `xml:"kerning,attr,optional,omitempty"`
	Strikeout  int    `xml:"strikeout,attr,optional,omitempty"`
	Wrap       int    `xml:"wrap,attr,optional,omitempty"`

	// data
	Value string `xml:",chardata"`
}

// Inflate the given text object, setting it's internal values
//
func (t *text) inflate() {
	if t.FontFamily == "" {
		t.FontFamily = common.DefaultFontFamily
	}
	if t.PixelSize == 0 {
		t.PixelSize = common.DefaultPixelSize
	}
	if t.AlignH == "" {
		t.AlignH = common.DefaultHAlign
	}
	if t.AlignV == "" {
		t.AlignV = common.DefaultVAlign
	}
}

func (t *text) TextColour() (*color.RGBA, error) {
	return decodeHexColour(t.Colour)
}

func (t *text) SetTextColour(rgba *color.RGBA) {
	t.Colour = encodeHexColour(rgba)
}
