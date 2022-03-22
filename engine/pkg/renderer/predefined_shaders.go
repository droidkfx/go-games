package renderer

import (
	"github.com/go-gl/gl/all-core/gl"
)

var (
	defaultShaderSrcs = map[uint32][]byte{
		gl.VERTEX_SHADER:   []byte(defaultVertShader + "\x00"),
		gl.FRAGMENT_SHADER: []byte(defaultFragShader + "\x00"),
	}
	textShaderSrcs = map[uint32][]byte{
		gl.VERTEX_SHADER:   []byte(textVertShader + "\x00"),
		gl.FRAGMENT_SHADER: []byte(textFragShader + "\x00"),
	}
)

const (
	defaultFragShader = `
#version 410

layout(location = 0) in vec3 Color;

layout(location = 0) out vec4 diffuseColor;

void main() {
    diffuseColor = vec4(Color.r, Color.g, Color.b, 1.0);
}
`
	defaultVertShader = `
#version 410

layout(location = 0) in vec2 position;
layout(location = 1) in vec3 color;

layout(location = 0) out vec3 Color;

void main() {
    Color = color;
    gl_Position = vec4(position, 0.0, 1.0);
}
`
	textFragShader = `
#version 410

layout(location = 0) in vec3 Color;

layout(location = 0) out vec4 diffuseColor;

void main() {
    diffuseColor = vec4(Color.r, Color.g, Color.b, 1.0);
}
`
	textVertShader = `
#version 410

layout(location = 0) in vec2 position;
layout(location = 1) in vec3 color;

layout(location = 0) out vec3 Color;

void main() {
    Color = color;
    gl_Position = vec4(position, -0.1, 1.0);
}
`
)
