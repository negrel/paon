package render

import (
	pdkstyle "github.com/negrel/paon/pkg/pdk/styles"
)

// Object define any object that can be rendered.
type Object interface {
	pdkstyle.Stylised

	Render(ctx Context)
}
