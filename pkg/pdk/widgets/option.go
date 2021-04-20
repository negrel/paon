package widgets

import (
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/flows"
	pdkstyle "github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/widgets/lifecycle"
)

type Option func(widget *widget)

func PrependOptions(opts []Option, toPrepend ...Option) []Option {
	return append(toPrepend, opts...)
}

// Bind binds the given variable to the widget.
func Bind(variable *Widget) Option {
	return func(widget *widget) {
		*variable = widget
	}
}

func Algo(algorithm flows.Algorithm) Option {
	return func(widget *widget) {
		widget.Cache.Algorithm = algorithm
	}
}

func DrawerFn(fn draw.DrawerFn) Option {
	return Drawer(fn)
}

func Drawer(drawable draw.Drawer) Option {
	return func(widget *widget) {
		widget.drawer = drawable
	}
}

// DefaultStyle applies the given styles to the widget.
func DefaultStyle(style pdkstyle.Style) Option {
	return func(widget *widget) {
		widget.theme = pdkstyle.NewTheme(style)
	}
}

// Props applies the given properties to the widget theme. This theme is not shareable
// and can only be modified using Set/Del method on the theme of the widget.
func Props(props ...property.Property) Option {
	return func(widget *widget) {
		if widget.theme == nil {
			widget.theme = pdkstyle.NewTheme(pdkstyle.NewStyle())
		}

		for _, prop := range props {
			widget.theme.Set(prop)
		}
	}
}

func lifecycleHook(step lifecycle.Step, hook func()) Option {
	return func(widget *widget) {
		widget.lifeCycleHooks[step] = hook
	}
}

func BeforeCreate(hook func()) Option {
	return func(widget *widget) {
		hook()
	}
}

func Created(hook func()) Option {
	return lifecycleHook(lifecycle.Created, hook)
}

func BeforeMount(hook func()) Option {
	return lifecycleHook(lifecycle.BeforeMount, hook)
}

func Mounted(hook func()) Option {
	return lifecycleHook(lifecycle.Mounted, hook)
}

func BeforeUpdate(hook func()) Option {
	return lifecycleHook(lifecycle.BeforeUpdate, hook)
}

func Updated(hook func()) Option {
	return lifecycleHook(lifecycle.Updated, hook)
}

func BeforeUnmount(hook func()) Option {
	return lifecycleHook(lifecycle.BeforeUnmount, hook)
}

func Unmounted(hook func()) Option {
	return lifecycleHook(lifecycle.Unmounted, hook)
}
