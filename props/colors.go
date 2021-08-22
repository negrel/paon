package props

import (
	"github.com/negrel/paon/styles/property"
	"github.com/negrel/paon/styles/value"
)

// FgColor returns a property.Color with the given color
// and a property.ForegroundColorID.
func FgColor(color value.Color) property.Color {
	return property.NewColor(property.ForegroundColorID(), color)
}

// BgColor returns a property.Color with the given color
// and a property.BackgroundColorID.
func BgColor(color value.Color) property.Color {
	return property.NewColor(property.BackgroundColorID(), color)
}
