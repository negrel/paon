package widgets

import (
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// DrawBoxOf draws the box of the given widget (border, padding and background color)
// The content box must be drawn by the widget itself.
func DrawBoxOf(widget Widget, ctx draw.Context) {
	style := widget.Style()
	box := widget.Box()

	// Background color
	bgColor, ok := style.Get(property.BackgroundColor()).(property.Color)
	if ok {
		ctx.SetFillColor(bgColor.Color)
	}
	ctx.FillRectangle(box.PaddingBox())
}

// MarkAsNeedRedraw marks the widget so that on the next rendering frame,
// the widget will be redrawed.
func MarkAsNeedRedraw(w Widget) {
	w.markAsNeedRedraw()
}

// MarkAsNeedReflow marks the widget so that on the next rendering frame,
// the widget size will be recomputed. The widgets is also as needing
// a redraw.
func MarkAsNeedReflow(w Widget) {
	w.markAsNeedReflow()
}
