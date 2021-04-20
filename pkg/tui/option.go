package tui

import (
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/widgets"
)

type Option = widgets.Option

func Props(props ...property.Property) widgets.Option {
	return widgets.Props(props...)
}
