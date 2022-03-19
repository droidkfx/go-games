package renderer

import (
	"log"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type RenderSystem interface {
	Init() error
	Process(obj components.RenderObject)
	Render()
}

type TypedRenderSystem interface {
	RenderSystem
	Type() components.RenderType
}

type MappingRenderSystem interface {
	RenderSystem
	SetMapping(s TypedRenderSystem)
}

func RoutingMultiRenderSystem(window *glfw.Window) MappingRenderSystem {
	return &routingMultiRenderSystem{renderers: map[components.RenderType]RenderSystem{}, w: window}
}

var _ MappingRenderSystem = (*routingMultiRenderSystem)(nil)

type routingMultiRenderSystem struct {
	w         *glfw.Window
	renderers map[components.RenderType]RenderSystem
}

func (r *routingMultiRenderSystem) SetMapping(s TypedRenderSystem) {
	r.renderers[s.Type()] = s
}

func (r *routingMultiRenderSystem) Init() error {
	for _, s := range r.renderers {
		if err := s.Init(); err != nil {
			return err
		}
	}
	return nil
}

func (r *routingMultiRenderSystem) Process(ro components.RenderObject) {
	if s, ok := r.renderers[ro.Type()]; ok {
		s.Process(ro)
	} else {
		log.Printf("Render type %+v not mapped to a render system. Skipping.\n", ro.Type())
	}
}

func (r *routingMultiRenderSystem) Render() {
	for _, s := range r.renderers {
		s.Render()
	}
	r.w.SwapBuffers()
	// Clean up for next render call
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
