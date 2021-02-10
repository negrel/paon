package draw

import (
	pdkstyle "github.com/negrel/paon/pkg/pdk/styles"
)

// Object defined any object that can be painted.
type Object interface {
	pdkstyle.Stylised

	Paint(ctx Context)
}
