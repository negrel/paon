package style

import (
	"github.com/negrel/paon/pkg/pdk/style/property"
	"github.com/negrel/paon/pkg/pdk/style/value"
)

func BgColor(color value.Color) property.Color {
	return property.MakeColor(property.IDBackgroundColor, color)
}

// TODO: add BgGradient property constructor
