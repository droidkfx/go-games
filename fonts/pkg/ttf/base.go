package ttf

import (
	"errors"
	"fmt"
	"io"
)

type TrueTypeFont struct {
	dir    *Directory
	tables map[TableType]Table
}

func ParseFont(data io.Reader) (*TrueTypeFont, error) {
	dir, parseDirErr := ParseTable[*Directory](TableType_Directory, data)
	if parseDirErr != nil {
		return nil, parseDirErr
	}
	font := &TrueTypeFont{
		dir:    dir,
		tables: map[TableType]Table{},
	}
	currentOffset := uint32(0)
	for i := uint16(0); i < font.dir.head.NumTables; i++ {
		nextType, hasNext := font.dir.nextTableTypeAfterOffset(currentOffset)
		if !hasNext {
			return nil, errors.New(fmt.Sprintf("expected %d tables but only found %d", font.dir.head.NumTables, i))
		}
		newTable, parseTableErr := ParseTable[Table](nextType.Table(), data)
		if parseTableErr != nil {
			return font, parseTableErr
		}
		font.tables[nextType.Table()] = newTable
		currentOffset = nextType.Offset
	}
	return font, nil
}
