package render

import pdkstyle "github.com/negrel/paon/pkg/pdk/style"

// Object define any object that can be rendered
type Object interface {
	pdkstyle.Themed

	ParentObject() Object
	Renderer() Renderer
}
