package main

import (
	"fmt"
	"github.com/voidshard/libtmx"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"io/ioutil"
)

var (
	inFile = kingpin.Arg("file", "Input .tmx file").String()
)

// Read in tmx file, print out variables
//
func main() {
	kingpin.Parse()

	f, err := os.Open(*inFile)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	parser := libtmx.CodecV1()
	xmap, err := parser.Unmarshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Map:", xmap.Width, xmap.Height, xmap.TileWidth, xmap.TileHeight, xmap.BackgroundColor)
	for _, prop := range xmap.Properties() {
		fmt.Println(" ", prop.Name(), prop.Type())
	}

	for _, tileset := range xmap.Tilesets() {
		fmt.Println("> Tileset:", tileset.Name, "Tiles:", len(tileset.Tiles()))

		for _, prop := range tileset.Properties() {
			fmt.Println("   ", prop.Name(), prop.Type())
		}

		for _, terrain := range tileset.Terrain() {
			fmt.Println("Terrain:", terrain)
			for _, prop := range terrain.Properties() {
				fmt.Println("      ", prop.Name(), prop.Type())
			}

		}
	}
}
