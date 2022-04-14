package renderer

import (
	"github.com/go-gl/gl/all-core/gl"
)

type RenderPipeline interface {
	Bind()
	UnBind()
}

var _ RenderPipeline = (*defaultRenderPipeline)(nil)

type defaultRenderPipeline struct {
	vbo, ebo, vao uint32
	shader        Shader
}

func (d defaultRenderPipeline) Bind() {
	d.shader.Use()
	gl.BindVertexArray(d.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, d.vbo)
}

func (d defaultRenderPipeline) UnBind() {
	gl.BindVertexArray(0)
	d.shader.Detach()
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

}
