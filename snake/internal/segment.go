package internal

import (
	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/components/render"
	"github.com/droidkfx/go-games/engine/pkg/components/render/shape"
	"github.com/droidkfx/go-games/engine/pkg/d_types"
)

var _ components.GameObject = (*SnakeSegment)(nil)
var _ render.BaseComponent = (*SnakeSegment)(nil)
var _ render.Mesh = (*SnakeSegment)(nil)

type SnakeSegment struct {
	render.RawMesh
	loc   d_types.V2f32
	size  float32
	color d_types.ColorRGB
}

func (m *SnakeSegment) GetMeshData() shape.Mesh {
	return shape.SolidSquareCenteredAt(m.size, m.loc, m.color)
}
