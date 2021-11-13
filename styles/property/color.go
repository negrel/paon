package property

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/pdk/id"
)

// ColorID define a uniqueID to accross colors properties.
type ColorID id.ID

var (
	colorRegistry = id.Registry{}
	colorMap      = id.NewStringMap()
)

// NewColorID returns a new unique color property ID.
func NewColorID(name string) ColorID {
	id := colorRegistry.New()
	colorMap.Set(id, name)

	return ColorID(id)
}

// String implements the fmt.Stringer interface.
func (ci ColorID) String() string {
	return colorMap.Get(id.ID(ci))
}

// ColorIDCount returns the number of ColorID generated.
func ColorIDCount() uint32 {
	_ = tcell.Color100
	return uint32(colorRegistry.Last())
}

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

var (
	_IDBorderColor     = NewColorID("border-color")
	_IDBackgroundColor = NewColorID("background-color")
	_IDForegroundColor = NewColorID("foreground-color")
)

// BorderColor returns ColorID of the "border-color" property.
func BorderColor() ColorID {
	return _IDBorderColor
}

// BackgroundColor returns ColorID of the "background-color" property.
func BackgroundColor() ColorID {
	return _IDBackgroundColor
}

// ForegroundColor returns ColorID of the "foreground-color" property.
func ForegroundColor() ColorID {
	return _IDForegroundColor
}
