package widgets

import (
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/flows"
	pdkstyle "github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
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

func lifecycleHook(step LifeCycleStep, hook func()) Option {
	return func(widget *widget) {
		remover := func() {}

		listener := &events.Listener{
			Type: LifeCycleEventType(),
			Handle: func(event events.Event) {
				lifeCycleEvent, isLifeCycleEvent := event.(LifeCycleEvent)
				if isLifeCycleEvent && lifeCycleEvent.Step == step {
					hook()
					remover()
				}
			},
		}

		// Remove the lifecycle hook if the lifecycle won't occure more than one time.
		if step == beforeCreateLifeCycleStep ||
			step == createdLifeCycleStep {
			remover = func() {
				widget.RemoveEventListener(listener)
			}
		}

		widget.AddEventListener(listener)
	}
}

func BeforeCreate(hook func()) Option {
	return func(widget *widget) {
		hook()
	}
}

func Created(hook func()) Option {
	return lifecycleHook(createdLifeCycleStep, hook)
}

func BeforeMount(hook func()) Option {
	return lifecycleHook(beforeMountLifeCycleStep, hook)
}

func Mounted(hook func()) Option {
	return lifecycleHook(mountedLifeCycleStep, hook)
}

func BeforeUpdate(hook func()) Option {
	return lifecycleHook(beforeUpdateLifeCycleStep, hook)
}

func Updated(hook func()) Option {
	return lifecycleHook(updatedLifeCycleStep, hook)
}

func BeforeUnmount(hook func()) Option {
	return lifecycleHook(beforeUnmountLifeCycleStep, hook)
}

func Unmounted(hook func()) Option {
	return lifecycleHook(unmountedLifeCycleStep, hook)
}
