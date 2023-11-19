package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/tree"
)

// Widget is a generic interface that define any component part of the widget/element tree.
// Any types that implement the Widget interface can be added to the widget tree.
// You must embed *BaseWidget or *BoxWidget to implement this interface as it
// contains private methods.
type Widget interface {
	events.Target
	layout.Layout
	draw.Drawer
	styles.Styled

	Node() *tree.Node[Widget]
	setNode(*tree.Node[Widget])

	// Swap swaps this widget node with node of the given one.
	Swap(Widget)
}

// Private events.Target for private embedding.
type eventTarget = events.Target

var _ Widget = &BaseWidget{}

// BaseWidget define a basic Widget object that implements the Widget interface.
// BaseWidget can either be used alone (see NewBaseWidget for the required options)
// or in composite struct.
type BaseWidget struct {
	eventTarget

	node *tree.Node[Widget]
}

// NewBaseWidget returns a new BaseWidget object configured with
// the given options.
func NewBaseWidget(nodeData Widget) *BaseWidget {
	widget := newBaseWidget(nodeData)

	return widget
}

func newBaseWidget(nodeData Widget) *BaseWidget {
	widget := &BaseWidget{
		eventTarget: events.NewTarget(),
		node:        tree.NewNode(nodeData),
	}

	return widget
}

// Swap implements Widget.
func (bw *BaseWidget) Swap(other Widget) {
	// Swap node reference.
	old := bw.node
	bw.node = other.Node()
	other.setNode(old)

	// Swap data contained in tree.Node.
	this := old.Swap(other)
	bw.node.Swap(this)
}

// Layout implements layout.Layout.
// This function panics, you must overwrite it.
func (bw *BaseWidget) Layout(co layout.Constraint) geometry.Size {
	panic("Layout must be overrided")
}

// Draw implements draw.Drawer.
// This function panics, you must overwrite it.
func (bw *BaseWidget) Draw(surface draw.Surface) {
	panic("Draw must be overrided")
}

// Style implements styles.Styled.
// BaseWidget has no style, it always return nil. Override this function
// if you widget needs styling.
func (bw *BaseWidget) Style() styles.Style {
	return nil
}

// Node implements the Widget interface.
func (bw *BaseWidget) Node() *tree.Node[Widget] {
	return bw.node
}

func (bw *BaseWidget) setNode(node *tree.Node[Widget]) {
	bw.node = node
}
