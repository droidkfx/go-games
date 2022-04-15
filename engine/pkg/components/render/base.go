package render

import (
	"github.com/droidkfx/go-games/engine/pkg/components"
)

// BaseComponent is a marker interface to allow the render system to take in any object that wishes to be rendered
type BaseComponent interface {
	components.GameObject
	Type() Type
}

type Type int

//goland:noinspection GoUnusedConst
const (
	TypeUnknown Type = iota
	TypeText
	TypeMesh
)
