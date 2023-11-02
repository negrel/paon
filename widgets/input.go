package widgets

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
	pdkevents "github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/layout"
	treevents "github.com/negrel/paon/pdk/tree/events"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
)

type Input struct {
	*pdkwidgets.BaseWidget
	value string
}

func NewInput(defaultValue string) *Input {
	w := &Input{
		value: defaultValue,
	}

	w.BaseWidget = pdkwidgets.NewBaseWidget(
		pdkwidgets.Wrap(w),
		pdkwidgets.LayoutFunc(
			func(co layout.Constraint) geometry.Size {
				return LayoutSpan(w.value, co)
			},
		),
		pdkwidgets.DrawerFunc(func(surface draw.Surface) {
			DrawSpan(w.value, surface)
		}),
	)

	keyPressEventType, keyPressHandler := events.KeyPressListener(func(event events.KeyPress) {
		switch event.Key {
		case events.KeyRune:
			w.value += string(event.Rune)

		case events.KeyDelete:
			fallthrough
		case events.KeyDEL:
			if w.value != "" {
				w.value = w.value[:len(w.value)-1]
			}
		}
	})

	w.AddEventListener(treevents.LifeCycleEventListener(func(event treevents.LifeCycleEvent) {
		switch event.Stage {
		case treevents.LCSMounted:
			w.Root().Unwrap().(pdkevents.Target).AddEventListener(keyPressEventType, keyPressHandler)

		case treevents.LCSBeforeUnmount:
			w.Root().Unwrap().(pdkevents.Target).RemoveEventListener(keyPressEventType, keyPressHandler)
		}
	}))

	return w
}

func (i *Input) Value() string {
	return i.value
}

func (i *Input) SetValue(v string) {
	i.value = v
}
