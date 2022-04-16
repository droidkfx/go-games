package renderer

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/droidkfx/go-games/engine/pkg/components/render"
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
	renderPipe          RenderPipeline
	currentVertexOffset uint32
	vertexList          []float32
	elementList         []uint32
}

func (s *singleBatchRenderSystem) Type() render.Type {
	return render.TypeMesh
}

func (s *singleBatchRenderSystem) Init() error {
	rp := &defaultRenderPipeline{}

	shader, shaderErr := ShaderFromSources(defaultShaderSrcs)
	if shaderErr != nil {
		return shaderErr
	}
	rp.shader = shader

	gl.GenVertexArrays(1, &rp.vao)
	gl.BindVertexArray(rp.vao)

	gl.GenBuffers(1, &rp.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, rp.vbo)

	gl.GenBuffers(1, &rp.ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, rp.ebo)

	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(0, 2, gl.FLOAT, false, 5*gl_util.SizeofFloat32, 0)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 5*gl_util.SizeofFloat32, 2*gl_util.SizeofFloat32)

	rp.UnBind()
	s.renderPipe = rp

	return nil
}

func (s *singleBatchRenderSystem) Process(ro render.BaseComponent) {
	if ro.Type() != render.TypeMesh {
		log.Println("Tried to paint non mesh render object")
		return
	}
	mro := ro.(render.Mesh)

	meshData := mro.GetMeshData()
	for i := 0; i < len(meshData.Elems); i++ {
		meshData.Elems[i] = meshData.Elems[i] + s.currentVertexOffset
	}
	s.vertexList = append(s.vertexList, meshData.Verts...)
	s.elementList = append(s.elementList, meshData.Elems...)

	s.currentVertexOffset += uint32(len(meshData.Elems))
}

func (s *singleBatchRenderSystem) Render() {
	if len(s.vertexList) == 0 {
		return
	}

	s.renderPipe.Bind()
	defer s.renderPipe.UnBind()

	gl.BufferData(gl.ARRAY_BUFFER, len(s.vertexList)*gl_util.SizeofFloat32, gl.Ptr(s.vertexList), gl.STREAM_DRAW)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(s.elementList)*gl_util.SizeofUint32, gl.Ptr(s.elementList), gl.STREAM_DRAW)
	gl.DrawElements(gl.TRIANGLES, int32(len(s.elementList)), gl.UNSIGNED_INT, unsafe.Pointer(nil))

	s.currentVertexOffset = 0
	s.vertexList = make([]float32, 0, len(s.vertexList))
	s.elementList = make([]uint32, 0, len(s.elementList))

	for glErr := gl.GetError(); glErr != gl.NO_ERROR; glErr = gl.GetError() {
		fmt.Printf("ERR: %v\n", glErr)
	}
}
