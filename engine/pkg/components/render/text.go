package render

import "github.com/droidkfx/go-games/engine/pkg/d_types"

type TextRender interface {
	BaseComponent
	GetPosition() d_types.V2f32
	GetSize() float32
	GetText() string
}

var _ BaseComponent = (*TextRenderObject)(nil)
var _ TextRender = (*TextRenderObject)(nil)

type TextRenderObject struct{}

func (t TextRenderObject) Type() Type {
	return TypeText
}

func (t TextRenderObject) GetPosition() d_types.V2f32 {
	return d_types.V2f32{}
}

func (t TextRenderObject) GetSize() float32 {
	return 0.05
}

func (t TextRenderObject) GetText() string {
	return "Hello World"
}
