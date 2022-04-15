package render

type Mesh interface {
	BaseComponent
	GetMeshData() ([]float32, []uint32)
}

var _ BaseComponent = (*RawMesh)(nil)
var _ Mesh = (*RawMesh)(nil)

type RawMesh struct{}

func (m *RawMesh) GetMeshData() ([]float32, []uint32) {
	return make([]float32, 0), make([]uint32, 0)
}

func (m *RawMesh) Type() Type {
	return TypeMesh
}
