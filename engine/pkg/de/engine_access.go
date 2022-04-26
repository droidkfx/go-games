package de

import "github.com/droidkfx/go-games/engine/pkg/components"

// TODO this is odd that it lives here. The issue is that it needs to reference the engine directly. The engine also
// needs to type check on it. So they cannot live in seperate packages. Perhaps a further level of abstraction would
// fix this...

type EngineAccessor interface {
	components.GameObject
	SetEngine(engine DEngine)
}

var _ EngineAccessor = (*DefaultEngineAccessor)(nil)

type DefaultEngineAccessor struct {
	engine DEngine
}

func (d *DefaultEngineAccessor) SetEngine(e DEngine) {
	d.engine = e
}

func (d *DefaultEngineAccessor) GetEngine() DEngine {
	return d.engine
}
