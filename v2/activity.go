package ginger

import (
	"image"

	"github.com/negrel/ginger/v2/events"
	"github.com/negrel/ginger/v2/widgets"
)

// Activity correspond to a view of your
// application.
type Activity struct {
	active bool

	root *widgets.Root
}

// NewActivity return a new activity.
func NewActivity(widget widgets.Widget) *Activity {
	ac := &Activity{
		active: false,
		root:   widgets.NewRoot(widget),
	}

	events.Emitter.AddClickListener(ac)
	events.Emitter.AddResizeListener(ac)

	return ac
}

/*****************************************************
 ********************** Events ***********************
 *****************************************************/
// ANCHOR Events

// OnClick implements the events.ClickListener interface.
func (ac *Activity) OnClick(ce *events.ClickEvent) {
	if ac.root.Attached() {
		for _, child := range ac.root.ItemAt(ce.Position()) {
			if listener, ok := child.(events.ClickListener); ok {
				listener.OnClick(ce)
			}
		}
	}
}

// OnResize implements the events.ResizeListener interface.
func (ac *Activity) OnResize(re *events.ResizeEvent) {
	if ac.active {
		ac.root.Refresh()
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Active return wether or not the activity is active
func (ac *Activity) Active() bool {
	return ac.root.Attached()
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Start the activity
func (ac *Activity) start(size func() (w, h int)) {
	// Set refresh function and refresh
	ac.root.Refresh = func() {
		width, height := size()
		renderer <- ac.root.Render(image.Rect(0, 0, width, height))
	}
	ac.root.Refresh()
}

// Stop the activity
func (ac *Activity) stop() {
	ac.root.Refresh = nil
}
