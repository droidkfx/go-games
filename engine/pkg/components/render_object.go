package components

// RenderObject is a marker interface to allow the render system to take in any object that wishes to be rendered
type RenderObject interface {
	GameObject
	Type() RenderType
}

type RenderType int

//goland:noinspection GoUnusedConst
const (
	RenderType_UNKNOWN RenderType = iota
	RenderType_MESH
)

type MeshRender interface {
	GetMeshData() ([]float32, []uint32)
}

var _ RenderObject = (*MeshRenderObject)(nil)
var _ MeshRender = (*MeshRenderObject)(nil)

type MeshRenderObject struct{}

func (m *MeshRenderObject) GetMeshData() ([]float32, []uint32) {
	return make([]float32, 0), make([]uint32, 0)
}

func (m *MeshRenderObject) Type() RenderType {
	return RenderType_MESH
}
