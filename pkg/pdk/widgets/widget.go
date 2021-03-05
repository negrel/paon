package widgets

import (
	"fmt"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/flows"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/widgets/themes"
)

// Widget define any object part of this Widget tree
// that can be rendered in the screen.
type Widget interface {
	fmt.Stringer
	tree.Node
	events.Target
	themes.Themed

	// ID returns the unique generated ID or the given one using the ID Option.
	ID() string

	// Root returns the root Widget in the tree.
	Root() Root

	// Parent returns the Layout that contain this Widget in the tree.
	Parent() Layout

	// Next returns the next sibling of this Widget.
	Next() Widget

	// Previous returns the previous sibling of this Widget.
	Previous() Widget

	// Flow returns the flows.BoxModel of this Widget.
	Flow(flows.Constraint) flows.BoxModel

	// Draw draws this widget using the given draw.Context.
	Draw(draw.Context)
}

var _ Widget = &widget{}

type widget struct {
	tree.Node
	events.Target

	id    string
	theme themes.Theme
}

// NewWidget returns a new Widget object customized with the given Option.
func NewWidget(opts ...Option) Widget {
	w := newWidget(tree.NewNode())

	for _, opt := range opts {
		opt(w)
	}

	return w
}

func newWidget(node tree.Node) *widget {
	w := &widget{
		Node:   node,
		Target: events.MakeTarget(),
	}
	w.theme = themes.New(func() themes.Themed { return w.Parent() })

	return w
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

// Theme implements the themes.Themed interface.
func (w *widget) Theme() themes.Theme {
	return w.theme
}

// Flow implements the Widget interface.
func (w *widget) Flow(constraint flows.Constraint) flows.BoxModel {
	return flows.GetFor(w.Theme()).Apply(w.Theme(), constraint)
}

// Draw implements the Widget interface.
func (w *widget) Draw(context draw.Context) {
	panic("implement me")
}
