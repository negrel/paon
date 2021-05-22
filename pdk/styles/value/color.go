package value

// Color define a 24-bit RGB color.
type Color struct {
	R, G, B uint8
}

// ColorFromRGB returns a Color object using the given color channel value.
func ColorFromRGB(r, g, b uint8) Color {
	return Color{
		R: r,
		G: g,
		B: b,
	}
}

// ColorFromHex returns a Color object using the given hexadecimal color code.
func ColorFromHex(rgb int32) Color {
	return Color{
		R: uint8((rgb & 0xFF0000) >> 16),
		G: uint8((rgb & 0xFF00) >> 8),
		B: uint8(rgb & 0xFF),
	}
}

// Int32 converts this Color to an int32.
func (c Color) Int32() int32 {
	return int32(c.R)<<16 | int32(c.G)<<8 | int32(c.B)
}
