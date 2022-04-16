package main

import (
	"log"

	"github.com/droidkfx/go-games/engine/pkg/components/c_impl"
	"github.com/droidkfx/go-games/engine/pkg/de"
	"github.com/droidkfx/go-games/engine/pkg/gl_util"
	"github.com/droidkfx/go-games/engine/pkg/renderer"
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

	engine.AddGameObject(c_impl.TestTriangle())

	if runErr := engine.Run(); runErr != nil {
		log.Fatalf(runErr.Error())
	}
}
