package common

import (
	"image/color"
)

var (
	defaultMapOpts = []MapOption{
		Width(20),
		Height(20),
		TileWidth(32),
		TileHeight(32),
		OrientationOrthogonal(),
	}
)

type Map struct {
	Width           int
	Height          int
	TileWidth       int
	TileHeight      int
	HexSideLength   int
	BackgroundColor *color.RGBA
	Version      string
	TiledVersion string

	nextObjectId int
	orientation  string
	renderOrder  string
	staggerAxis  string
	staggerIndex string

	tilesets    []*Tileset
	tileLayers  []*TileLayer
	imageLayers []*ImageLayer
	properties  map[string]*Property
}

// Set Ids on child Tiles, Tilesets & Terrain (called before Map is written out)
func (m *Map) FinalizeIDs() {
	globaltilecount := 0

	for _, tileset := range m.tilesets {
		tileset.FirstGID = globaltilecount + 1 // Note that global tile id 0 is reserved for 'no tile'

		terrainCount := 0
		for _, terrain := range tileset.terrain {
			terrain.Id = terrainCount
			terrainCount += 1
		}

		localtilecount := 0
		for _, tile := range tileset.tiles {
			tile.Id = localtilecount
			localtilecount += 1
		}
		globaltilecount += localtilecount
	}
}

func (m *Map) Orientation() string {
	return m.orientation
}

func (m *Map) RenderOrder() string {
	return m.renderOrder
}

func (m *Map) StaggerAxis() string {
	return m.staggerAxis
}

func (m *Map) StaggerIndex() string {
	return m.staggerIndex
}

func (m *Map) Tilesets() []*Tileset {
	return m.tilesets
}

func (m *Map) Property(name string) (*Property, bool) {
	prop, ok := m.properties[name]
	return prop, ok
}

func (m *Map) Properties() []*Property {
	results := []*Property{}
	for _, p := range m.properties {
		results = append(results, p)
	}
	return results
}

func (m *Map) TileLayers() []*TileLayer {
	return m.tileLayers
}

func (m *Map) ImageLayers() []*ImageLayer {
	return m.imageLayers
}

func (m *Map) UpdateProperties(props ...*Property) {
	for _, prop := range props {
		m.properties[prop.Name()] = prop
	}
}

func (m *Map) NewTileLayer(name string) *TileLayer {
	tiles := make([]*Tile, m.Height * m.Height + m.Width)
	layer := &TileLayer{
		parent: m,
		Name: name,
		tiles: tiles,
		Visible: true,
		properties: make(map[string]*Property),
	}
	m.tileLayers = append(m.tileLayers, layer)
	return layer
}

func (m *Map) NewImageLayer(name, imageSource string) *ImageLayer {
	layer := &ImageLayer{
		parent: m,
		Name: name,
		Visible: true,
		ImageSource: imageSource,
		properties: make(map[string]*Property),
	}
	m.imageLayers = append(m.imageLayers, layer)
	return layer
}

func NewMap(opts ...MapOption) *Map {
	mp := &Map{
		tilesets: []*Tileset{},
		tileLayers: []*TileLayer{},
		imageLayers: []*ImageLayer{},
		properties: make(map[string]*Property),
		Version: DefaultTmxVersion,
	}
	for _, opt := range defaultMapOpts {
		opt(mp) // apply our defaults
	}
	for _, opt := range opts {
		opt(mp) // apply given settings
	}
	return mp
}