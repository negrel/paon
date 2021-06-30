package value

// Color define a 24-bit RGB color.
type Color struct {
	R, G, B uint8
	isSet   bool
}

// ColorFromRGB returns a Color object using the given color channels values.
func ColorFromRGB(r, g, b uint8) Color {
	return Color{
		R:     r,
		G:     g,
		B:     b,
		isSet: true,
	}
}

// ColorFromHex returns a Color object using the given hexadecimal color code.
func ColorFromHex(rgb int32) Color {
	return Color{
		R:     uint8((rgb & 0xFF0000) >> 16),
		G:     uint8((rgb & 0xFF00) >> 8),
		B:     uint8(rgb & 0xFF),
		isSet: true,
	}
}

// ColorUnset returns an unset Color.
// Editing the RGB channels of an unset color is useless.
func ColorUnset() Color {
	return Color{}
}

// Hex converts this Color to an int32.
func (c Color) Hex() int32 {
	return int32(c.R)<<16 | int32(c.G)<<8 | int32(c.B)
}

// IsSet returns true if the color is set (to make distinction between black and unset colors).
func (c Color) IsSet() bool {
	return c.isSet
}
