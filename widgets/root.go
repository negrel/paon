package widgets

import (
	"fmt"

	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
)

var (
	NeedRenderEventType = events.NewType("NeedRender")
)

// NeedRenderEventListener returns an events.Listener that will call the given handler
// on resize events.
func NeedRenderEventListener(handler func(events.Event)) (events.Type, events.Listener) {
	return NeedRenderEventType, events.NewListenerFunc(handler)
}

var _ Widget = &Root{}

// Root is a special enhancer type that is always the root of a widget tree.
type Root struct {
	BaseEnhancer
}

// NewRoot returns a new Root widget that wraps the given Widget and its
// render.Renderable.
func NewRoot(w Widget) *Root {
	root := &Root{}
	root.BaseEnhancer = NewBaseEnhancer(root)

	root.AddEventListener(events.ResizeListener(func(_ events.Event, _ events.ResizeEventData) {
		root.DispatchEvent(events.NewEvent(NeedRenderEventType, nil))
	}))

	_, err := root.SwapWidget(w)
	if err != nil {
		panic(fmt.Errorf("failed to add widget to root: %w", err))
	}

	return root
}

// DispatchEvent forwards event to child after calling all of root's listeners.
func (r *Root) DispatchEvent(ev events.Event) {
	r.Target.DispatchEvent(ev)

	// Forward events
	w := r.Widget()
	if w != nil {
		w.DispatchEvent(ev)
	}
}

// Layout implements layout.Layout.
func (r *Root) Layout(co layout.Constraint, ctx LayoutContext) layout.SizeHint {
	child := widgetOrNil(r.node.FirstChild())
	if child == nil {
		r.ChildLayout.Widget = nil
		return layout.SizeHint{
			MinSize:          geometry.Size{},
			Size:             geometry.Size{},
			VerticalPolicy:   layout.IgnoreFlag,
			HorizontalPolicy: layout.IgnoreFlag,
		}
	}

	ctx.ParentSize = ctx.RootSize
	childSizeHint := child.Layout(co, ctx)

	// Store child layout.
	r.ChildLayout.Widget = child
	r.ChildLayout.Bounds = geometry.Rectangle{
		Origin:   geometry.Vec2D{},
		RectSize: childSizeHint.Size,
	}

	if childSizeHint.ExpandingDirections().Horizontal() {
		r.ChildLayout.Bounds.RectSize = childSizeHint.Size.WithWidth(ctx.RootSize.Width)
	}
	if childSizeHint.ExpandingDirections().Vertical() {
		r.ChildLayout.Bounds.RectSize = childSizeHint.Size.WithHeight(ctx.RootSize.Height)
	}

	// Ignored...
	return childSizeHint
}
