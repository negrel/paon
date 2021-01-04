package widgets

import (
	"fmt"
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/pdk/styles"
)

// Widget define any object part of the Widget tree
// that can be rendered in the screen.
type Widget interface {
	fmt.Stringer
	tree.Node
	events.Target
	render.Object
	styles.Stylised

	// Theme return the Widget Theme.
	Theme() Theme

	// ID return the unique generated ID or the given one using the ID Option.
	ID() string

	// Root return the root Widget in the tree.
	Root() Root

	// Parent return the Layout that contain this Widget in the tree.
	Parent() Layout

	// Next return the next sibling of the Widget.
	Next() Widget

	// Previous return the previous sibling of the Widget.
	Previous() Widget
}

var _ Widget = &widget{}

type widget struct {
	tree.Node
	events.Target

	theme theme

	id string
}

// NewWidget return a new Widget object customized with the given Option.
func NewWidget(opts ...Option) Widget {
	w := newWidget(tree.NewNode())

	for _, opt := range opts {
		opt(w)
	}

	return w
}

func newWidget(node tree.Node) *widget {
	return &widget{
		Node:   node,
		Target: events.MakeTarget(),
	}
}

// ID implements the Widget interface.
func (w *widget) ID() string {
	return w.id
}

// String implements fmt.Stringer interface.
func (w *widget) String() string {
	return w.ID()
}

// Root implements the Widget interface.
func (w *widget) Root() Root {
	if r := w.RootNode(); r != nil {
		return r.(Root)
	}

	return nil
}

// Parent implements the Widget interface.
func (w *widget) Parent() Layout {
	if p := w.ParentNode(); p != nil {
		return p.(Layout)
	}
	return nil
}

// Next implements the Widget interface.
func (w *widget) Next() Widget {
	if n := w.NextNode(); n != nil {
		return n.(Widget)
	}

	return nil
}

// Previous implements the Widget interface.
func (w *widget) Previous() Widget {
	if p := w.PreviousNode(); p != nil {
		return p.(Widget)
	}

	return nil
}

// Style implements the styles.Stylised interface.
func (w *widget) Style() styles.Style {
	return w.theme
}

// Theme return the theme of the widget.
func (w *widget) Theme() Theme {
	return w.theme
}

// ParentObject implements the render.Object interface.
func (w *widget) ParentObject() render.Object {
	return w.Parent()
}
