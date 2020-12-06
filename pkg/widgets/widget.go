package widgets

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/styles"
)

type Widget interface {
	tree.Node
	events.EventTarget
	render.Object

	Parent() Layout
	Next() Widget
	Previous() Widget
}

var _ Widget = &widget{}

type widget struct {
	tree.Node
	events.Target

	styles.Style
}

func NewWidget(name string, opts ...Option) Widget {
	node := tree.NewNode(name)
	w := newWidget(node)

	for _, opt := range opts {
		opt(w)
	}

	return w
}

func newWidget(node tree.Node) *widget {
	return &widget{
		Node:   node,
		Target: events.MakeTarget(),
		Style:  styles.New(),
	}
}

func (w *widget) Parent() Layout {
	if p := w.ParentNode(); p != nil {
		return p.(Layout)
	}
	return nil
}

func (w *widget) ParentObject() render.Object {
	return w.Parent()
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
