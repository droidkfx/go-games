package c_impl

import (
	"time"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/components/render"
	"github.com/droidkfx/go-games/engine/pkg/components/render/shape"
	"github.com/droidkfx/go-games/engine/pkg/d_types"
)

var _ components.GameObject = (*SimpleTriangle)(nil)
var _ render.BaseComponent = (*SimpleTriangle)(nil)
var _ render.Mesh = (*SimpleTriangle)(nil)
var _ components.UpdatableObject = (*SimpleTriangle)(nil)
var _ components.KeyInputListener = (*SimpleTriangle)(nil)

type SimpleTriangle struct {
	SimpleTriMesh
	components.SimpleKeyLoggerListener
	modifier float32
	minSize  float32
	maxSize  float32
}

func TestTriangle() components.GameObject {
	return &SimpleTriangle{SimpleTriMesh: SimpleTriMesh{size: 0.5}, modifier: 0.3, minSize: 0.1, maxSize: 1.0}
}

func (s *SimpleTriangle) Update(delta time.Duration) {
	s.size += (float32(delta) / float32(time.Second)) * s.modifier
	if (s.size > s.maxSize && s.modifier > 0) || (s.size < s.minSize && s.modifier < 0) {
		s.modifier *= -1
	}
}

type SimpleTriMesh struct {
	render.RawMesh
	position d_types.V2f32
	size     float32
}

func (s *SimpleTriMesh) GetMeshData() shape.Mesh {
	top := s.position.Add(d_types.V2f32{Y: s.size})
	botLeft := s.position.Add(d_types.V2f32{X: -s.size, Y: -s.size})
	botRight := s.position.Add(d_types.V2f32{X: s.size, Y: -s.size})
	return shape.Mesh{
		Verts: []float32{
			top.X, top.Y, 1.0, 0.0, 0.0,
			botLeft.X, botLeft.Y, 0.0, 1.0, 0.0,
			botRight.X, botRight.Y, 0.0, 0.0, 1.0,
		},
		Elems: []uint32{0, 1, 2},
	}
}
