package main

import (
	"fmt"
	"os"

	"github.com/droidkfx/go-games/engine/pkg/font/ttf"
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

	_, err = ttf.ParseFont(fontBytes)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
