package d_types

type Number interface {
	int | int8 | int16 | int32 | int64
	float32 | float64
}

type V2f32 struct {
	X, Y float32
}

type V2f64 struct {
	X, Y float32
}

func (v V2f32) Eq(other V2f32) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v V2f32) Add(other V2f32) V2f32 {
	return V2f32{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v V2f32) DivC(c float32) V2f32 {
	return V2f32{
		X: v.X / c,
		Y: v.Y / c,
	}
}
