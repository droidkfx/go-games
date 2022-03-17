#version 410

layout(location = 0) in vec3 Color;

layout(location = 0) out vec4 diffuseColor;

void main() {
    diffuseColor = vec4(Color.r, Color.g, Color.b, 1.0);
}