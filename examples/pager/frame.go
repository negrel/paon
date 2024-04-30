package main

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

// Frame is our pager frame.
type Frame struct {
	widgets.BaseLayout
}

// NewFrame returns a new frame layout.
func NewFrame(lines []string) *Frame {
	f := &Frame{}
	f.BaseLayout = widgets.NewBaseLayout(f)

	scroll := NewScroll(lines)

	err := f.AppendChild(scroll)
	if err != nil { // should never happen
		panic(err)
	}

	// Key press handler.
	f.AddEventListener(events.KeyListener(func(ev events.Event, data events.KeyEventData) {
		if data.Rune == 'j' {
			scroll.ScrollDown()
		} else if data.Rune == 'k' {
			scroll.ScrollUp()
		}
	}))
	f.AddEventListener(events.MouseEventListener(func(ev events.Event, data events.MouseEventData) {
		if data.Buttons&events.WheelDown != 0 {
			scroll.ScrollDown()
		} else if data.Buttons&events.WheelUp != 0 {
			scroll.ScrollUp()
		}
	}))

	return f
}

// Layout implements widgets.Layout.
func (f *Frame) Layout(co layout.Constraint, ctx widgets.LayoutContext) layout.SizeHint {
	scroll := f.FirstChild()

	// Compute scroll view size.
	scrollSize := scroll.Layout(co, ctx)
	// Store it so BaseLayout.Draw knows it size when widgets will be drawn.
	f.ChildrenLayout.Append(widgets.ChildLayout{
		Widget: scroll,
		Bounds: geometry.Rectangle{
			Origin:   geometry.Vec2D{},
			RectSize: scrollSize.Size,
		},
	})

	// Use all available space.
	return layout.SizeHint{
		MinSize:          co.MaxSize,
		Size:             co.MaxSize,
		VerticalPolicy:   layout.ExpandingSizePolicy,
		HorizontalPolicy: layout.ExpandingSizePolicy,
	}
}
