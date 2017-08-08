package common


type Tileset struct {
	parent *Map

	FirstGID int // the global id (across all tilesets) for the first tile in this set
	Name string
	OffsetX int
	OffsetY int
	TileWidth       int
	TileHeight      int
	Spacing int
	Margin int

	terrain []*Terrain
	properties  map[string]*Property
	tiles   []*Tile
}

func (t *Tileset) Tiles() []*Tile {
	return t.tiles
}

func (t *Tileset) AddTerrain(in ...*Terrain) {
	for _, ter := range in {
		ter.Id = len(t.terrain)
		t.terrain = append(t.terrain, ter)
	}
}

func (t *Tileset) Terrain() []*Terrain {
	return t.terrain
}

func (t *Tileset) UpdateProperties(props ...*Property) {
	for _, prop := range props {
		t.properties[prop.Name()] = prop
	}
}

func (t *Tileset) Property(name string) (*Property, bool) {
	prop, ok := t.properties[name]
	return prop, ok
}

func (t *Tileset) Properties() []*Property {
	results := []*Property{}
	for _, p := range t.properties {
		results = append(results, p)
	}
	return results
}

func (t *Tileset) TileCount() int {
	return len(t.tiles)
}

func (t *Tileset) AddTiles(tiles ...*Tile) {
	for _, tile := range tiles {
		if tile.Source == "" {
			continue
		}

		tile.parent = t
		t.tiles = append(t.tiles, tile)
	}
}

func (m *Map) NewTileset(name string, tiles ...*Tile) *Tileset {
	tset := &Tileset{
		parent: m,
		Name: name,
		properties: make(map[string]*Property),
		terrain: []*Terrain{},
		tiles: []*Tile{},
		TileWidth: m.TileWidth,
		TileHeight: m.TileHeight,
	}

	m.tilesets = append(m.tilesets, tset)
	tset.AddTiles(tiles...)
	return tset
}

type Tile struct {
	parent      *Tileset
	terrain     []*Terrain
	properties  map[string]*Property

	Id int
	Type    string
	Source string
	Width int
	Height int
	Probability float64
	Animation *Animation
}

func NewTile(source string) *Tile {
	return &Tile{
		Source: source,
		properties: make(map[string]*Property),
		terrain: make([]*Terrain, 4),
	}
}

func (t *Tile) GlobalID() int {
	return t.Id + t.parent.FirstGID
}

func (t *Tile) TopLeftTerrain() *Terrain {
	return t.terrain[0]
}

func (t *Tile) TopRightTerrain() *Terrain {
	return t.terrain[1]
}

func (t *Tile) BottomLeftTerrain() *Terrain {
	return t.terrain[2]
}

func (t *Tile) BottomRightTerrain() *Terrain {
	return t.terrain[3]
}

func (t *Tile) SetTopLeftTerrain(in *Terrain) {
	t.terrain[0] = in
}

func (t *Tile) SetTopRightTerrain(in *Terrain) {
	t.terrain[1] = in
}

func (t *Tile) SetBottomLeftTerrain(in *Terrain) {
	t.terrain[2] = in
}

func (t *Tile) SetBottomRightTerrain(in *Terrain) {
	t.terrain[3] = in
}

func (t *Tile) UpdateProperties(props ...*Property) {
	for _, prop := range props {
		t.properties[prop.Name()] = prop
	}
}

func (t *Tile) Terrain() []*Terrain {
	return t.terrain
}

func (t *Tile) Property(name string) (*Property, bool) {
	prop, ok := t.properties[name]
	return prop, ok
}

func (t *Tile) Properties() []*Property {
	results := []*Property{}
	for _, p := range t.properties {
		results = append(results, p)
	}
	return results
}

func (t *Tile) SetAnimation(frames ...*Frame) {
	animation := newAnimation(frames...)
	t.Animation = animation
	animation.parent = t
}

type Animation struct {
	parent *Tile
	Frames []*Frame
}

type Frame struct {
	parent *Animation
	Tile *Tile
	Duration int
}

func newAnimation(frames ...*Frame) *Animation {
	ani := &Animation{Frames: []*Frame{}}

	for _, fr := range frames {
		fr.parent = ani
		ani.Frames = append(ani.Frames, fr)
	}

	return ani
}

type Terrain struct {
	Id int
	Name       string
	Tile       *Tile
	properties map[string]*Property
}

func NewTerrain(name string) *Terrain {
	return &Terrain{
		Name: name,
		properties: make(map[string]*Property),
	}
}

func (t *Terrain) UpdateProperties(props ...*Property) {
	for _, prop := range props {
		t.properties[prop.Name()] = prop
	}
}

func (t *Terrain) Property(name string) (*Property, bool) {
	prop, ok := t.properties[name]
	return prop, ok
}

func (t *Terrain) Properties() []*Property {
	results := []*Property{}
	for _, p := range t.properties {
		results = append(results, p)
	}
	return results
}
