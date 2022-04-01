package ttf

type Table interface {
	Type() TableType
	SizeInFile() uint
}

type TableType string

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage,GoNameStartsWithPackageName
const (
	TableType_acnt      TableType = "acnt"
	TableType_ankr      TableType = "ankr"
	TableType_avar      TableType = "avar"
	TableType_bdat      TableType = "bdat"
	TableType_bhed      TableType = "bhed"
	TableType_bloc      TableType = "bloc"
	TableType_bsln      TableType = "bsln"
	TableType_cmap      TableType = "cmap"
	TableType_cvar      TableType = "cvar"
	TableType_cvt       TableType = "cvt "
	TableType_EBSC      TableType = "EBSC"
	TableType_fdsc      TableType = "fdsc"
	TableType_feat      TableType = "feat"
	TableType_fmtx      TableType = "fmtx"
	TableType_fond      TableType = "fond"
	TableType_fpgm      TableType = "fpgm"
	TableType_fvar      TableType = "fvar"
	TableType_gasp      TableType = "gasp"
	TableType_gcid      TableType = "gcid"
	TableType_glyf      TableType = "glyf"
	TableType_gvar      TableType = "gvar"
	TableType_hdmx      TableType = "hdmx"
	TableType_head      TableType = "head"
	TableType_hhea      TableType = "hhea"
	TableType_hmtx      TableType = "hmtx"
	TableType_just      TableType = "just"
	TableType_kern      TableType = "kern"
	TableType_kerx      TableType = "kerx"
	TableType_lcar      TableType = "lcar"
	TableType_loca      TableType = "loca"
	TableType_ltag      TableType = "ltag"
	TableType_maxp      TableType = "maxp"
	TableType_meta      TableType = "meta"
	TableType_mort      TableType = "mort"
	TableType_morx      TableType = "morx"
	TableType_name      TableType = "name"
	TableType_opbd      TableType = "opbd"
	TableType_OS2       TableType = "OS/2"
	TableType_post      TableType = "post"
	TableType_prep      TableType = "prep"
	TableType_prop      TableType = "prop"
	TableType_sbix      TableType = "sbix"
	TableType_Directory TableType = "snft"
	TableType_trak      TableType = "trak"
	TableType_vhea      TableType = "vhea"
	TableType_vmtx      TableType = "vmtx"
	TableType_xref      TableType = "xref"
	TableType_Zapf      TableType = "Zapf"
)
