package widgets

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/tree"
)

type Widget interface {
	tree.Node
	events.EventTarget

	ParentL() Layout
	NextW() Widget
	PreviousW() Widget

	layer() *render.Layer
	Render(bounds geometry.Rectangle) render.Patch
}

var _ Widget = &widget{}

type widget struct {
	tree.Node
	events.Target
	*render.Layer
}

func NewWidget(name string, opts ...Option) Widget {
	node := tree.NewNode(name)
	return newWidget(node, opts...)
}

func newWidget(node tree.Node, opts ...Option) *widget {
	w := &widget{
		Node:  node,
		Layer: render.NewLayer(node),
	}

	for _, opt := range opts {
		opt(w)
	}

	return w
}

func (w *widget) ParentL() Layout {
	if p := w.Parent(); p != nil {
		return w.Parent().(Layout)
	}
	return nil
}

func (w *widget) NextW() Widget {
	if n := w.Next(); n != nil {
		return n.(Widget)
	}

	return nil
}

func (w *widget) PreviousW() Widget {
	if p := w.Previous(); p != nil {
		return p.(Widget)
	}

	return nil
}

func (w *widget) layer() *render.Layer {
	return w.Layer
}
