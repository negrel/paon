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
		R: uint8((rgb & 0xFF0000) >> 16),
		G: uint8((rgb & 0xFF00) >> 8),
		B: uint8(rgb & 0xFF),
	}
}

func (c Color) Int32() int32 {
	return int32(c.R)<<16 | int32(c.G)<<8 | int32(c.B)
}
