package de

import (
	"log"
	"time"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/renderer"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var currentEngine *engine

type DEngine interface {
	Run() error
}

func CreateEngine(window *glfw.Window) (DEngine, error) {
	rSys := renderer.SingleBatch(window)
	if rSysErr := rSys.Init(); rSysErr != nil {
		return nil, rSysErr
	}
	currentEngine = &engine{window: window, rSys: rSys}
	return currentEngine, nil
}

var _ DEngine = (*engine)(nil)

type engine struct {
	window *glfw.Window

	objects []components.GameObject
	rSys    renderer.RenderSystem
}

func AddGameObject(g components.GameObject) {
	if currentEngine == nil {
		return
	}
	currentEngine.objects = append(currentEngine.objects, g)
}

func (e *engine) Run() error {
	log.Println("Starting DE")

	delta := time.Duration(0)
	lastTime := time.Now()
	for !e.window.ShouldClose() {
		for i := 0; i < len(e.objects); i++ {
			obj := e.objects[i]
			if uObj, ok := obj.(components.UpdatableObject); ok {
				uObj.Update(delta)
			}
		}

		for i := 0; i < len(e.objects); i++ {
			obj := e.objects[i]
			if rObj, ok := obj.(components.RenderObject); ok {
				e.rSys.Process(rObj)
			}
		}
		e.rSys.Render()

		loopTime := time.Now()
		delta = loopTime.Sub(lastTime)
		lastTime = loopTime
		glfw.PollEvents()
	}

	return nil
}
