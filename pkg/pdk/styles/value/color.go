package value

type Color struct {
	R uint8
	G uint8
	B uint8
}

func ColorFromRGB(r, g, b uint8) Color {
	return Color{
		R: r,
		G: g,
		B: b,
	}
}

func ColorFromHex(rgb int32) Color {
	return Color{
		R: uint8(rgb & 0xFF0000),
		G: uint8(rgb & 0xFF00),
		B: uint8(rgb & 0xFF),
	}
}
