package draw

import (
	pdkstyle "github.com/negrel/paon/pkg/pdk/styles"
)

// Object represents thing that can be painted by a Painter.
type Object interface {
	pdkstyle.Stylised
	Drawable
}
