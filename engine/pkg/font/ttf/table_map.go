package ttf

import (
	"io"
)

func ParseTable[T any](tableType TableType, reader io.Reader) T {
	if v, ok := parserMap[tableType]; ok {
		return v(reader).(T)
	}
	return *new(T)
}

var (
	parserMap = map[TableType]func(reader io.Reader) any{
		TableType_acnt:      noop,
		TableType_ankr:      noop,
		TableType_avar:      noop,
		TableType_bdat:      noop,
		TableType_bhed:      noop,
		TableType_bloc:      noop,
		TableType_bsln:      noop,
		TableType_cmap:      noop,
		TableType_cvar:      noop,
		TableType_cvt:       noop,
		TableType_EBSC:      noop,
		TableType_fdsc:      noop,
		TableType_feat:      noop,
		TableType_fmtx:      noop,
		TableType_fond:      noop,
		TableType_fpgm:      noop,
		TableType_fvar:      noop,
		TableType_gasp:      noop,
		TableType_gcid:      noop,
		TableType_glyf:      noop,
		TableType_gvar:      noop,
		TableType_hdmx:      noop,
		TableType_head:      noop,
		TableType_hhea:      noop,
		TableType_hmtx:      noop,
		TableType_just:      noop,
		TableType_kern:      noop,
		TableType_kerx:      noop,
		TableType_lcar:      noop,
		TableType_loca:      noop,
		TableType_ltag:      noop,
		TableType_maxp:      noop,
		TableType_meta:      noop,
		TableType_mort:      noop,
		TableType_morx:      noop,
		TableType_name:      noop,
		TableType_opbd:      noop,
		TableType_OS2:       noop,
		TableType_post:      noop,
		TableType_prep:      noop,
		TableType_prop:      noop,
		TableType_sbix:      noop,
		TableType_Directory: parseDirectory,
		TableType_trak:      noop,
		TableType_vhea:      noop,
		TableType_vmtx:      noop,
		TableType_xref:      noop,
		TableType_Zapf:      noop,
	}
)

func noop(_ io.Reader) any {
	return nil
}
