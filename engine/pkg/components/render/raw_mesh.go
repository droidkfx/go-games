package render

import (
	"github.com/droidkfx/go-games/engine/pkg/components/render/shape"
)

type Mesh interface {
	BaseComponent
	GetMeshData() shape.Mesh
}

var _ BaseComponent = (*RawMesh)(nil)
var _ Mesh = (*RawMesh)(nil)

type RawMesh struct{}

func (m *RawMesh) GetMeshData() shape.Mesh {
	return shape.Mesh{}
}

func (m *RawMesh) Type() Type {
	return TypeMesh
}
