package v1

import (
	"encoding/xml"
	"strconv"
	"github.com/voidshard/libtmx/common"
)

// Represents a list of Property objects
//
type properties struct {
	XMLName xml.Name `xml:"properties"`

	// subsections
	Properties []property `xml:"property,optional,omitempty"`
}

func (p *properties) inflate() []*common.Property {
	result := []*common.Property{}
	for _, prop := range p.Properties {
		result = append(result, prop.inflate())
	}
	return result
}

type property struct {
	XMLName xml.Name `xml:"property"`

	// attrs
	Name      string `xml:"name,attr"`
	ValueType string `xml:"type,attr"`
	Value     string `xml:"value,attr"`
}

func (p *property) inflate() *common.Property {
	prop := common.NewProp(p.Name)

	if p.ValueType == common.PropertyTypeString {
		prop.SetString(p.Value)
	} else if p.ValueType == common.PropertyTypeFile {
		prop.SetFilepath(p.Value)
	} else if p.ValueType == common.PropertyTypeColour {
		col, _ := decodeHexColour(p.Value)
		prop.SetColour(col)
	} else if p.ValueType == common.PropertyTypeFloat {
		val, _ := strconv.ParseFloat(p.Value, 64)
		prop.SetFloat(val)
	} else if p.ValueType == common.PropertyTypeInt {
		val, _ := strconv.Atoi(p.Value)
		prop.SetInt(val)
	} else if p.ValueType == common.PropertyTypeBool {
		prop.SetBool(p.Value == "true")
	}

	return prop
}

func deflateProperty(in *common.Property) (out property) {
	out.Name = in.Name()
	out.ValueType = in.Type()

	if out.ValueType == common.PropertyTypeString {
		out.Value = in.AsString()
	} else if out.ValueType == common.PropertyTypeFile {
		out.Value = in.AsFilepath()
	} else if out.ValueType == common.PropertyTypeFloat {
		out.Value = strconv.FormatFloat(in.AsFloat(), 'f', -1, 64)
	} else if out.ValueType == common.PropertyTypeInt {
		out.Value = strconv.Itoa(in.AsInt())
	} else if out.ValueType == common.PropertyTypeColour {
		out.Value = encodeHexColour(in.AsColour())
	} else if out.ValueType == common.PropertyTypeBool {
		out.Value = "false"
		if in.AsBool() {
			out.Value = "true"
		}
	}

	return
}

func deflateProperties(in []*common.Property) (out properties) {
	for _, prop := range in {
		out.Properties = append(out.Properties, deflateProperty(prop))
	}
	return
}