package shape

const (
	sqrt2 = 1.4142135
)

// func NewEqTriAt(px, py, size, r, g, b float32) Mesh {
// 	c2 := size * 0.5
// 	c1 := c2 * sqrt2
// 	return NewTriFromPoints(px, py+c1, px+c2, py-c2, px-c1)
// }
//
// func NewTriFromPoints(p1x, p1y, p2x, p2y, p3x, p3y, r, g, b float32) Mesh {
// 	return Mesh{
// 		Verts: []float32{
// 			p1x, p1y, r, g, b,
// 			p2x, p2y, r, g, b,
// 			p3x, p3y, r, g, b,
// 		},
// 		Elems: []uint32{0, 1, 2},
// 	}
// }
