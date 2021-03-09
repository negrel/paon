package style

import (
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

func BgColor(color value.Color) property.Color {
	return property.MakeColor(property.BackgroundColor(), color)
}

// TODO: add BgGradient property constructor
