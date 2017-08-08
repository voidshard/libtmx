package v1

import (
	"encoding/xml"
	"github.com/voidshard/libtmx/common"
)

// Image layer - a single image embedded as a layer
//
type imageLayer struct {
	XMLName xml.Name `xml:"imagelayer"`

	// attrs
	Name    string  `xml:"name"`
	X       int     `xml:"x,attr,optional,omitempty"`
	Y       int     `xml:"y,attr,optional,omitempty"`
	Visible int     `xml:"visible,attr"`
	OffsetX int     `xml:"offsetx,attr"`
	OffsetY int     `xml:"offsety,attr"`
	Opacity float64 `xml:"opacity,attr"`

	// subsections
	Properties properties `xml:"properties,optional,omitempty"`
	Image      imageData  `xml:"image,optional,omitempty"`
}

// Inflate this layer to be a common.ImageLayer and add it to the given map
//
func (o *imageLayer) inflate(parent *common.Map) {
	layer := parent.NewImageLayer(o.Name, o.Image.Source)

	layer.OffsetX = o.OffsetX
	layer.OffsetY = o.OffsetY
	layer.Opacity = o.Opacity
	layer.Height = o.Image.Height
	layer.Width = o.Image.Width
	layer.ImageSource = o.Image.Source
	layer.ImageFormat = o.Image.Format

	col, err := o.Image.TransparentColour()
	if err == nil {
		layer.TransparentColour = col
	}
	layer.UpdateProperties(o.Properties.inflate()...)
}

// Deflate the given common.ImageLayer to be an imageLayer for writing to XML
//
func deflateImageLayer(in *common.ImageLayer) imageLayer {
	return imageLayer{
		Name: in.Name,
		Visible: boolToInt(in.Visible),
		OffsetX: in.OffsetX,
		OffsetY: in.OffsetY,
		Opacity: in.Opacity,
		Properties: deflateProperties(in.Properties()),
		Image: imageData{
			Source: in.ImageSource,
			TransColour: encodeHexColour(in.TransparentColour),
			Width: in.Width,
			Height: in.Height,
			Format: in.ImageFormat,
		},
	}
}
