package c_impl

import (
	"time"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/components/render"
	"github.com/droidkfx/go-games/engine/pkg/components/render/shape"
	"github.com/droidkfx/go-games/engine/pkg/d_types"
)

var _ components.GameObject = (*SimpleQuad)(nil)
var _ render.BaseComponent = (*SimpleQuad)(nil)
var _ render.Mesh = (*SimpleQuad)(nil)
var _ components.UpdatableObject = (*SimpleQuad)(nil)
var _ components.KeyInputListener = (*SimpleQuad)(nil)
var _ components.InitObject = (*SimpleQuad)(nil)

type SimpleQuad struct {
	SimpleQuadMesh
	components.KeyInputLogger
	components.InitLogger
	modifier float32
	minSize  float32
	maxSize  float32
}

func TestQuad() components.GameObject {
	return &SimpleQuad{SimpleQuadMesh: SimpleQuadMesh{size: 0.5}, modifier: 0.3, minSize: 0.1, maxSize: 1.0}
}

func (s *SimpleQuad) Update(delta time.Duration) {
	s.size += (float32(delta) / float32(time.Second)) * s.modifier
	if (s.size > s.maxSize && s.modifier > 0) || (s.size < s.minSize && s.modifier < 0) {
		s.modifier *= -1
	}
}

type SimpleQuadMesh struct {
	render.RawMesh
	position d_types.V2f32
	size     float32
}

func (s *SimpleQuadMesh) GetMeshData() shape.Mesh {
	return shape.SolidSquareCenteredAt(s.size, s.position, d_types.Color_RED)
}
