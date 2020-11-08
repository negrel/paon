package widgets

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/style"
)

type Widget interface {
	tree.Node
	events.EventTarget

	ParentL() Layout
	NextW() Widget
	PreviousW() Widget

	Theme() *style.Theme
	Render(render.Patch) render.Patch
}

var _ Widget = &widget{}

type widget struct {
	tree.Node
	events.Target

	theme *style.Theme
}

func NewWidget(name string, opts ...Option) Widget {
	node := tree.NewNode(name)
	return newWidget(node, opts...)
}

func newWidget(node tree.Node, opts ...Option) *widget {
	w := &widget{
		Node:   node,
		Target: events.MakeTarget(),
		theme:  style.NewTheme(),
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

func (w *widget) Theme() *style.Theme {
	return w.theme
}

func (w *widget) Render(patch render.Patch) render.Patch {
	w.theme.ApplyOn(&patch)

	return patch
}
