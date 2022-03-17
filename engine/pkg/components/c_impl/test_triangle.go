package c_impl

import (
	"time"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/d_types"
)

var _ components.GameObject = (*SimpleTriangle)(nil)
var _ components.RenderObject = (*SimpleTriangle)(nil)
var _ components.MeshRender = (*SimpleTriangle)(nil)
var _ components.UpdatableObject = (*SimpleTriangle)(nil)

type SimpleTriangle struct {
	components.MeshRenderObject
	position d_types.V2f32
	size     float32
	modifier float32
}

func TestTriangle() components.GameObject {
	return &SimpleTriangle{size: 0.5, modifier: 0.3}
}

func (s *SimpleTriangle) Update(delta time.Duration) {
	s.size += (float32(delta) / float32(time.Second)) * s.modifier
	if s.size > 1 || s.size < 0.1 {
		s.modifier *= -1
	}
}

func (s *SimpleTriangle) GetMeshData() ([]float32, []uint32) {
	top := s.position.Add(d_types.V2f32{Y: s.size})
	botLeft := s.position.Add(d_types.V2f32{X: -s.size, Y: -s.size})
	botRight := s.position.Add(d_types.V2f32{X: s.size, Y: -s.size})
	return []float32{
		top.X, top.Y, 1.0, 0.0, 0.0,
		botLeft.X, botLeft.Y, 0.0, 1.0, 0.0,
		botRight.X, botRight.Y, 0.0, 0.0, 1.0,
	}, []uint32{0, 1, 2}
}
