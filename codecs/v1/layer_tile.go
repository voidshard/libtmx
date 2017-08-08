package v1

import (
	"encoding/xml"
	"github.com/voidshard/libtmx/common"
)

var (
	// Control char for new lines used by XML
	controlCharXml = "&#xA;"
)

// Represents a layer of tiles (tile x,y -> id data being in Data.Value)
//
type tileLayer struct {
	XMLName xml.Name `xml:"layer"`

	// attrs
	Name string `xml:"name,attr"`

	// attrs optional
	//X       int     `xml:"x,attr,optional,omitempty"` // defaults to 0, cannot be changed
	//Y       int     `xml:"y,attr,optional,omitempty"` // defaults to 0, cannot be changed
	Visible int     `xml:"visible,attr,optional,omitempty"`
	OffsetX int     `xml:"offsetx,attr,optional,omitempty"`
	OffsetY int     `xml:"offsety,attr,optional,omitempty"`
	Width int `xml:"width,attr,optional,omitempty"`
	Height int `xml:"height,attr,optional,omitempty"`
	Opacity float64 `xml:"opacity,attr,optional,omitempty"`

	// subsections
	Data       dataBlock  `xml:"data,optional,omitempty"`
	Properties properties `xml:"properties,optional,omitempty"`
}

// Deflate the given common.TileLayer to be a codec v1 tileLayer for writing to xml
//
func deflateTileLayer(in *common.TileLayer) tileLayer {
	return tileLayer{
		Name: in.Name,
		Visible: boolToInt(in.Visible),
		OffsetX: in.OffsetX,
		OffsetY: in.OffsetY,
		Opacity: in.Opacity,
		Width: in.Width(),
		Height: in.Height(),
		Data: dataBlock{
			Value: encodeTileDataCsv(in.TileIds()),
			Encoding: common.DataEncodingCsv,
		},
		Properties: deflateProperties(in.Properties()),
	}
}

// Inflate this tileLayer from it's xml format to a common.TileLayer and add
// it to the given map.
//
func (t *tileLayer) inflate(parent *common.Map) error {
	layer := parent.NewTileLayer(t.Name)

	tileIds, err := t.TileIds()
	if err != nil {
		return err
	}

	layer.Opacity = t.Opacity
	layer.Visible = t.Visible == 1
	layer.OffsetX = t.OffsetX
	layer.OffsetX = t.OffsetY
	layer.UpdateProperties(t.Properties.inflate()...)

	for y, row := range tileIds {
		for x, tileId := range row { // Nb: these are global tile ids
			if tileId == 0 {  // tile id of 0 means no tile is there
				continue
			}

			tile, ok := globalTilemap[tileId]
			if !ok {
				continue
			}
			layer.Put(x, y, tile.inflatedTile)
		}
	}

	return nil
}

// Return the tile data as a slice of int slices representing a TileId at a given x,y.
//
func (t *tileLayer) TileIds() ([][]int, error) {
	var tiledata [][]int
	var err error

	if t.Data.Encoding == common.DataEncodingBase64 {
		tiledata, err = decodeTileDataBase64(t.Data.Value)
	} else {
		tiledata, err = decodeTileDataCsv(t.Data.Value)
	}

	if err != nil {
		return nil, err
	}

	return tiledata, err
}
