package d_types

type Number interface {
	int | int8 | int16 | int32 | int64
	float32 | float64
}

type V2f32 struct {
	X, Y float32
}

func (vf32 V2f32) Add(other V2f32) V2f32 {
	return V2f32{
		X: vf32.X + other.X,
		Y: vf32.Y + other.Y,
	}
}

type V2f64 struct {
	X, Y float32
}
