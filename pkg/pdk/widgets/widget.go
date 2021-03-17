package widgets

import (
	"fmt"
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/flows"
	"github.com/negrel/paon/pkg/pdk/id"
	pdkstyles "github.com/negrel/paon/pkg/pdk/styles"
)

// Widget define any object part of the Widget tree
// that can be rendered in the screen.
type Widget interface {
	fmt.Stringer
	id.Identifiable
	tree.Node
	events.Target
	pdkstyles.Stylable
	flows.Flowable
	draw.Drawable

	// Box returns the current flows.BoxModel of this Widget.
	Box() flows.BoxModel

	// Root returns the root Widget in the tree.
	Root() Root

	// Parent returns the Layout that contain this Widget in the tree.
	Parent() Layout

	// Next returns the next sibling of this Widget.
	Next() Widget

	// Previous returns the previous sibling of this Widget.
	Previous() Widget
}

var _ Widget = &widget{}

type widget struct {
	tree.Node
	events.Target
	*flows.Cache
	draw.Drawable

	theme pdkstyles.Theme
}

// NewWidget returns a new Widget object customized with the given Option.
func NewWidget(opts ...Option) Widget {
	return newWidget(tree.NewNode(), opts...)
}

func newWidget(node tree.Node, opts ...Option) *widget {
	w := &widget{
		Node:   node,
		Target: events.MakeTarget(),
	}

	for _, opt := range opts {
		opt(w)
	}

	if w.theme == nil {
		w.theme = pdkstyles.NewTheme(pdkstyles.NewStyle())
	}

	return w
}

// String implements fmt.Stringer interface.
func (w *widget) String() string {
	return string(w.ID())
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

// ParentStyle implements the styles.Stylable interface.
func (w *widget) ParentStyle() pdkstyles.Style {
	if parent := w.Parent(); parent != nil {
		return parent.Style()
	}

	return nil
}

// Style implements the styles.Stylable interface.
func (w *widget) Style() pdkstyles.Style {
	return w.theme
}

// Draw implements the Widget interface.
func (w *widget) Draw(ctx draw.Context) {
	assert.NotNil(w.Drawable)
	DrawBoxOf(w, ctx)
	w.Drawable.Draw(ctx)
}
