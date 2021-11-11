package widgets

import (
	pdkevents "github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/render"
	"github.com/negrel/paon/pdk/tree"
	treevents "github.com/negrel/paon/pdk/tree/events"
	"github.com/negrel/paon/styles"
)

// Widget is a generic interface that define any component part of the widget/element tree.
// A widget is basically a renderable object in a tree. It support events and styling/theming.
//
// Any types that implements the Widget interface can be added to the widget tree. The recommended
// way to create custom widgets is using the BaseWidget implementation.
type Widget interface {
	id.Identifiable
	pdkevents.Target
	styles.Styled
	styles.Themed
	render.Renderable

	LifeCycleStage() treevents.LifeCycleStage

	// Node returns a tree.Node containing this Widget.
	Node() treevents.Node

	// Parent returns the parent Layout of this Node.
	Parent() Layout

	// Root returns the root of the widget tree.
	Root() *Root

	// NextSibling returns the next sibling widget in the parent child list.
	NextSibling() Widget

	// PreviousSibling returns the previous sibling widget in the parent child list.
	PreviousSibling() Widget
}

type globalListener struct {
	pdkevents.Type
	pdkevents.Listener
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

// NeedRender marks the widget (and its ancestors) as dirty so it will be rendered
// on the next update.
func NeedRender(bw *BaseWidget) {
	bw.needRender()
}

// IsDirty returns true if the given BaseWidget is tagged as dirty.
func IsDirty(bw *BaseWidget) bool {
	return bw.isDirty
}

// MarkAsClean marks the widget as clean which means that
// it is up to date and don't need to be rendered.
func MarkAsClean(bw *BaseWidget) {
	bw.isDirty = false
}

var _ Widget = &BaseWidget{}

// Type alias to avoid collision with Widget.Node method.
type node = treevents.Node

// BaseWidget define a basic Widget object that implements the Widget interface.
// BaseWidget can either be used alone (see NewBaseWidget for the required options)
// or in composite struct.
type BaseWidget struct {
	node

	isDirty         bool
	theme           styles.Theme
	globalListeners []globalListener
}

// NewBaseWidget returns a new BaseWidget object configured with
// the given options.
// To embed this layout in composite struct, use the Wrap widget options.
func NewBaseWidget(options ...WidgetOption) *BaseWidget {
	widget := newBaseWidget(options...)

	widget.AddEventListener(treevents.LifeCycleEventListener(func(lce treevents.LifeCycleEvent) {
		switch lce.Stage {
		case treevents.LCSBeforeUnmount:
			for _, l := range widget.globalListeners {
				widget.Root().RemoveEventListener(l.Type, l.Listener)
			}

		case treevents.LCSMounted:
			widget.needRender()

		}
	}))

	widget.AddEventListener(NeedRenderListener(func(nre NeedRenderEvent) {
		widget.needRender()
	}))

	return widget
}

var emptyStyle = styles.New(pdkevents.NewNoOpTarget())

func newBaseWidget(options ...WidgetOption) *BaseWidget {
	widget := &BaseWidget{
		globalListeners: make([]globalListener, 0, 8),
	}

	widgetConf := &baseWidgetOption{
		BaseWidget:  widget,
		nodeOptions: make([]treevents.NodeOption, 0, 8),
	}

	for _, option := range options {
		option(widgetConf)
	}

	widget.node = treevents.NewBaseNode(widgetConf.nodeOptions...)

	if widgetConf.defaultStyle == nil {
		widgetConf.defaultStyle = emptyStyle
	}

	if widget.theme == nil {
		widget.theme = styles.NewTheme(styles.NewWeighted(widgetConf.defaultStyle, -1))
	}

	return widget
}

// ID implements the id.Identifiable interface.
func (bw BaseWidget) ID() id.ID {
	return bw.node.ID()
}

// IsSame implements the id.Identifiable interface.
func (bw BaseWidget) IsSame(other id.Identifiable) bool {
	return bw.node.IsSame(other)
}

// Widget is an helper method that returns the Widget wrapped by the internal tree.Node.
// This method act kinda like the `this` keyword in Java.
func (bw BaseWidget) Widget() Widget {
	return widgetOrNil(bw.node)
}

// Node implements the Widget interface.
func (bw BaseWidget) Node() treevents.Node {
	return bw.node
}

// Theme implements the styles.Themed interface.
func (bw BaseWidget) Theme() styles.Theme {
	return bw.theme
}

// Style implements the styles.Styled interface.
func (bw BaseWidget) Style() styles.Style {
	return bw.theme
}

// Parent implements the Widget interface.
func (bw BaseWidget) Parent() Layout {
	if parent := bw.node.Parent(); parent != nil {
		return parent.Unwrap().(Layout)
	}

	return nil
}

// Root implements the Widget interface.
func (bw BaseWidget) Root() *Root {
	return bw.node.Root().Unwrap().(*Root)
}

// NextSibling implements the Widget interface.
func (bw BaseWidget) NextSibling() Widget {
	return widgetOrNil(bw.node.Next())
}

// PreviousSibling implements the Widget interface.
func (bw BaseWidget) PreviousSibling() Widget {
	return widgetOrNil(bw.node.Previous())
}

func (bw *BaseWidget) needRender() {
	if bw.isDirty || bw.LifeCycleStage() != treevents.LCSMounted {
		return
	}

	bw.isDirty = true
	bw.Parent().DispatchEvent(NewNeedRenderEvent(bw.Widget()))
}

// Layer implements the render.Renderable interface.
func (bw BaseWidget) Layer() *render.Layer {
	panic("Layer is not implemented")
}

// Render implements the render.Renderable interface.
func (bw BaseWidget) Render(ctx render.Context) {
	panic("Render is not implemented")
}

// AddGlobalEventListener adds the given event listener to
// the Root events.Target. The events added with this function
// are automatically removed during the before unmount lifecycle stage.
// This function panics if the widget is not mounted.
func (bw *BaseWidget) AddGlobalEventListener(t pdkevents.Type, l pdkevents.Listener) {
	bw.globalListeners = append(bw.globalListeners, globalListener{t, l})
	bw.Root().AddEventListener(t, l)
}
