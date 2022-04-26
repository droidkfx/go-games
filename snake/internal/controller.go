package internal

import (
	"math/rand"
	"time"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/d_types"
	"github.com/droidkfx/go-games/engine/pkg/de"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var _ components.InitObject = (*Controller)(nil)
var _ components.UpdatableObject = (*Controller)(nil)
var _ components.KeyInputListener = (*Controller)(nil)
var _ de.EngineAccessor = (*Controller)(nil)

var (
	gridSize        float32 = 0.3
	snakeSegSize            = gridSize * 0.95
	foodSize                = snakeSegSize / 3
	maxGridIndex            = int(1.0 / gridSize)
	minGridLoc              = float32(maxGridIndex) * -gridSize
	maxGridLoc              = float32(maxGridIndex) * gridSize
	totalGridSize           = maxGridIndex * 2
	initialSnakeLen         = 5
	liveSnakeColor          = d_types.Color_GREEN
	deadSnakeColor          = d_types.Color_RED
	foodEqEps               = float64(0.5 * gridSize)
)

type Controller struct {
	de.DefaultEngineAccessor
	snake *Snake
	food  *Unit
}

func (c *Controller) HandleKeyInput(key glfw.Key, action glfw.Action, _ glfw.ModifierKey) {
	if action == glfw.Press {
		if key == glfw.Key0 {
			c.placeFood()
		}
	}
}

func (c *Controller) Update(_ time.Duration) {
	headLoc := c.snake.segments[0].loc
	if headLoc.EqEps(c.food.loc, foodEqEps) {
		c.placeFood()
		c.snake.Grow()
	}
	if headLoc.X >= maxGridLoc || headLoc.Y >= maxGridLoc || headLoc.X <= minGridLoc || headLoc.Y <= minGridLoc {
		c.snake.Die()
	}
}

func (c *Controller) Init() {
	c.snake = &Snake{}
	c.food = &Unit{color: d_types.Color_BLUE, size: foodSize}
	c.placeFood()
	c.GetEngine().AddGameObject(c.snake)
	c.GetEngine().AddGameObject(c.food)
}

func (c *Controller) placeFood() {
	c.food.loc = d_types.V2f32{
		X: gridSize * float32(rand.Intn(totalGridSize-1)-maxGridIndex+1),
		Y: gridSize * float32(rand.Intn(totalGridSize-1)-maxGridIndex+1),
	}
	for _, seg := range c.snake.segments {
		if seg.loc.EqEps(c.food.loc, foodEqEps) {
			c.placeFood()
		}
	}
}
