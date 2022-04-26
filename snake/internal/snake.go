package internal

import (
	"time"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/d_types"
	"github.com/droidkfx/go-games/engine/pkg/de"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var _ components.InitObject = (*Snake)(nil)
var _ components.UpdatableObject = (*Snake)(nil)
var _ components.KeyInputListener = (*Snake)(nil)
var _ de.EngineAccessor = (*Snake)(nil)

var (
	snakeSegSize   float32 = 0.05
	gridSize       float32 = 0.06
	initialSize            = 5
	liveSnakeColor         = d_types.Color_GREEN
	deadSnakeColor         = d_types.Color_RED
)

type Snake struct {
	de.DefaultEngineAccessor
	segments  []*SnakeSegment
	direction d_types.V2f32
	dead      bool
	ticker
}

func (s *Snake) HandleKeyInput(key glfw.Key, action glfw.Action, _ glfw.ModifierKey) {
	if action == glfw.Press {
		switch key {
		case glfw.KeyA:
			if s.direction.X == 0.0 {
				s.direction = d_types.V2f32{X: -gridSize}
			}
		case glfw.KeyD:
			if s.direction.X == 0.0 {
				s.direction = d_types.V2f32{X: gridSize}
			}
		case glfw.KeyW:
			if s.direction.Y == 0.0 {
				s.direction = d_types.V2f32{Y: gridSize}
			}
		case glfw.KeyS:
			if s.direction.Y == 0.0 {
				s.direction = d_types.V2f32{Y: -gridSize}
			}
		}
	}
}

func (s *Snake) Update(delta time.Duration) {
	s.CallOnTick(delta, s.Tick)
}

func (s *Snake) Tick() {
	if s.dead {
		return
	}
	s.move()
	// check first if we have set a direction. In the beginning we would not have
	if (s.direction.X != 0.0 || s.direction.Y != 0.0) && s.hasOverlap() {
		s.die()
	}
}

func (s *Snake) Init() {
	s.ticker.tickTime = time.Millisecond * 250
	s.setupInitialSegments()
	for _, segment := range s.segments {
		s.GetEngine().AddGameObject(segment)
	}
}

func (s *Snake) move() {
	for i := len(s.segments) - 1; i >= 0; i-- {
		if i == 0 {
			s.segments[i].loc = s.segments[i].loc.Add(s.direction)
		} else {
			s.segments[i].loc = s.segments[i-1].loc
		}
	}
}

func (s *Snake) setupInitialSegments() {
	s.segments = make([]*SnakeSegment, 0, initialSize)
	for i := 0; i < initialSize; i++ {
		s.segments = append(s.segments, &SnakeSegment{size: snakeSegSize, color: liveSnakeColor})
	}
}

func (s *Snake) hasOverlap() bool {
	head := s.segments[0]
	for i := 1; i < len(s.segments); i++ {
		if head.loc.Eq(s.segments[i].loc) {
			return true
		}
	}
	return false
}

func (s *Snake) die() {
	s.dead = true
	for _, seg := range s.segments {
		seg.color = deadSnakeColor
	}
}
