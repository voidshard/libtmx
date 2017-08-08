package common

const (
	// Available settings for various fields.
	//  See TMX format docs: http://doc.mapeditor.org/reference/tmx-map-format/
	MapOrientationOrthogonal = "orthogonal"
	MapOrientationIsometric  = "isometric"
	MapOrientationStaggered  = "staggered"

	MapRenderOrderRightDown = "right-down"
	MapRenderOrderRightUp   = "right-up"
	MapRenderOrderLeftDown  = "left-down"
	MapRenderOrderLeftUp    = "left-up"

	MapStaggerAxisX = "x"
	MapStaggerAxisY = "y"

	MapStaggerIndexEven = "even"
	MapStaggerIndexOdd  = "odd"

	DataEncodingBase64 = "base64"
	DataEncodingCsv    = "csv"

	DataCompressionGzip = "gzip"
	DataCompressionZlib = "zlib"

	ObjectGroupDrawOrderIndex   = "index"
	ObjectGroupDrawOrderTopDown = "topdown"

	TextHAlignLeft   = "left"
	TextHAlignRight  = "right"
	TextHAlignCentre = "center"

	TextVAlignTop    = "top"
	TextVAlignBottom = "bottom"
	TextVAlignCentre = "center"

	PropertyTypeString = "string"
	PropertyTypeInt    = "int"
	PropertyTypeFloat  = "float"
	PropertyTypeBool   = "bool"
	PropertyTypeColour = "color"
	PropertyTypeFile   = "file"

	ObjectTypeEllipse = "e"
	ObjectTypePolygon = "p"
	ObjectTypePolyline = "l"
	ObjectTypeText = "t"

	// Defaults that shall be enforced on object creation or parsing
	DefaultPropertyType = PropertyTypeString
	DefaultDrawOrder    = ObjectGroupDrawOrderTopDown
	DefaultRenderOrder  = MapRenderOrderRightDown
	DefaultVAlign       = TextVAlignTop
	DefaultHAlign       = TextHAlignLeft
	DefaultStaggerAxis  = MapStaggerAxisX
	DefaultStaggerIndex = MapStaggerIndexEven
	DefaultFontFamily   = "sand-serif"
	DefaultPixelSize    = 16
	DefaultTmxVersion   = "1.0"
)
