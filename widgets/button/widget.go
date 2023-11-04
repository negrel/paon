package button

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/span"
)

type Widget struct {
	*widgets.BaseWidget
	text string
}

func New(text string, onclick func(mouse.ClickEvent)) *Widget {
	w := &Widget{
		text: text,
	}

	w.BaseWidget = widgets.NewBaseWidget(
		widgets.Wrap(w),
		widgets.LayoutFunc(
			func(co layout.Constraint) geometry.Size {
				return span.Layout(w.text, co)
			},
		),
		widgets.DrawerFunc(func(surface draw.Surface) {
			span.Draw(surface, w.text, w.Style().CellStyle)
		}),
		widgets.Style(styles.Style{
			CellStyle: draw.CellStyle{
				Foreground:    0,
				Background:    0,
				Bold:          false,
				Blink:         false,
				Reverse:       true,
				Underline:     false,
				Dim:           false,
				Italic:        false,
				StrikeThrough: false,
			},
			Extras: map[string]any{},
		}),
	)

	w.AddEventListener(mouse.PressListener(func(event mouse.PressEvent) {
		style := w.Style()
		style.Reverse = !style.Reverse
	}))
	w.AddEventListener(mouse.UpListener(func(event mouse.UpEvent) {
		style := w.Style()
		style.Reverse = !style.Reverse
	}))

	w.AddEventListener(mouse.ClickListener(onclick))

	return w
}
