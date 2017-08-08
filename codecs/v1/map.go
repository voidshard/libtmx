package v1

import (
	"encoding/xml"
	"image/color"
	"os"
	"github.com/voidshard/libtmx/common"
	"io/ioutil"
)

type tileMap struct {
	XMLName xml.Name `xml:"map"`

	// attrs
	Width      int `xml:"width,attr"`
	Height     int `xml:"height,attr"`
	TileWidth  int `xml:"tilewidth,attr"`
	TileHeight int `xml:"tileheight,attr"`

	// attrs optional
	Version         string `xml:"version,attr,optional,omitempty"`
	TiledVersion    string `xml:"tiledversion,attr,optional,omitempty"`
	Orientation     string `xml:"orientation,attr,optional,omitempty"`
	RenderOrder     string `xml:"renderorder,attr,optional,omitempty"`
	HexSideLength   int    `xml:"hexsidelength,attr,optional,omitempty"`
	StaggerAxis     string `xml:"staggeraxis,attr,optional,omitempty"`
	StaggerIndex    string `xml:"staggerindex,attr,optional,omitempty"`
	BackgroundColor string `xml:"backgroundcolor,attr,optional,omitempty"`
	NextObjectId    int    `xml:"nextobjectid,attr,optional,omitempty"`

	// subsections
	Tilesets   []tileset   `xml:"tileset"`
	TileLayers []tileLayer `xml:"layer"`

	// subsections optional
	Properties   properties    `xml:"properties,optional,omitempty"`
	ObjectGroups []objectGroup `xml:"objectgroup,optional,omitempty"`
	ImageLayers  []imageLayer  `xml:"imagelayer,optional,omitempty"`
	Groups       []group       `xml:"group,optional,omitempty"`
}

// Inflate xml encoding struct into our common.Map struct.
//
func (m *tileMap) inflate() (*common.Map, error) {
	settings := []common.MapOption{}

	c, err := m.Background()
	if err != nil {
		return nil, err
	}
	settings = append(settings, common.Background(c))

	if m.Orientation == common.MapOrientationIsometric {
		settings = append(settings, common.OrientationIsometric())
	} else if m.Orientation == common.MapOrientationStaggered {
		settings = append(settings, common.OrientationStaggered(m.StaggerAxis, m.StaggerIndex))
	} else {
		settings = append(settings, common.OrientationOrthogonal())
	}

	if m.RenderOrder == common.MapRenderOrderRightDown {
		settings = append(settings, common.RenderOrderRightDown())
	} else if m.RenderOrder == common.MapRenderOrderLeftDown {
		settings = append(settings, common.RenderOrderLeftDown())
	} else if m.RenderOrder == common.MapRenderOrderRightUp {
		settings = append(settings, common.RenderOrderRightUp())
	} else if m.RenderOrder == common.MapRenderOrderLeftUp {
		settings = append(settings, common.RenderOrderLeftUp())
	}

	out := common.NewMap(settings...)
	out.Height = m.Height
	out.Width = m.Width
	out.TileWidth = m.TileWidth
	out.TileHeight = m.TileHeight

	out.UpdateProperties(m.Properties.inflate()...)

	for _, tileset := range m.Tilesets {
		tileset.inflate(out)
	}
	for _, layer := range m.TileLayers {
		layer.inflate(out)
	}
	for _, layer := range m.ImageLayers {
		layer.inflate(out)
	}
	return out, nil
}

// Save Map as tmx
//
func (m *tileMap) SaveTmx(in string) error {
	data, err := xml.Marshal(m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(in, data, 0644)
}

// Get map background colour
//
func (m *tileMap) Background() (*color.RGBA, error) {
	return decodeHexColour(m.BackgroundColor)
}

// Set map background colour
//
func (m *tileMap) SetBackground(in *color.RGBA) {
	m.BackgroundColor = encodeHexColour(in)
}
