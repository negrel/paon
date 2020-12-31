package styles

type ColorValue struct {
	R uint8
	G uint8
	B uint8
}

func RGB(r, g, b uint8) ColorValue {
	return ColorValue{
		R: r,
		G: g,
		B: b,
	}
}

func RGBFromHex(rgb int32) ColorValue {
	return ColorValue{
		R: uint8(rgb & 0xFF0000),
		G: uint8(rgb & 0xFF00),
		B: uint8(rgb & 0xFF),
	}
}
