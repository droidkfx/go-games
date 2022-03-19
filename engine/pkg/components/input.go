package components

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type KeyInputListener interface {
	GameObject
	HandleKeyInput(glfw.Key, glfw.Action, glfw.ModifierKey)
}
