package v1

import (
	"encoding/xml"
	"github.com/voidshard/libtmx/common"
)

var globalTilemap map[int]*tile
var globalTerrain map[int]*common.Terrain
func init () {
	globalTilemap = make(map[int]*tile)
	globalTerrain = make(map[int]*common.Terrain)
}

type tileset struct {
	XMLName xml.Name `xml:"tileset"`

	// attrs
	Name string `xml:"name,attr"`

	// attrs optional
	FirstGID   int    `xml:"firstgid,attr,optional,omitempty"`
	Source     string `xml:"source,attr,optional,omitempty"`
	TileWidth  int    `xml:"tilewidth,attr,optional,omitempty"`
	TileHeight int    `xml:"tileheight,attr,optional,omitempty"`
	Spacing    int    `xml:"spacing,attr,optional,omitempty"`
	Margin     int    `xml:"margin,attr,optional,omitempty"`
	Tilecount  int    `xml:"tilecount,attr,optional,omitempty"`
	Columns    int    `xml:"columns,attr,optional,omitempty"`

	// subsections
	Tiles []tile `xml:"tile"`

	// subsections optional
	Offset     tileOffset   `xml:"tileoffset,optional,omitempty"`
	Properties properties   `xml:"properties,optional,omitempty"`
	Terrain    terrainTypes `xml:"terraintypes,optional,omitempty"`
}

func deflateTileset(in *common.Tileset) tileset {
	tset := tileset{
		Name: in.Name,
		FirstGID: in.FirstGID,
		TileWidth: in.TileWidth,
		TileHeight: in.TileHeight,
		Spacing: in.Spacing,
		Margin: in.Margin,
		Tilecount: len(in.Tiles()),
		Offset: tileOffset{
			X: in.OffsetX,
			Y: in.OffsetY,
		},
		Properties: deflateProperties(in.Properties()),
		Tiles: []tile{},
		Terrain: terrainTypes{
			Terrain: []terrain{},
		},
	}

	for _, terr := range in.Terrain() {
		if terr != nil {
			tset.Terrain.Terrain = append(tset.Terrain.Terrain, deflateTerrain(terr))
		}
	}

	for _, tile := range in.Tiles() {
		tset.Tiles = append(tset.Tiles, deflateTile(tile))
	}

	return tset
}

func (t *tileset) inflate(parent *common.Map) *common.Tileset {
	// Build a map Id->Tile as we're going to need to match TileId(s) to Tiles
	// even when fully inflating Tiles & Terrain ..
	// That is, Tile and Terrain can reference other Tile(s)
	tiles := []*common.Tile{}
	tilewrappers := []*tile{}

	for _, tmp := range t.Tiles {
		tilecopy := &tile{
			Id: tmp.Id,
			Type: tmp.Type,
			RawTerrain: tmp.RawTerrain,
			Probability: tmp.Probability,
			Properties: tmp.Properties,
			Image: tmp.Image,
			Animation: tmp.Animation,
			ObjectGroup: tmp.ObjectGroup,
		}

		tileId := tilecopy.Id + t.FirstGID
		globalTilemap[tileId] = tilecopy
		tilecopy.inflatedTile = common.NewTile(tilecopy.Image.Source)
		tiles = append(tiles, tilecopy.inflatedTile)
		tilewrappers = append(tilewrappers, tilecopy)
	}

	obj := parent.NewTileset(t.Name, tiles...)
	obj.UpdateProperties(t.Properties.inflate()...)
	obj.AddTerrain(t.Terrain.inflate()...)

	obj.TileWidth = t.TileWidth
	obj.TileHeight = t.TileHeight
	obj.Margin = t.Margin
	obj.Spacing = t.Spacing
	obj.OffsetX = t.Offset.X
	obj.OffsetY = t.Offset.Y

	for _, tw := range tilewrappers {
		tw.inflate()
	}

	return obj
}
