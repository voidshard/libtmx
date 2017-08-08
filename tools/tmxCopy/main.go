package main

import (
	"github.com/voidshard/libtmx"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"io/ioutil"
	"fmt"
)

var (
	inFile = kingpin.Arg("file", "Input .tmx file").String()
	outFile = kingpin.Arg("output", "Output .tmx file").String()
)

// Test of our marshal / unmarshal functions
//  - read in tmx file, parse to internal format & write out copy.
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

	fmt.Println(xmap.BackgroundColor)

	rawdata, err := parser.Marshal(xmap)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*outFile, rawdata, 0644)
	if err != nil {
		panic(err)
	}
}
