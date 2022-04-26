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

type Snake struct {
	de.DefaultEngineAccessor
	segments      []*Unit
	direction     d_types.V2f32
	nextDirection d_types.V2f32
	dead          bool
	grow          bool
	ticker
}

func (s *Snake) HandleKeyInput(key glfw.Key, action glfw.Action, _ glfw.ModifierKey) {
	if action == glfw.Press {
		switch key {
		case glfw.KeyA:
			if s.direction.X == 0.0 {
				s.nextDirection = d_types.V2f32{X: -gridSize}
			}
		case glfw.KeyD:
			if s.direction.X == 0.0 {
				s.nextDirection = d_types.V2f32{X: gridSize}
			}
		case glfw.KeyW:
			if s.direction.Y == 0.0 {
				s.nextDirection = d_types.V2f32{Y: gridSize}
			}
		case glfw.KeyS:
			if s.direction.Y == 0.0 {
				s.nextDirection = d_types.V2f32{Y: -gridSize}
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
		s.Die()
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
	s.direction = s.nextDirection
	var newSegment *Unit
	if s.grow {
		s.grow = false
		newSegLoc := s.segments[len(s.segments)-1].loc
		newSegment = &Unit{loc: newSegLoc, color: liveSnakeColor, size: snakeSegSize}
	}
	for i := len(s.segments) - 1; i >= 0; i-- {
		if i == 0 {
			s.segments[i].loc = s.segments[i].loc.Add(s.direction)
		} else {
			s.segments[i].loc = s.segments[i-1].loc
		}
	}
	if newSegment != nil {
		s.segments = append(s.segments, newSegment)
		s.GetEngine().AddGameObject(newSegment)
	}
}

func (s *Snake) setupInitialSegments() {
	s.segments = make([]*Unit, 0, initialSnakeLen)
	for i := 0; i < initialSnakeLen; i++ {
		s.segments = append(s.segments, &Unit{size: snakeSegSize, color: liveSnakeColor})
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

func (s *Snake) Die() {
	s.dead = true
	for _, seg := range s.segments {
		seg.color = deadSnakeColor
	}
}

func (s *Snake) Grow() {
	s.grow = true
}
