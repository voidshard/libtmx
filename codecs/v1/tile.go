package v1

import (
	"encoding/xml"
	"github.com/voidshard/libtmx/common"
)

type tile struct {
	XMLName xml.Name `xml:"tile"`

	// attrs
	Id          int     `xml:"id,attr"`
	Type        string  `xml:"type,attr,optional,omitempty"`
	RawTerrain  string  `xml:"terrain,attr,optional,omitempty"`
	Probability float64 `xml:"probability,attr,optional,omitempty"`

	// subsections
	Properties  properties  `xml:"properties,optional,omitempty"`
	Image       imageData   `xml:"image,optional,omitempty"`
	Animation   animation   `xml:"animation,optional,omitempty"`
	ObjectGroup objectGroup `xml:"objectgroup,optional,omitempty"`

	// Wrapped common.Tile that represents this xml parsed Tile
	// (we have to create this in bits as Terrain & other tiles are loaded)
	inflatedTile *common.Tile
}

func (t *tile) inflate() {
	// Continue to setup Tile obj
	t.inflatedTile.UpdateProperties(t.Properties.inflate()...)
	t.inflatedTile.Source = t.Image.Source
	t.inflatedTile.Type = t.Type
	t.inflatedTile.Probability = t.Probability
	t.inflatedTile.Width = t.Image.Width
	t.inflatedTile.Height = t.Image.Height

	terrains, err := t.Terrain()
	if err == nil {
		if terrains[0] > -1 {
			ter, ok := globalTerrain[terrains[0]]
			if ok {
				t.inflatedTile.SetTopLeftTerrain(ter)
			}
		}

		if terrains[1] > -1 {
			ter, ok := globalTerrain[terrains[1]]
			if ok {
				t.inflatedTile.SetTopRightTerrain(ter)
			}
		}

		if terrains[2] > -1 {
			ter, ok := globalTerrain[terrains[2]]
			if ok {
				t.inflatedTile.SetBottomLeftTerrain(ter)
			}
		}

		if terrains[3] > -1 {
			ter, ok := globalTerrain[terrains[3]]
			if ok {
				t.inflatedTile.SetBottomRightTerrain(ter)
			}
		}
	}

	frames := []*common.Frame{}
	for _, fr := range t.Animation.Frames {
		tile, ok := globalTilemap[fr.TileId]
		if !ok {
			continue
		}

		frames = append(frames, &common.Frame{
			Duration: fr.Duration,
			Tile: tile.inflatedTile,
		})
	}
	t.inflatedTile.SetAnimation(frames...)
}

func deflateTile(in *common.Tile) tile {
	rawter := make([]int, 4)
	for i, terr := range in.Terrain() {
		if terr == nil {
			rawter[i] = -1
		} else {
			rawter[i] = terr.Id
		}
	}

	return tile{
		Id: in.Id,
		Type: in.Type,
		RawTerrain: encodeTerrain(rawter),
		Properties: deflateProperties(in.Properties()),
		Animation: deflateAnimation(in.Animation),
		Probability: in.Probability,
		Image: imageData{
			Source: in.Source,
			Width: in.Width,
			Height: in.Height,
		},
	}
}

func (t *tile) Terrain() ([4]int, error) {
	return decodeTerrain(t.RawTerrain)
}

func (t *tile) SetTerrain(in []int) {
	t.RawTerrain = encodeTerrain(in)
}

type animation struct {
	XMLName xml.Name `xml:"animation"`
	Frames   []frame  `xml:"frame"`
}

type frame struct {
	XMLName xml.Name `xml:"frame"`

	// attrs
	TileId   int `xml:"tileid,attr"`
	Duration int `xml:"duration,attr"`
}

func deflateAnimation(in *common.Animation) animation {
	ani := animation{
		Frames: []frame{},
	}
	for _, fr := range in.Frames {
		ani.Frames = append(ani.Frames, frame{
			TileId: fr.Tile.Id,
			Duration: fr.Duration,
		})
	}
	return ani
}

func deflateTerrain(in *common.Terrain) terrain {
	tileid := -1
	if in.Tile != nil {
		tileid = in.Tile.GlobalID()
	}

	return terrain{
		Name: in.Name,
		Tile: tileid,
		Properties: deflateProperties(in.Properties()),
	}
}