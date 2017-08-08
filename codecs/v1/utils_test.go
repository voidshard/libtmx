package v1

import (
	"encoding/base64"
	"image"
	"image/color"
	"reflect"
	"testing"
)

const (
	encodedcsv = "&#xA;0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,&#xA;0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,&#xA;0,0,0,0,0,140,140,140,140,140,140,140,140,140,140,140,0,0,0,0,&#xA;0,0,0,0,0,140,140,140,140,140,140,140,140,140,140,140,0,0,0,0,&#xA;0,0,0,0,0,140,140,140,140,140,140,140,140,140,140,140,0,0,0,0,&#xA;0,0,0,0,0,140,140,140,140,140,140,140,140,140,140,140,0,0,0,0,&#xA;0,0,0,0,0,140,140,140,140,140,140,140,140,140,140,140,0,0,0,0,&#xA;0,0,0,0,0,140,140,140,140,140,140,140,140,140,140,140,0,0,0,0,&#xA;0,0,0,0,0,140,140,140,140,140,140,140,140,140,140,140,0,0,0,0,&#xA;0,0,0,0,0,140,140,140,140,140,140,140,140,140,140,140,0,0,0,0,&#xA;0,0,0,0,0,140,140,140,140,140,140,140,140,140,140,140,0,0,0,0,&#xA;0,0,0,0,0,140,140,140,140,140,140,140,140,140,140,140,0,0,0,0,&#xA;0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,&#xA;0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,&#xA;0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,&#xA;0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,&#xA;0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,&#xA;0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,&#xA;0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,&#xA;0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0&#xA;"
)

var (
	decodedcsv = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
)

func TestDecodePoints(t *testing.T) {
	cases := []struct {
		In     string
		Expect []image.Point
	}{
		{
			"0,0 146,-13 164,165 58,193 -80,152 117,80 1,2",
			[]image.Point{
				image.Pt(0, 0),
				image.Pt(146, -13),
				image.Pt(164, 165),
				image.Pt(58, 193),
				image.Pt(-80, 152),
				image.Pt(117, 80),
				image.Pt(1, 2),
			},
		},
	}

	for _, test := range cases {
		result, err := decodePoints(test.In)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, test.Expect) {
			t.Error("Given", test.In, "expected", test.Expect, "got", result)
		}
	}
}

func TestEncodePoints(t *testing.T) {
	cases := []struct {
		Expect string
		In     []image.Point
	}{
		{
			"0,0 146,-13 164,165 58,193 -80,152 117,80 1,2",
			[]image.Point{
				image.Pt(0, 0),
				image.Pt(146, -13),
				image.Pt(164, 165),
				image.Pt(58, 193),
				image.Pt(-80, 152),
				image.Pt(117, 80),
				image.Pt(1, 2),
			},
		},
	}

	for _, test := range cases {
		result := encodePoints(test.In)

		if !reflect.DeepEqual(result, test.Expect) {
			t.Error("Given", test.In, "expected", test.Expect, "got", result)
		}
	}
}

func TestDecodeTerrain(t *testing.T) {
	cases := []struct {
		In     string
		Expect [4]int
	}{
		{"", [4]int{-1, -1, -1, -1}},
		{",,,", [4]int{-1, -1, -1, -1}},
		{"0,1,2,3", [4]int{0, 1, 2, 3}},
		{"3,,,2", [4]int{3, -1, -1, 2}},
	}

	for _, test := range cases {
		result, err := decodeTerrain(test.In)
		if err != nil {
			t.Error(err)
		}

		if result != test.Expect {
			t.Error("Given", test.In, "expected", test.Expect, "got", result)
		}
	}
}

func TestEncodeTerrain(t *testing.T) {
	cases := []struct {
		Expect string
		In     []int
	}{
		{"", []int{-1, -1, -1, -1}},
		{"0,1,2,3", []int{0, 1, 2, 3}},
		{"3,,,2", []int{3, -1, -1, 2}},
	}

	for _, test := range cases {
		result := encodeTerrain(test.In)
		if result != test.Expect {
			t.Error("Given", test.In, "expected", test.Expect, "got", result)
		}
	}
}

func TestDecodeHexColour(t *testing.T) {
	cases := []struct {
		In     string
		Expect *color.RGBA
	}{
		{"#FF00FF00", &color.RGBA{0, 255, 0, 255}},
		{"FF00FF00", &color.RGBA{0, 255, 0, 255}},
		{"#00FF00FF", &color.RGBA{255, 0, 255, 0}},
		{"00FF00FF", &color.RGBA{255, 0, 255, 0}},
		{"FF00FF", &color.RGBA{255, 0, 255, 0}},
		{"FF00FF", &color.RGBA{255, 0, 255, 0}},
		{"#00FF00", &color.RGBA{0, 255, 0, 0}},
		{"00FF00", &color.RGBA{0, 255, 0, 0}},
	}

	for _, test := range cases {
		result, err := decodeHexColour(test.In)
		if err != nil {
			t.Error(err)
		}

		if result.A != test.Expect.A || result.R != test.Expect.R || result.G != test.Expect.G || result.B != test.Expect.B {
			t.Error("Given", test.In, "expected", test.Expect, "got", result)
		}
	}
}

func TestEncodeHexColour(t *testing.T) {
	cases := []struct {
		Expect string
		In     *color.RGBA
	}{
		{"#00FF00", &color.RGBA{0, 255, 0, 255}},
		{"#FF00FF", &color.RGBA{255, 0, 255, 0}},
		{"#00FF00", &color.RGBA{0, 255, 0, 0}},
	}

	for _, test := range cases {
		result := encodeHexColour(test.In)
		if result != test.Expect {
			t.Error("Given", test.In, "expected", test.Expect, "got", result)
		}
	}
}

func TestDecodeTileDataBase64(t *testing.T) {
	cases := []struct {
		In     string
		Expect [][]int
	}{
		{base64.StdEncoding.EncodeToString([]byte(encodedcsv)), decodedcsv},
	}

	for _, test := range cases {
		result, err := decodeTileDataBase64(test.In)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, test.Expect) {
			t.Error("Given", test.In, "expected", test.Expect, "got", result)
		}
	}
}

func TestDecodeTileDataCsv(t *testing.T) {
	cases := []struct {
		In     string
		Expect [][]int
	}{
		{encodedcsv, decodedcsv},
	}

	for _, test := range cases {
		result, err := decodeTileDataCsv(test.In)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, test.Expect) {
			t.Error("Given", test.In, "expected", test.Expect, "got", result)
		}
	}
}
