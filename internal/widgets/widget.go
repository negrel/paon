package widgets

import (
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/tree"
)

type Widget interface {
	tree.Node

	Render(buffer render.Surface)
}

var _ Widget = &widget{}

type widget struct {
	tree.Node

	options []Option
}

func NewWidget(name string, opts ...Option) Widget {
	return newWidget(name, opts...)
}

func newWidget(name string, opts ...Option) *widget {
	w := &widget{
		Node:    tree.NewNode(name),
		options: opts,
	}

	return w
}

func (w *widget) Render(buffer render.Surface) {
	for _, opt := range w.options {
		opt.apply(buffer)
	}
}
