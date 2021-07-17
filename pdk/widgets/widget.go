package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/tree"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

// Widget is a generic interface that define any component part of the widget/element tree.
// Any types that implement the Widget interface can be added to the widget tree. However, it is strongly
// recommended to create custom widgets using the BaseWidget implementation.
type Widget interface {
	id.Identifiable
	events.Target
	draw.Drawer
	layout.Manager
	layout.Boxed
	styles.Styled
	styles.Themed

	// LifeCycleStage returns the current LifeCycleStage of this Widget.
	LifeCycleStage() LifeCycleStage

	// Node returns the underlying tree.Node used by this Widget.
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

// Reflow dispatch a ReflowEvent event to the parent widget.
func Reflow(bw *BaseWidget) {
	if !bw.cache.IsValid() || bw.Parent() == nil {
		return
	}

	bw.reflow()
}

// Redraw enqueue the widget to the redraw queue if
func Redraw(bw *BaseWidget) {
	bw.root.DispatchEvent(NewRedrawEvent(bw.ID(), bw.Widget()))
}

var _ Widget = &BaseWidget{}

// Type alias to avoid collision with Widget.Node method.
type node = tree.Node

// BaseWidget define a basic Widget object that implements the Widget interface.
// BaseWidget can either be used alone (see NewBaseWidget for the required options)
// or in composite struct.
// BaseWidget takes care of the following things for you:
// - Constant access time to the root.
// - Caching the layout.BoxModel
// - Updating the lifecycle stage
// - Limit the draw area of the context to the widget border box.
type BaseWidget struct {
	node node
	events.Target

	root   *Root
	stage  LifeCycleStage
	theme  styles.Theme
	cache  *layout.Cache
	drawer draw.Drawer
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
		if event.Stage == LCSMounted {
			root := widget.Parent().Root()
			assert.NotNil(root)
			widget.root = root
		} else if event.Stage == LCSUnmounted {
			widget.root = nil
		}
	}))

	widget.AddEventListener(ReflowListener(func(re ReflowEvent) {
		widget.reflow()
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

	assert.NotNil(widget.cache)
	assert.NotNil(widget.drawer)

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

// HaveFixedSize returns true if this widget have a fixed size.
func (bw *BaseWidget) HaveFixedSize() bool {
	return bw.theme.Get(property.WidthID()) != nil && bw.theme.Get(property.HeightID()) != nil
}

// Widget is an helper method that returns the Widget wrapped by the internal tree.Node.
// This method act kinda like the `this` keyword in Java.
func (bw *BaseWidget) Widget() Widget {
	return widgetOrNil(bw.node)
}

// Draw implements the draw.Drawer interface.
func (bw *BaseWidget) Draw(ctx draw.Context) {
	boxCtx := draw.SubContext(ctx, bw.Box().BorderBox())
	_ = boxCtx

	contentCtx := ctx.Canvas().NewContext(bw.Box().ContentBox())
	bw.drawer.Draw(contentCtx)
}

// Layout implements the layout.Algo interface.
func (bw *BaseWidget) Layout(c layout.Constraint) layout.BoxModel {
	return bw.cache.Layout(c)
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

// Box implements the Widget interface.
func (bw *BaseWidget) Box() layout.BoxModel {
	return bw.cache.Box()
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

func (bw *BaseWidget) reflow() {
	if bw.HaveFixedSize() {
		bw.root.DispatchEvent(NewReflowEvent(bw.ID(), bw.Widget()))
	} else {
		bw.Parent().DispatchEvent(NewReflowEvent(bw.ID(), bw.Widget()))
	}

	Redraw(bw)
	bw.cache.Invalidate()
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
