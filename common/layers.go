package common

import (
	"image/color"
)

type ImageLayer struct {
	parent  *Map
	Name    string
	Opacity float64
	Visible bool
	OffsetX int
	OffsetY int
	ImageSource string
	ImageFormat string
	Width int
	Height int
	Format string
	TransparentColour *color.RGBA
	properties map[string]*Property
}

func (m *ImageLayer) UpdateProperties(props ...*Property) {
	for _, prop := range props {
		m.properties[prop.Name()] = prop
	}
}

func (m *ImageLayer) Property(name string) (*Property, bool) {
	prop, ok := m.properties[name]
	return prop, ok
}

func (m *ImageLayer) Properties() []*Property {
	results := []*Property{}
	for _, p := range m.properties {
		results = append(results, p)
	}
	return results
}

type TileLayer struct {
	parent     *Map
	tiles      []*Tile
	Name       string
	Opacity    float64
	Visible    bool
	OffsetX    int
	OffsetY    int
	properties map[string]*Property
}

func (t *TileLayer) Width() int {
	return t.parent.Width
}

func (t *TileLayer) Height() int {
	return t.parent.Height
}

func (t *TileLayer) UpdateProperties(props ...*Property) {
	for _, prop := range props {
		t.properties[prop.Name()] = prop
	}
}

func (t *TileLayer) Property(name string) (*Property, bool) {
	prop, ok := t.properties[name]
	return prop, ok
}

func (t *TileLayer) Properties() []*Property {
	results := []*Property{}
	for _, p := range t.properties {
		results = append(results, p)
	}
	return results
}

func (t *TileLayer) TileIds() [][]int {
	result := make([][]int, t.parent.Height)
	for y := 0; y < t.parent.Height; y++ {
		result[y] = make([]int, t.parent.Width)

		for x := 0; x < t.parent.Width; x++ {
			tile := t.Get(x, y)
			if tile == nil {
				result[y][x] = 0
				continue
			}
			result[y][x] = tile.GlobalID()
		}
	}
	return result
}

func (t *TileLayer) Get(x, y int) *Tile {
	if x < 0 || x >= t.parent.Width {
		return nil
	}
	if y < 0 || y >= t.parent.Height {
		return nil
	}
	return t.tiles[y * t.parent.Height + x]
}

func (t *TileLayer) Put(x, y int, tile *Tile) {
	if x < 0 || x >= t.parent.Width {
		return
	}
	if y < 0 || y >= t.parent.Width {
		return
	}
	t.tiles[y * t.parent.Height + x] = tile
}