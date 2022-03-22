package components

import (
	"github.com/droidkfx/go-games/engine/pkg/d_types"
)

// RenderObject is a marker interface to allow the render system to take in any object that wishes to be rendered
type RenderObject interface {
	GameObject
	Type() RenderType
}

type RenderType int

//goland:noinspection GoUnusedConst
const (
	RenderType_UNKNOWN RenderType = iota
	RenderType_MESH
	RenderType_TEXT
)

type MeshRender interface {
	RenderObject
	GetMeshData() ([]float32, []uint32)
}

var _ RenderObject = (*MeshRenderObject)(nil)
var _ MeshRender = (*MeshRenderObject)(nil)

type MeshRenderObject struct{}

func (m *MeshRenderObject) GetMeshData() ([]float32, []uint32) {
	return make([]float32, 0), make([]uint32, 0)
}

func (m *MeshRenderObject) Type() RenderType {
	return RenderType_MESH
}

type TextRender interface {
	RenderObject
	GetPosition() d_types.V2f32
	GetSize() float32
	GetText() string
}

var _ RenderObject = (*TextRenderObject)(nil)
var _ TextRender = (*TextRenderObject)(nil)

type TextRenderObject struct{}

func (t TextRenderObject) Type() RenderType {
	return RenderType_TEXT
}

func (t TextRenderObject) GetPosition() d_types.V2f32 {
	return d_types.V2f32{}
}

func (t TextRenderObject) GetSize() float32 {
	return 0.05
}

func (t TextRenderObject) GetText() string {
	return "Hello World"
}
