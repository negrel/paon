package widgets

import (
	"fmt"
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/pdk/style"
)

type Widget interface {
	fmt.Stringer
	tree.Node
	events.Target
	render.Object
	style.Themed

	ID() string

	Root() Root
	Parent() Layout
	Next() Widget
	Previous() Widget
}

var _ Widget = &widget{}

type widget struct {
	tree.Node
	events.Target

	theme theme

	id string
}

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

func (w *widget) ID() string {
	return w.id
}

func (w *widget) String() string {
	return w.ID()
}

func (w *widget) Root() Root {
	if r := w.RootNode(); r != nil {
		return r.(Root)
	}

	return nil
}

func (w *widget) Parent() Layout {
	if p := w.ParentNode(); p != nil {
		return p.(Layout)
	}
	return nil
}

func (w *widget) Next() Widget {
	if n := w.NextNode(); n != nil {
		return n.(Widget)
	}

	return nil
}

func (w *widget) Previous() Widget {
	if p := w.PreviousNode(); p != nil {
		return p.(Widget)
	}

	return nil
}

func (w *widget) Theme() style.Theme {
	assert.NotNil(w.theme, "%v widget type doesn't instantiate a style.Theme object", w)

	return w.theme
}

func (w *widget) ParentObject() render.Object {
	return w.Parent()
}

func (w *widget) Renderer() render.Renderer {
	panic("implement me")
}
