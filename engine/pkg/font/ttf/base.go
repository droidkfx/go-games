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
		dir: ParseTable[Directory](TableType_Directory, data),
	}
	fmt.Printf("%+v %v", font, font.dir.SizeInFile())
	return font, nil
}
