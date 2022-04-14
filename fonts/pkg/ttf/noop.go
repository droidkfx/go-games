package ttf

import (
	"io"
)

var _ Table = (*noopTable)(nil)

type noopTable struct{}

func (n noopTable) IsOpenTypeTable() bool { return false }
func (n noopTable) SizeInFile() uint      { return 0 }
func (n noopTable) Type() TableType       { return TableType_UNKNOWN }

func noop(_ io.Reader) any { return &noopTable{} }
