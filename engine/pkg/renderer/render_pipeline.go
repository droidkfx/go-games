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

func (d *defaultRenderPipeline) Bind() {
	d.shader.Use()
	gl.BindVertexArray(d.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, d.vbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, d.ebo)
}

func (d *defaultRenderPipeline) UnBind() {
	d.shader.Detach()
	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

}
