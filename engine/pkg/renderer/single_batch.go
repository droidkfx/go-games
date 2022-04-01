package renderer

import (
	"log"

	"github.com/droidkfx/go-games/engine/pkg/components"
	"github.com/droidkfx/go-games/engine/pkg/gl_util"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func SingleBatch(window *glfw.Window) TypedRenderSystem {
	sys := &singleBatchRenderSystem{
		window: window,
	}
	return sys
}

var _ TypedRenderSystem = (*singleBatchRenderSystem)(nil)

type singleBatchRenderSystem struct {
	window              *glfw.Window
	vbo, ebo, vao       uint32
	currentVertexOffset uint32
	shader              Shader
	vertexList          []float32
	elementList         []uint32
}

func (s *singleBatchRenderSystem) Type() components.RenderType {
	return components.RenderType_MESH
}

func (s *singleBatchRenderSystem) Init() error {
	gl.GenBuffers(1, &s.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)

	gl.GenBuffers(1, &s.ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, s.ebo)

	gl.GenVertexArrays(1, &s.vao)
	gl.BindVertexArray(s.vao)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 2, gl.FLOAT, false, 5*gl_util.SizeofFloat32, 0)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 5*gl_util.SizeofFloat32, 2*gl_util.SizeofFloat32)
	shader, shaderErr := ShaderFromSources(defaultShaderSrcs)
	if shaderErr != nil {
		return shaderErr
	}
	s.shader = shader

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

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
	if len(s.vertexList) == 0 {
		return
	}
	s.shader.Use()
	defer s.shader.Detach()

	gl.BindVertexArray(s.vao)
	defer gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)
	defer gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BufferData(gl.ARRAY_BUFFER, len(s.vertexList)*gl_util.SizeofFloat32, gl.Ptr(s.vertexList), gl.STREAM_DRAW)
	gl.DrawElements(gl.TRIANGLES, int32(len(s.elementList)), gl.UNSIGNED_INT, gl.Ptr(s.elementList))

	s.currentVertexOffset = 0
	s.vertexList = make([]float32, 0, len(s.vertexList))
	s.elementList = make([]uint32, 0, len(s.elementList))
}