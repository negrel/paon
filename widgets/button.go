package widgets

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
)

type Button struct {
	*pdkwidgets.BaseWidget
	text string
}

func NewButton(text string, onclick func(events.Click)) *Button {
	w := &Button{
		text: text,
	}

	w.BaseWidget = pdkwidgets.NewBaseWidget(
		pdkwidgets.Wrap(w),
		pdkwidgets.LayoutFunc(func(co layout.Constraint) geometry.Size {
			return LayoutSpan(w.text, co)
		}),
		pdkwidgets.DrawerFunc(func(surface draw.Surface) {
			DrawSpan(w.text, surface)
		}),
	)

	w.AddEventListener(events.ClickListener(onclick))

	return w
}
