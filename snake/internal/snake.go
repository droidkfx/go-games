package internal

import (
	"time"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/d_types"
	"github.com/droidkfx/go-games/engine/pkg/de"
)

var _ components.InitObject = (*Snake)(nil)
var _ components.UpdatableObject = (*Snake)(nil)
var _ de.EngineAccessor = (*Snake)(nil)

var snakeSegSize float32 = 0.05
var gridSize float32 = 0.06

type Snake struct {
	de.DefaultEngineAccessor
	segments []*SnakeSegment

	ticker
}

func (s *Snake) Update(delta time.Duration) {
	s.CallOnTick(delta, s.Tick)
}

func (s *Snake) Tick() {
	for i, seg := range s.segments {
		s.segments[i].loc = seg.loc.Add(d_types.V2f32{Y: gridSize})
	}
}

func (s *Snake) Init() {
	s.ticker.tickTime = time.Millisecond * 250
	s.setupInitialSegments()
	for _, segment := range s.segments {
		s.GetEngine().AddGameObject(segment)
	}
}

func (s *Snake) setupInitialSegments() {
	s.segments = []*SnakeSegment{
		{
			loc:   d_types.V2f32{X: 0, Y: gridSize * -5},
			size:  snakeSegSize,
			color: d_types.Color_RED,
		},
		{
			loc:   d_types.V2f32{X: 0, Y: gridSize * -6},
			size:  snakeSegSize,
			color: d_types.Color_RED,
		},
		{
			loc:   d_types.V2f32{X: 0, Y: gridSize * -7},
			size:  snakeSegSize,
			color: d_types.Color_RED,
		},
		{
			loc:   d_types.V2f32{X: 0, Y: gridSize * -8},
			size:  snakeSegSize,
			color: d_types.Color_RED,
		},
	}
}
