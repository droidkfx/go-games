package shape

type Mesh struct {
	Verts []MeshVert
	Elems [][]uint32
}

func (m *Mesh) OffsetElems(offset uint32) {
	for _, elemL := range m.Elems {
		for i := 0; i < len(elemL); i++ {
			elemL[i] = elemL[i] + offset
		}
	}
}

type MeshVert interface {
	GetVertData() []float32
}

var _ MeshVert = (*GenericMeshVert)(nil)

type GenericMeshVert struct {
	Data []float32
}

func (g *GenericMeshVert) GetVertData() []float32 {
	return g.Data
}
