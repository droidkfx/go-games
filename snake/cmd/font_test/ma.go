package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/droidkfx/go-games/fonts/pkg/ttf"
)

func main() {
	fontBytes, err := os.Open("C:\\src\\games\\snake\\assets\\fonts\\MontereyFLF.ttf")
	defer func(fontBytes *os.File) {
		closeErr := fontBytes.Close()
		if closeErr != nil {
			fmt.Printf("%+v\n", err)
		}
	}(fontBytes)
	if err != nil {
		panic(err)
	}

	font, err := ttf.ParseFont(fontBytes)
	fmt.Printf("ParseErr: %s\n\n", err)
	spew.Dump(font)
}
