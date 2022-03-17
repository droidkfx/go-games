package renderer

import (
	"log"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/gl_util"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type RenderSystem interface {
	Init() error
	Process(obj components.RenderObject)
	Render()
}

func SingleBatch(window *glfw.Window) RenderSystem {
	sys := &singleBatchRenderSystem{
		window: window,
	}

	gl.GenBuffers(1, &sys.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, sys.vbo)

	gl.GenBuffers(1, &sys.ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, sys.ebo)

	gl.GenVertexArrays(1, &sys.vao)
	gl.BindVertexArray(sys.vao)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 2, gl.FLOAT, false, 5*gl_util.SizeofFloat32, 0)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 5*gl_util.SizeofFloat32, 2*gl_util.SizeofFloat32)

	return sys
}

var _ RenderSystem = (*singleBatchRenderSystem)(nil)

type singleBatchRenderSystem struct {
	window              *glfw.Window
	vbo, ebo, vao       uint32
	currentVertexOffset uint32
	shader              Shader
	vertexList          []float32
	elementList         []uint32
}

func (s *singleBatchRenderSystem) Init() error {
	shader, shaderErr := NewShader("default")
	if shaderErr != nil {
		return shaderErr
	}
	s.shader = shader
	s.shader.Use()
	return nil
}

func (s *singleBatchRenderSystem) Process(ro components.RenderObject) {
	if ro.Type() != components.RenderType_MESH {
		log.Println("Tried to paint non mesh render object")
		return
	}
	mro := ro.(components.MeshRender)

	vD, eD := mro.GetMeshData()
	for i := 0; i < len(eD); i++ {
		eD[i] = eD[i] + s.currentVertexOffset
	}
	s.vertexList = append(s.vertexList, vD...)
	s.elementList = append(s.elementList, eD...)

	s.currentVertexOffset += uint32(len(eD))
}

func (s *singleBatchRenderSystem) Render() {

	gl.BufferData(gl.ARRAY_BUFFER, len(s.vertexList)*gl_util.SizeofFloat32, gl.Ptr(s.vertexList), gl.STREAM_DRAW)
	gl.DrawElements(gl.TRIANGLES, int32(len(s.elementList)), gl.UNSIGNED_INT, gl.Ptr(s.elementList))

	s.window.SwapBuffers()

	// Clean up for next render call
	s.currentVertexOffset = 0
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	s.vertexList = make([]float32, 0, len(s.vertexList))
	s.elementList = make([]uint32, 0, len(s.elementList))
}
