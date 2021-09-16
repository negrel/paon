package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/render"
	"github.com/negrel/paon/pdk/tree"
	"github.com/negrel/paon/styles"
)

// Widget is a generic interface that define any component part of the widget/element tree.
// A widget is basically a renderable object in a tree. It support events and styling/theming.
//
// Any types that implements the Widget interface can be added to the widget tree. The recommended
// way to create custom widgets is using the BaseWidget implementation.
type Widget interface {
	id.Identifiable
	events.Target
	styles.Styled
	styles.Themed
	render.Renderable

	// LifeCycleStage returns the current LifeCycleStage of this Widget.
	LifeCycleStage() LifeCycleStage

	// Node returns a tree.Node containing this Widget.
	Node() tree.Node

	// Parent returns the parent Layout of this Node.
	Parent() Layout

	// Root returns the root of the widget tree.
	Root() *Root

	// NextSibling returns the next sibling widget in the parent child list.
	NextSibling() Widget

	// PreviousSibling returns the previous sibling widget in the parent child list.
	PreviousSibling() Widget
}

func widgetOrNil(n tree.Node) Widget {
	if n == nil {
		return nil
	}

	return n.Unwrap().(Widget)
}

func nodeOrNil(w Widget) tree.Node {
	if w == nil {
		return nil
	}

	return w.Node()
}

var _ Widget = &BaseWidget{}

// Type alias to avoid collision with Widget.Node method.
type node = tree.Node

// BaseWidget define a basic Widget object that implements the Widget interface.
// BaseWidget can either be used alone (see NewBaseWidget for the required options)
// or in composite struct.
// BaseWidget takes care of the following things for you:
// - O(1) access time to the root.
// - Caching the layout.BoxModel
// - Updating the lifecycle stage
// - Limit the draw area of the context to the widget border box.
type BaseWidget struct {
	node node
	events.Target

	root       *Root
	stage      LifeCycleStage
	theme      styles.Theme
	renderable render.Renderable
}

// NewBaseWidget returns a new BaseWidget object configured with
// the given options.
// The LayoutAlgo and Drawer widget options are required.
// To embed this layout in composite struct, use the Wrap widget options.
func NewBaseWidget(options ...WidgetOption) *BaseWidget {
	widget := newBaseWidget(options...)
	widget.AddEventListener(LifeCycleEventListener(func(event LifeCycleEvent) {
		widget.stage = event.Stage

		// Update root field on mount/unmount
		switch event.Stage {
		case LCSMounted:
			root := widget.Parent().Root()
			assert.NotNil(root)
			widget.root = root

			widget.needRender()

		case LCSBeforeUnmount:
			// Clean node before unmounting (drop layer in layer tree)

		case LCSUnmounted:
			widget.root = nil
			widget.cache.Invalidate()
		}
	}))

	widget.AddEventListener(NeedRenderListener(func(nre NeedRenderEvent) {
		widget.needRender()
	}))

	return widget
}

func newBaseWidget(options ...WidgetOption) *BaseWidget {
	widget := &BaseWidget{
		stage: LCSInitial,
	}

	widgetConf := &baseWidgetOption{
		BaseWidget:      widget,
		nodeConstructor: tree.NewLeafNode,
		data:            widget,
		defaultStyle:    styles.New(),
	}

	for _, option := range options {
		option(widgetConf)
	}

	widget.node = widgetConf.nodeConstructor(widgetConf.data)

	if widget.Target == nil {
		widget.Target = events.NewTarget()
	}

	for _, listener := range widgetConf.listeners {
		widget.AddEventListener(listener)
	}

	if widget.theme == nil {
		widget.theme = styles.NewTheme(styles.NewWeighted(widgetConf.defaultStyle, -1))
	}

	return widget
}

// ID implements the id.Identifiable interface.
func (bw *BaseWidget) ID() id.ID {
	return bw.node.ID()
}

// IsSame implements the id.Identifiable interface.
func (bw *BaseWidget) IsSame(other id.Identifiable) bool {
	return bw.node.IsSame(other)
}

// Widget is an helper method that returns the Widget wrapped by the internal tree.Node.
// This method act kinda like the `this` keyword in Java.
func (bw *BaseWidget) Widget() Widget {
	return widgetOrNil(bw.node)
}

// Node implements the Widget interface.
func (bw *BaseWidget) Node() tree.Node {
	return nodeWrapper{
		Node:   bw.node,
		widget: bw.Widget(),
	}
}

// LifeCycleStage implements the Widget interface
func (bw *BaseWidget) LifeCycleStage() LifeCycleStage {
	return bw.stage
}

// Theme implements the styles.Themed interface.
func (bw *BaseWidget) Theme() styles.Theme {
	return bw.theme
}

// Style implements the styles.Styled interface.
func (bw *BaseWidget) Style() styles.Style {
	return bw.theme
}

// Parent implements the Widget interface.
func (bw *BaseWidget) Parent() Layout {
	if parent := bw.node.Parent(); parent != nil {
		return parent.Unwrap().(Layout)
	}

	return nil
}

// Root implements the Widget interface.
func (bw *BaseWidget) Root() *Root {
	return bw.root
}

// NextSibling implements the Widget interface.
func (bw *BaseWidget) NextSibling() Widget {
	return widgetOrNil(bw.node.Next())
}

// PreviousSibling implements the Widget interface.
func (bw *BaseWidget) PreviousSibling() Widget {
	return widgetOrNil(bw.node.Previous())
}

func (bw *BaseWidget) needRender() {
	if bw.Parent() == nil {
		return
	}

	bw.Parent().DispatchEvent(NewNeedRenderEvent(bw))
}

// Render implements the render.Renderable interface.
func (bw *BaseWidget) Render(ctx render.Context) {
	bw.renderable.Render(ctx)
}

type nodeWrapper struct {
	tree.Node
	widget Widget
}

func (nw nodeWrapper) SetParent(parent tree.Node) {
	if parent == nil {
		assert.NotNil(nw.widget.Root())
		nw.widget.DispatchEvent(NewLifeCycleEvent(nw.widget, LCSBeforeUnmount))
	} else {
		parentW := parent.Unwrap().(Widget)
		stage := LCSBeforeMount
		if parentW.Root() == nil { // parent is not mounted
			stage = LCSBeforeUnmount
		}
		nw.widget.DispatchEvent(NewLifeCycleEvent(nw.widget, stage))
	}

	nw.Node.SetParent(parent)
}
