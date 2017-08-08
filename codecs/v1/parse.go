package v1

import (
	"github.com/voidshard/libtmx/common"
	"encoding/xml"
)

type CodecV1 struct {}

// Unmarshal a common.Map from the given data (that is, []byte read from a tmx .xml file)
//
func (c *CodecV1) Unmarshal(data []byte) (*common.Map, error) {
	var xmlmap tileMap
	err := xml.Unmarshal(data, &xmlmap)
	if err != nil {
		return nil, err
	}
	return xmlmap.inflate()
}

// Given a map, marshal it back into it's tmx .xml format, compatible with Tiled.
//
func (c *CodecV1) Marshal(in *common.Map) ([]byte, error) {
	in.FinalizeIDs()

	tmap := &tileMap{
		Height: in.Height,
		Width: in.Width,
		TileHeight: in.TileHeight,
		TileWidth: in.TileWidth,
		Version: in.Version,
		TiledVersion: in.TiledVersion,
		Orientation: in.Orientation(),
		RenderOrder: in.RenderOrder(),
		HexSideLength: in.HexSideLength,
		StaggerAxis: in.StaggerAxis(),
		StaggerIndex: in.StaggerIndex(),
		BackgroundColor: encodeHexColour(in.BackgroundColor),
		NextObjectId: 0,
		Properties: deflateProperties(in.Properties()),
		Tilesets:   []tileset{},
		TileLayers: []tileLayer{},
		ImageLayers:  []imageLayer{},
	}

	for _, tset := range in.Tilesets() {
		tmap.Tilesets = append(tmap.Tilesets, deflateTileset(tset))
	}

	for _, layer := range in.TileLayers() {
		tmap.TileLayers = append(tmap.TileLayers, deflateTileLayer(layer))
	}

	for _, layer := range in.ImageLayers() {
		tmap.ImageLayers = append(tmap.ImageLayers, deflateImageLayer(layer))
	}

	return xml.Marshal(tmap)
}
