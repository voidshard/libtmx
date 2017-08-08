package v1

import (
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"image/color"
	"strconv"
	"strings"
	"encoding/base64"
)

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Turn space separated string of x,y coords to []image.Point
func decodePoints(s string) ([]image.Point, error) {
	points := []image.Point{}
	for _, bit := range strings.Split(s, " ") {
		coords := strings.Split(bit, ",")
		if len(coords) != 2 {
			return nil, errors.New(fmt.Sprintf("Expected x,y coord from %s got %s", bit, coords))
		}

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return nil, err
		}
		points = append(points, image.Pt(x, y))
	}
	return points, nil
}

// Turn list of points to space separated string of x,y coords
func encodePoints(in []image.Point) string {
	bits := []string{}
	for _, p := range in {
		bits = append(bits, fmt.Sprintf("%d,%d", p.X, p.Y))
	}
	return strings.Join(bits, " ")
}

// Parse a colour from hex #AARRGGBB or #RRGGBB
func decodeHexColour(s string) (*color.RGBA, error) {
	data, err := hex.DecodeString(strings.TrimLeft(s, "#"))
	if err != nil {
		return nil, err
	}

	dlen := len(data)
	c := &color.RGBA{R: data[dlen-3], G: data[dlen-2], B: data[dlen-1]}
	if dlen > 3 {
		c.A = data[0]
	}
	return c, nil
}

// Turn a RGBA colour back into #RRGGBB format (Experimentation finds that #AARRGGBB doesn't work?)
func encodeHexColour(in *color.RGBA) string {
	chex := hex.EncodeToString([]byte{in.R, in.G, in.B})
	return fmt.Sprintf("#%s", strings.ToUpper(chex))
}

// Turn terrain Id csv to []int
//   - it is possible there is no terrain info (we get simply "")
//   - if there is terrain info, there will 4 segments separated by ','
//   - even if there is terrain info, some segments may be empty ("")
//   - note that I use '-1' to indicate that no data was found
//
//  Examples: ",,,2", "1,2,,2", "", "1,2,3,4"
//
func decodeTerrain(in string) ([4]int, error) {
	res := [4]int{-1, -1, -1, -1}
	if in == "" {
		return res, nil
	}

	bits := strings.SplitN(in, ",", 4)
	for i := 0; i < 4; i++ {
		if bits[i] == "" {
			continue
		}

		x, err := strconv.Atoi(bits[i])
		if err != nil {
			return res, nil
		}
		res[i] = x
	}
	return res, nil
}

// Inverse of decodeTerrain - turn int[4] into .tmx compatible csv terrain string
func encodeTerrain(in []int) string {
	if in[0] + in[1] + in[2] + in[3] == -4 {
		return "" // special case if no terrain is set
	}

	bits := []string{"", "", "", ""}
	for index, val := range in {
		x := ""
		if val > -1 {
			x = strconv.Itoa(val)
		}
		bits[index] = x
	}
	return strings.Join(bits, ",")
}

// Decode tile data into [][]int - assuming it's base64 encoded csv data
func decodeTileDataBase64(in string) ([][]int, error) {
	tmp, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return nil, err
	}
	return decodeTileDataCsv(string(tmp))
}

// Encode tile id row/column data back into csv / newline delimited format
func encodeTileDataCsv(tiles [][]int) string {
	lastRow := len(tiles) -1
	rows := []string{}
	for rowNum, tileIdRow := range tiles {
		row := []string{}
		for _, tileid := range tileIdRow {
			row = append(row, strconv.Itoa(tileid))
		}
		if rowNum == lastRow {
			rows = append(rows, strings.Join(row, ","))
		} else {
			rows = append(rows, fmt.Sprintf("%s,", strings.Join(row, ",")))
		}
	}
	return strings.Join(rows, "\n")
}

// Decode tile data into [][]int - assuming it's simply csv & newline delimited
func decodeTileDataCsv(csv string) ([][]int, error) {
	// It'd be more efficient to return a simple []int where tiles are placed via the usual id*y+x
	// but I think this way is easier to understand & modify.
	in := strings.Replace(csv, "\n", controlCharXml, -1) // remove differences between \n and xml char

	result := [][]int{}
	for _, row := range strings.Split(in, controlCharXml) {
		row_ids := []int{}
		if row == "" {
			continue
		}

		for _, column := range strings.Split(row, ",") {
			if column == "" {
				continue
			}

			value, err := strconv.Atoi(column)
			if err != nil {
				return nil, err
			}
			row_ids = append(row_ids, value)
		}
		result = append(result, row_ids)
	}
	return result, nil
}