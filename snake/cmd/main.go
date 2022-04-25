package main

import (
	"log"

	"github.com/droidkfx/go-games/engine/pkg/d_types"
	"github.com/droidkfx/go-games/engine/pkg/de"
	"github.com/droidkfx/go-games/engine/pkg/gl_util"
	"github.com/droidkfx/go-games/engine/pkg/renderer"
	"github.com/droidkfx/go-games/snake/internal"
)

func main() {
	log.Println("starting snake")
	window, initErr := gl_util.InitializeGlfwWindow(gl_util.DefaultGlfwConfig())
	if initErr != nil {
		log.Fatalf(initErr.Error())
	}
	defer gl_util.UnInitialize()

	rootRenderer := renderer.RoutingMultiRenderSystem(window)
	rootRenderer.SetMapping(renderer.SingleBatch(window))
	rootRenderer.SetMapping(renderer.TextRenderSystem(window))
	engine := de.Builder().Window(window).RenderSystem(rootRenderer).Build()

	//engine.AddGameObject(internal.NewSegment(d_types.V2f32{X: 0.000, Y: 0.000}))
	//engine.AddGameObject(internal.NewSegment(d_types.V2f32{X: 0.125, Y: 0.125}))
	for i := float32(0.0); i < 1.0; i += 0.125 {
		engine.AddGameObject(internal.NewSegment(d_types.V2f32{X: i, Y: i}))
	}

	if runErr := engine.Run(); runErr != nil {
		log.Fatalf(runErr.Error())
	}
}
