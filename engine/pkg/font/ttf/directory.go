package ttf

import (
	"io"
)

type Table string

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage,GoNameStartsWithPackageName
const (
	TUFTTable_AccentAttachment Table = "acnt"
	TTFTable_ankr              Table = "ankr"
	TTFTable_avar              Table = "avar"
	TTFTable_bdat              Table = "bdat"
	TTFTable_bhed              Table = "bhed"
	TTFTable_bloc              Table = "bloc"
	TTFTable_bsln              Table = "bsln"
	TTFTable_cmap              Table = "cmap"
	TTFTable_cvar              Table = "cvar"
	TTFTable_cvt               Table = "cvt "
	TTFTable_EBSC              Table = "EBSC"
	TTFTable_fdsc              Table = "fdsc"
	TTFTable_feat              Table = "feat"
	TTFTable_fmtx              Table = "fmtx"
	TTFTable_fond              Table = "fond"
	TTFTable_fpgm              Table = "fpgm"
	TTFTable_fvar              Table = "fvar"
	TTFTable_gasp              Table = "gasp"
	TTFTable_gcid              Table = "gcid"
	TTFTable_glyf              Table = "glyf"
	TTFTable_gvar              Table = "gvar"
	TTFTable_hdmx              Table = "hdmx"
	TTFTable_head              Table = "head"
	TTFTable_hhea              Table = "hhea"
	TTFTable_hmtx              Table = "hmtx"
	TTFTable_just              Table = "just"
	TTFTable_kern              Table = "kern"
	TTFTable_kerx              Table = "kerx"
	TTFTable_lcar              Table = "lcar"
	TTFTable_loca              Table = "loca"
	TTFTable_ltag              Table = "ltag"
	TTFTable_maxp              Table = "maxp"
	TTFTable_meta              Table = "meta"
	TTFTable_mort              Table = "mort"
	TTFTable_morx              Table = "morx"
	TTFTable_name              Table = "name"
	TTFTable_opbd              Table = "opbd"
	TTFTable_OS2               Table = "OS/2"
	TTFTable_post              Table = "post"
	TTFTable_prep              Table = "prep"
	TTFTable_prop              Table = "prop"
	TTFTable_sbix              Table = "sbix"
	TTFTable_trak              Table = "trak"
	TTFTable_vhea              Table = "vhea"
	TTFTable_vmtx              Table = "vmtx"
	TTFTable_xref              Table = "xref"
	TTFTable_Zapf              Table = "Zapf"
)

type Directory struct {
	head   DirectoryHeader
	tables map[Table]DirectoryRow
}

func (d Directory) sizeInFile() uint {
	return uint((4 + (2 * 4)) + (len(d.tables) * (4 * 4)))
}

func parseDirectory(data io.Reader) Directory {
	head := parseDirectoryHead(data)
	tMap := map[Table]DirectoryRow{}
	for i := uint16(0); i < head.NumTables; i++ {
		row := parseDirectoryRow(data)
		tMap[row.Table()] = row
	}
	return Directory{
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

func (d DirectoryRow) Table() Table {
	return Table(uint32ToString(d.Tag))
}

func parseDirectoryRow(data io.Reader) DirectoryRow {
	return parse[DirectoryRow](data)
}
