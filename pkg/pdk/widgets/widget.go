package widgets

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/tree"
)

type node tree.Node

type Widget interface {
	node
	events.EventTarget

	Root() Root
	Parent() Layout
	Next() Widget
	Previous() Widget
}

var _ Widget = &widget{}

type widget struct {
	tree.Node
	events.Target
}

func NewWidget(name string, opts ...Option) Widget {
	w := newWidget(tree.NewNode(name))

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
