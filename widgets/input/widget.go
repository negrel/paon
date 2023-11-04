package input

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	treevents "github.com/negrel/paon/tree/events"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/span"
)

type Widget struct {
	*widgets.BaseWidget
	value string
}

func New(defaultValue string) *Widget {
	w := &Widget{
		value: defaultValue,
	}

	w.BaseWidget = widgets.NewBaseWidget(
		widgets.Wrap(w),
		widgets.LayoutFunc(
			func(co layout.Constraint) geometry.Size {
				size := span.Layout(w.value, co)
				return geometry.NewSize(size.Width()+1, size.Height())
			},
		),
		widgets.DrawerFunc(func(surface draw.Surface) {
			span.Draw(surface, w.value, w.Style().CellStyle)

			surface.Set(geometry.NewVec2D(surface.Size().Width()-1, 0), draw.Cell{
				Style: draw.CellStyle{
					Foreground:    0,
					Background:    0,
					Bold:          false,
					Blink:         true,
					Reverse:       true,
					Underline:     false,
					Dim:           false,
					Italic:        false,
					StrikeThrough: false,
				},
				Content: ' ',
			})
		}),
	)

	keyPressEventType, keyPressHandler := keypress.Listener(func(event keypress.Event) {
		switch event.Key {
		case keypress.KeyRune:
			w.value += string(event.Rune)

		case keypress.KeyDelete:
			fallthrough
		case keypress.KeyDEL:
			if w.value != "" {
				w.value = w.value[:len(w.value)-1]
			}
		}
	})

	w.AddEventListener(treevents.LifeCycleEventListener(func(event treevents.LifeCycleEvent) {
		switch event.Stage {
		case treevents.LCSMounted:
			w.Root().Unwrap().(events.Target).AddEventListener(keyPressEventType, keyPressHandler)

		case treevents.LCSBeforeUnmount:
			w.Root().Unwrap().(events.Target).RemoveEventListener(keyPressEventType, keyPressHandler)
		}
	}))

	return w
}

func (w *Widget) Value() string {
	return w.value
}

func (w *Widget) SetValue(v string) {
	w.value = v
}
