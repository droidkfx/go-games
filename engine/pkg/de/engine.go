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
	currentEngine = &engine{window: window, rSys: rSys, minFrameTime: time.Millisecond * 200}
	return currentEngine, nil
}

var _ DEngine = (*engine)(nil)

type engine struct {
	window       *glfw.Window
	objects      []components.GameObject
	uObjs        []components.UpdatableObject
	rObjs        []components.RenderObject
	iLs          []components.KeyInputListener
	ogKeyCb      glfw.KeyCallback
	rSys         renderer.RenderSystem
	minFrameTime time.Duration
}

func AddGameObject(g components.GameObject) {
	if currentEngine == nil {
		return
	}
	currentEngine.AddGameObject(g)
}

func (e *engine) AddGameObject(g components.GameObject) {
	e.objects = append(e.objects, g)
	if liObj, ok := g.(components.KeyInputListener); ok {
		e.iLs = append(e.iLs, liObj)
	}
	if rObj, ok := g.(components.RenderObject); ok {
		e.rObjs = append(e.rObjs, rObj)
	}
	if uObj, ok := g.(components.UpdatableObject); ok {
		e.uObjs = append(e.uObjs, uObj)
	}
}

func (e *engine) handleKeySubs(_ *glfw.Window, key glfw.Key, _ int, action glfw.Action, mods glfw.ModifierKey) {
	for _, ikls := range e.iLs {
		ikls.HandleKeyInput(key, action, mods)
	}
}

func (e *engine) Run() error {
	log.Println("starting DE")

	delta := time.Duration(0)
	e.ogKeyCb = e.window.SetKeyCallback(e.handleKeySubs)
	defer e.window.SetKeyCallback(e.ogKeyCb)
	lastTime := time.Now()
	for !e.window.ShouldClose() {
		for i := 0; i < len(e.uObjs); i++ {
			e.uObjs[i].Update(delta)
		}

		for i := 0; i < len(e.rObjs); i++ {
			e.rSys.Process(e.rObjs[i])
		}
		e.rSys.Render()

		loopTime := time.Now()
		delta = loopTime.Sub(lastTime)
		if delta > e.minFrameTime {
			log.Printf("Frame drop, application is running slow. Delta was %v. Using %v instead\n", delta, e.minFrameTime)
			delta = e.minFrameTime
		}
		lastTime = loopTime
		glfw.PollEvents()
	}

	return nil
}
