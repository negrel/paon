package value

type Color struct {
	R, G, B uint8
	isSet   bool
}

func ColorFromRGB(r, g, b uint8) Color {
	return Color{
		R:     r,
		G:     g,
		B:     b,
		isSet: true,
	}
}

func ColorFromHex(rgb int32) Color {
	return Color{
		R:     uint8((rgb & 0xFF0000) >> 16),
		G:     uint8((rgb & 0xFF00) >> 8),
		B:     uint8(rgb & 0xFF),
		isSet: true,
	}
}

// Int32 converts this Color to an int32.
func (c Color) Int32() int32 {
	return int32(c.R)<<16 | int32(c.G)<<8 | int32(c.B)
}

// IsSet returns whether the color is BLACK or just unset.
func (c Color) IsSet() bool {
	return c.isSet
}
