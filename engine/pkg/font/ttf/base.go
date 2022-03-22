package ttf

import (
	"fmt"
	"io"
)

type TrueTypeFont struct {
	dir Directory
}

func ParseFont(data io.Reader) (*TrueTypeFont, error) {
	font := &TrueTypeFont{
		dir: parseDirectory(data),
	}
	fmt.Printf("%+v %v", font, font.dir.sizeInFile())
	return font, nil
}
