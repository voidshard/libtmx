package common

import (
	"image/color"
)

type Property struct {
	name      string
	valueType string

	// Rather keep only "value" then convert & cast between types,
	// we keep a slot for each of the data types and provide an "As()" func to
	// return the appropriate type.
	// ToDo: Improve .. !
	valueString     string
	valueInt int
	valueBool bool
	valueColour *color.RGBA
	valueFloat float64
}

func (p *Property) Name() string {
	return p.name
}

func (p *Property) Type() string {
	return p.valueType
}

func NewProp(name string) *Property {
	return &Property{
		name: name,
		valueType: DefaultPropertyType,
	}
}

func (p *Property) clear() {
	p.valueString = ""
	p.valueInt = 0
	p.valueBool = false
	p.valueColour = &color.RGBA{0, 0, 0, 0}
	p.valueFloat = 0
}

func (p *Property) AsInt() int {
	return p.valueInt
}

func (p *Property) AsFloat() float64 {
	return p.valueFloat
}

func (p *Property) AsColour() *color.RGBA{
	return p.valueColour
}

func (p *Property) AsBool() bool {
	return p.valueBool
}

func (p *Property) AsString() string {
	return p.valueString
}

func (p *Property) AsFilepath() string {
	return p.valueString
}

func (p *Property) SetFilepath(in string) *Property {
	p.clear()
	p.valueString = in
	p.valueType = PropertyTypeFile
	return p
}

func (p *Property) SetString (in string) *Property {
	p.clear()
	p.valueString = in
	p.valueType = PropertyTypeString
	return p
}

func (p *Property) SetFloat(in float64) *Property {
	p.clear()
	p.valueFloat = in
	p.valueType = PropertyTypeFloat
	return p
}

func (p *Property) SetColour(in *color.RGBA) *Property {
	p.clear()
	p.valueColour = in
	p.valueType = PropertyTypeColour
	return p
}

func (p *Property) SetInt(in int) *Property {
	p.clear()
	p.valueInt = in
	p.valueType = PropertyTypeInt
	return p
}

func (p *Property) SetBool(in bool) *Property {
	p.clear()
	p.valueBool = in
	p.valueType = PropertyTypeBool
	return p
}
