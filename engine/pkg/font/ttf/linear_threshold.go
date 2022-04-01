package ttf

import (
	"io"
)

var _ Table = (*noopTable)(nil)

type LinearThreshold struct {
	// Version number (starts at 0).
	Version uint16
	// NumGlyphs Number of glyphs (from “numGlyphs” in 'maxp' table).
	NumGlyphs uint16
	// YPels The vertical pel height at which the glyph can be assumed to scale linearly.On a per glyph basis.
	YPels []uint8
}

func (n LinearThreshold) IsOpenTypeTable() bool { return true }
func (n LinearThreshold) Type() TableType       { return TableType_LinearThreshold }
func (n LinearThreshold) SizeInFile() uint {
	return sizeOfUint16 + sizeOfUint16 + sizeOfUint8*uint(len(n.YPels))
}

func parseLinearThreshold(data io.Reader) any {
	table := &LinearThreshold{
		Version:   ui16(data),
		NumGlyphs: ui16(data),
	}
	table.YPels = make([]uint8, 0, table.NumGlyphs)
	for i := uint16(0); i < table.NumGlyphs; i++ {
		table.YPels = append(table.YPels, ui8(data))
	}
	return table
}
