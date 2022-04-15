package de

import (
	"log"
	"time"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/components/render"
	"github.com/droidkfx/go-games/engine/pkg/renderer"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type DEngine interface {
	Run() error
	AddGameObject(g components.GameObject)
}

type DEngineBuilder struct {
	engine *engine
}

func Builder() *DEngineBuilder {
	return &DEngineBuilder{
		engine: &engine{minFrameTime: time.Millisecond * 200},
	}
}

func (db *DEngineBuilder) MinFrameTime(duration time.Duration) *DEngineBuilder {
	db.engine.minFrameTime = duration
	return db
}

func (db *DEngineBuilder) Window(w *glfw.Window) *DEngineBuilder {
	db.engine.window = w
	db.engine.rSys = nil // We have to reset this since the render system depends on this
	return db
}

func (db *DEngineBuilder) RenderSystem(rSys renderer.RenderSystem) *DEngineBuilder {
	db.engine.rSys = rSys
	return db
}

func (db *DEngineBuilder) Build() DEngine {
	return db.engine
}

var _ DEngine = (*engine)(nil)

type engine struct {
	window       *glfw.Window
	objects      []components.GameObject
	uObjs        []components.UpdatableObject
	rObjs        []render.BaseComponent
	iLs          []components.KeyInputListener
	ogKeyCb      glfw.KeyCallback
	rSys         renderer.RenderSystem
	minFrameTime time.Duration
}

func (e *engine) AddGameObject(g components.GameObject) {
	e.objects = append(e.objects, g)
	if liObj, ok := g.(components.KeyInputListener); ok {
		e.iLs = append(e.iLs, liObj)
	}
	if rObj, ok := g.(render.BaseComponent); ok {
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
	initErr := e.rSys.Init()
	if initErr != nil {
		return initErr
	}

	delta := time.Duration(0)
	e.ogKeyCb = e.window.SetKeyCallback(e.handleKeySubs)
	defer e.window.SetKeyCallback(e.ogKeyCb)
	lastTime := time.Now()
	log.Println("DE initialized, entering game loop")
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
