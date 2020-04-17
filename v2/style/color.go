package style

// Color is an alias for int32. It should be use
// to write hexadecimal color.
// Red: 0xFF0000
// Green: 0x00FF00
// Blue: 0x0000FF
type Color = int32

// Basic color
const (
	RED   Color = 0xFF0000
	GREEN Color = 0x00FF00
	BLUE  Color = 0x0000FF
)

// DefaultColor is used to leave the Color unchanged
// from whatever system or teminal default may exist.
var DefaultColor Color = -1

// RGB return a color object from the given value.
func RGB(r, g, b uint8) Color {
	return Color(
		int32(uint32(r)<<16) + int32(uint16(g)<<8) + int32(b),
	)
}
