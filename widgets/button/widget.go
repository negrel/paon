package button

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events/click"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/span"
)

type Widget struct {
	*widgets.BaseWidget
	text string
}

func New(text string, onclick func(click.Event)) *Widget {
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
			span.Draw(w.text, surface)
		}),
	)

	w.AddEventListener(click.Listener(onclick))

	return w
}
