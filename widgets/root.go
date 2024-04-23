package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
)

var (
	NeedRenderEventType = events.NewType("NeedRender")
	ScheduleEventType   = events.NewType("Schedule")
)

// NeedRenderEventListener returns an events.Listener that will call the given handler
// on resize events.
func NeedRenderEventListener(handler func(events.Event)) (events.Type, events.Listener) {
	return NeedRenderEventType, events.NewListenerFunc(handler)
}

var _ Widget = &Root{}

// Root define root of widgets tree.
type Root struct {
	BaseLayout
}

// NewRoot returns a new Root widget that wraps the given Widget and its
// render.Renderable.
func NewRoot(w Widget) *Root {
	root := &Root{}
	root.BaseLayout = NewBaseLayout(root)

	root.AddEventListener(events.ResizeListener(func(_ events.Event, _ events.ResizeEventData) {
		root.DispatchEvent(events.NewEvent(NeedRenderEventType, nil))
	}))

	root.node.AppendChild(nodeOrNil(w))

	return root
}

// Layout implements layout.Layout.
func (r *Root) Layout(co layout.Constraint) geometry.Size {
	r.ChildrenLayout.Reset()

	child := widgetOrNil(r.node.FirstChild())
	if child != nil {
		childSize := child.Layout(co)

		r.ChildrenLayout.Append(ChildLayout{
			Widget: child,
			Bounds: geometry.Rectangle{
				Origin:   geometry.Vec2D{},
				RectSize: childSize,
			},
		})

		return childSize
	}

	return geometry.Size{}
}

// Draw implements draw.Drawer.
func (r *Root) Draw(srf draw.Surface) {
	child := widgetOrNil(r.node.FirstChild())
	if child != nil {
		child.Draw(srf)
	}
}
