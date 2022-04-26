package main

import (
	"log"

	"github.com/droidkfx/go-games/engine/pkg/de"
	"github.com/droidkfx/go-games/engine/pkg/gl_util"
	"github.com/droidkfx/go-games/engine/pkg/renderer"
	"github.com/droidkfx/go-games/snake/internal"
)

func main() {
	log.Println("starting snake")
	config := gl_util.DefaultGlfwConfig()
	config.WindowSize(1000, 1000)
	window, initErr := gl_util.InitializeGlfwWindow(config)
	if initErr != nil {
		log.Fatalf(initErr.Error())
	}
	defer gl_util.UnInitialize()

	rootRenderer := renderer.RoutingMultiRenderSystem(window)
	rootRenderer.SetMapping(renderer.SingleBatch(window))
	rootRenderer.SetMapping(renderer.TextRenderSystem(window))
	engine := de.Builder().Window(window).RenderSystem(rootRenderer).Build()

	engine.AddGameObject(&internal.Snake{})

	if runErr := engine.Run(); runErr != nil {
		log.Fatalf(runErr.Error())
	}
}
