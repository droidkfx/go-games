package ttf

import (
	"io"
	"math"
)

var _ Table = (*Directory)(nil)

type Directory struct {
	head   DirectoryHeader
	tables map[TableType]DirectoryRow
}

func (d Directory) IsOpenTypeTable() bool {
	return false
}

func (d Directory) Type() TableType {
	return TableType_Directory
}

func (d Directory) SizeInFile() uint {
	return uint((sizeOfUint32 + (sizeOfUint16 * 4)) + (len(d.tables) * (sizeOfUint32 * 4)))
}

func (d Directory) nextTableTypeAfterOffset(offset uint32) (DirectoryRow, bool) {
	var nextTable DirectoryRow
	hasNext := false
	lowestOffset := uint32(math.MaxUint32)
	for _, dirRow := range d.tables {
		if dirRow.Offset < lowestOffset && dirRow.Offset > offset {
			hasNext = true
			nextTable = dirRow
			lowestOffset = dirRow.Offset
		}
	}
	return nextTable, hasNext
}

func parseDirectory(data io.Reader) any {
	head := parseDirectoryHead(data)
	tMap := map[TableType]DirectoryRow{}
	for i := uint16(0); i < head.NumTables; i++ {
		row := parseDirectoryRow(data)
		tMap[row.Table()] = row
	}
	return &Directory{
		head:   head,
		tables: tMap,
	}
}

type DirectoryHeader struct {
	// ScalarType A Tag to indicate the OFA scalar to be used to rasterize this font
	// The values 'true' (0x74727565) and 0x00010000 are recognized by OS X and iOS as referring to TrueType fonts.
	// The value 'typ1' (0x74797031) is recognized as referring to the old style of PostScript font housed in a
	// sfnt wrapper. The value 'OTTO' (0x4F54544F) indicates an OpenType font with PostScript outlines (that is, a
	// 'CFF ' table instead of a 'glyf' table). Other values are not currently supported.
	ScalarType uint32
	// NumTables the number of tables
	NumTables uint16
	// SearchRange (maximum power of 2 <= NumTables)*16
	SearchRange uint16
	// EntrySelector log2(maximum power of 2 <= NumTables)
	EntrySelector uint16
	// RangeShift NumTables*16-SearchRange
	RangeShift uint16
}

func parseDirectoryHead(data io.Reader) DirectoryHeader {
	return parse[DirectoryHeader](data)
}

type DirectoryRow struct {
	// Tag 4-byte identifier
	Tag uint32
	// CheckSum for this table
	CheckSum uint32
	// Offset from beginning of sfnt
	Offset uint32
	// Length of this table in byte (actual Length not padded Length)
	Length uint32
}

func (d DirectoryRow) Table() TableType {
	return TableType(uint32ToString(d.Tag))
}

func parseDirectoryRow(data io.Reader) DirectoryRow {
	return parse[DirectoryRow](data)
}
