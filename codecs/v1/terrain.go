package v1

import (
	"encoding/xml"
	"github.com/voidshard/libtmx/common"
)

type terrainTypes struct {
	XMLName xml.Name `xml:"terraintypes"`

	// subsections
	Terrain []terrain `xml:"terrain"`
}

func (t *terrainTypes) inflate() []*common.Terrain {
	result := []*common.Terrain{}
	for terrainId, ter := range t.Terrain {
		terrain := common.NewTerrain(ter.Name)
		globalTerrain[terrainId] = terrain

		if ter.Tile > 0 {
			tile, ok := globalTilemap[ter.Tile]
			if ok {
				terrain.Tile = tile.inflatedTile
			}
		}

		terrain.UpdateProperties(ter.Properties.inflate()...)
		result = append(result, terrain)
	}
	return result
}

type terrain struct {
	XMLName xml.Name `xml:"terrain"`

	// attrs
	Name string `xml:"name,attr"`
	Tile int    `xml:"tile,attr,optional,omitempty"`

	// subsections
	Properties properties `xml:"properties,optional,omitempty"`
}
