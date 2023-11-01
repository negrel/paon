package colors

// Color defines a 32-bit RGBA color.
type Color uint32

// ColorFromRGB returns a Color object using the given color channels values.
func ColorFromRGB(r, g, b uint8) Color {
	return ColorFromRGBA(r, g, b, 255)
}

// ColorFromRGBA returns a Color object using the given color channels values.
func ColorFromRGBA(r, g, b, a uint8) Color {
	return ColorFromHexa(uint32(a)<<24 | uint32(r)<<16 | uint32(g)<<8 | uint32(b))
}

// ColorFromHex returns a Color object using the given hexadecimal color code.
// The color is interpreted as OxRRGGBB
func ColorFromHex(rgb uint32) Color {
	return ColorFromHexa((rgb | 0xFF000000))
}

// ColorFromHexa returns a Color object using the given hexadecimal color code.
// The color is interpreted as OxAARRGGBB
func ColorFromHexa(rgba uint32) Color {
	return Color(rgba)
}

// ColorUnset returns an unset Color.
// Editing the RGB channels of an unset color is useless.
func ColorUnset() Color {
	return Color(0)
}

// Hex converts this Color to an uint32. The color is stored as 0xRRGGBB
func (c Color) Hex() uint32 {
	return uint32(c) & 0xFFFFFF
}

// Hexa converts this Color to an uint32. The color is stored as 0xAARRGGBB
func (c Color) Hexa() uint32 {
	return uint32(c)
}

// R returns the value of the red channel.
func (c Color) R() uint8 {
	return uint8(c & 0xFF0000 >> 16)
}

// G returns the value of the green channel.
func (c Color) G() uint8 {
	return uint8(c & 0xFF00 >> 8)
}

// B returns the value of the blue channel.
func (c Color) B() uint8 {
	return uint8(c & 0xFF)
}

// A returns the value of the alpha channel.
func (c Color) A() uint8 {
	return uint8(c & 0xFF000000 >> 24)
}
