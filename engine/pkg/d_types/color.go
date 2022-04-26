package d_types

//goland:noinspection GoSnakeCaseUsage,GoUnusedGlobalVariable
var (
	Color_RED   = ColorRGB{R: 1.0}
	Color_BLUE  = ColorRGB{B: 1.0}
	Color_GREEN = ColorRGB{G: 1.0}
)

type ColorRGB struct {
	R, G, B float32
}
