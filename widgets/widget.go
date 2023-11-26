package widgets

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/render"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/tree"
)

// Widget is a generic interface that define any component part of the widget/element tree.
// Any types that implement the Widget interface can be added to the widget tree.
// You must embed *BaseWidget or *BoxWidget to implement this interface as it
// contains private methods.
type Widget interface {
	events.Target
	styles.Styled
	tree.NodeAccessor
	render.RenderableAccessor

	setNode(*tree.Node)

	// Swap swaps this widget node with node of the given one.
	Swap(Widget)
}

// IsMounted return whether a Node is mounted on an active Widget tree.
func IsMounted(n *tree.Node) bool {
	root := n.Root()
	if root == nil {
		return false
	}

	_, isRoot := root.Unwrap().(Root)
	return isRoot
}

// Private events.Target for private embedding.
type eventTarget = events.Target

var _ Widget = &PanicWidget{}

// PanicWidget define a minimal (and incomplete) Widget type.
// PanicWidget implements panic methods for styles.Styled and render.RenderableAccessor
// interfaces.
type PanicWidget struct {
	eventTarget

	node *tree.Node
}

// NewPanicWidget returns a new PanicWidget object configured with
// the given options.
func NewPanicWidget(nodeData Widget) *PanicWidget {
	widget := newBaseWidget(nodeData)

	return widget
}

func newBaseWidget(nodeData Widget) *PanicWidget {
	widget := &PanicWidget{
		eventTarget: events.NewTarget(),
		node:        tree.NewNode(nodeData),
	}

	return widget
}

// Swap implements Widget.
func (bw *PanicWidget) Swap(other Widget) {
	// Swap node reference.
	old := bw.node
	bw.node = other.Node()
	other.setNode(old)

	// Swap data contained in tree.Node.
	this := old.Swap(other.Node().Unwrap())
	bw.node.Swap(this)
}

// Node implements the Widget interface.
func (bw *PanicWidget) Node() *tree.Node {
	return bw.node
}

func (bw *PanicWidget) setNode(node *tree.Node) {
	bw.node = node
}

// Renderable implements Widget.
func (*PanicWidget) Renderable() render.Renderable {
	panic("unimplemented")
}

// Style implements styles.Styled.
func (*PanicWidget) Style() styles.Style {
	panic("unimplemented")
}
