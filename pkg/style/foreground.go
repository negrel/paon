package style

import (
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

func FgColor(color value.Color) property.Color {
	return property.MakeColor(property.IDForegroundColor, color)
}
