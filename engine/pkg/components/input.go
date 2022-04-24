package components

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type KeyInputListener interface {
	GameObject
	HandleKeyInput(glfw.Key, glfw.Action, glfw.ModifierKey)
}

var _ KeyInputListener = (*KeyInputLogger)(nil)

type KeyInputLogger struct{}

func (s *KeyInputLogger) HandleKeyInput(key glfw.Key, action glfw.Action, mod glfw.ModifierKey) {
	log.Printf("Key action: `%+v` Key Code: `%+v` Key Modifier: `%+v` \n", action, key, mod)
}
