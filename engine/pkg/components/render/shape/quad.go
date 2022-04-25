package shape

import "github.com/droidkfx/go-games/engine/pkg/d_types"

func SolidSquareCenteredAt(size float32, loc d_types.V2f32, color d_types.ColorRGB) Mesh {
	return SquareCenteredAt(size, loc, ConstColor(color))
}

func SquareCenteredAt(size float32, loc d_types.V2f32, cmap ColorMap) Mesh {
	return QuadCenteredAt(d_types.V2f32{X: size, Y: size}, loc, cmap)
}

func SolidQuadCenteredAt(size, loc d_types.V2f32, color d_types.ColorRGB) Mesh {
	return QuadCenteredAt(size, loc, ConstColor(color))
}

func QuadCenteredAt(size, loc d_types.V2f32, cmap ColorMap) Mesh {
	halfSize := size.DivC(2)
	locUL := d_types.V2f32{X: loc.X - halfSize.X, Y: loc.Y + halfSize.Y}
	locLR := d_types.V2f32{X: loc.X + halfSize.X, Y: loc.Y - halfSize.Y}
	return Quad(locUL, locLR, cmap)
}

func SolidQuad(locUL, locLR d_types.V2f32, color d_types.ColorRGB) Mesh {
	return Quad(locUL, locLR, ConstColor(color))
}

func Quad(locUL, locLR d_types.V2f32, cmap ColorMap) Mesh {
	cUL := Color(locUL, cmap)
	locUR := d_types.V2f32{X: locLR.X, Y: locUL.Y}
	cUR := Color(locUR, cmap)
	locLL := d_types.V2f32{X: locUL.X, Y: locLR.Y}
	cLL := Color(locLL, cmap)
	cLR := Color(locLR, cmap)

	return Mesh{
		Verts: []MeshVert{
			&GenericMeshVert{Data: []float32{locUL.X, locUL.Y, cUL.R, cUL.G, cUL.B}}, // UL
			&GenericMeshVert{Data: []float32{locLR.X, locUL.Y, cUR.R, cUR.G, cUR.B}}, // UR
			&GenericMeshVert{Data: []float32{locUL.X, locLR.Y, cLL.R, cLL.G, cLL.B}}, // LL
			&GenericMeshVert{Data: []float32{locLR.X, locLR.Y, cLR.R, cLR.G, cLR.B}}, // LR
		},
		Elems: [][]uint32{
			{0, 3, 2},
			{0, 1, 3},
		},
	}
}
