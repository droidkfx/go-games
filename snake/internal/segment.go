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

var snakeSegmentSize float32 = 0.05
var snakeColor = d_types.Color_RED

func NewSegment(loc d_types.V2f32) *SnakeSegment {
	return &SnakeSegment{loc: loc}
}

type SnakeSegment struct {
	render.RawMesh
	loc d_types.V2f32
}

func (m *SnakeSegment) GetMeshData() shape.Mesh {
	return shape.SolidSquareCenteredAt(snakeSegmentSize, m.loc, snakeColor)
}
