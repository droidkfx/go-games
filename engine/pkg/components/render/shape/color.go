package shape

import "github.com/droidkfx/go-games/engine/pkg/d_types"

type ColorMap func(float32, float32) d_types.ColorRGB

func Color(loc d_types.V2f32, cmap ColorMap) d_types.ColorRGB {
	return cmap(loc.X, loc.Y)
}

func ConstColor(color d_types.ColorRGB) ColorMap {
	return func(_ float32, _ float32) d_types.ColorRGB {
		return color
	}
}
