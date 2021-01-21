package render

import (
	"github.com/negrel/paon/internal/geometry"
	pdkstyle "github.com/negrel/paon/pkg/pdk/styles"
)

// Object define any object that can be rendered.
type Object interface {
	geometry.Sized
	geometry.Positioned
	pdkstyle.Stylised

	ParentObject() Object
}
