package widgets

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/tree"
)

type pNode = tree.ParentNode

// Layout is a Widget that can contain child Widget.
type Layout interface {
	pNode
	Widget

	FirstChild() Widget
	LastChild() Widget

	AppendChild(child Widget) error
	InsertBefore(reference, child Widget) error
	RemoveChild(child Widget)
}

var _ Layout = &layout{}

type layout struct {
	pNode
	events.Target
}

func NewLayout(name string, opts ...Option) Layout {
	l := newLayout(tree.NewParent(name))

	for _, opt := range opts {
		opt(l)
	}

	return l
}

func newLayout(node tree.ParentNode) *layout {
	return &layout{
		pNode:  node,
		Target: events.MakeTarget(),
	}
}

func (l *layout) Root() *root {
	if r := l.RootNode(); r != nil {

	}

	return nil
}

func (l *layout) Parent() Layout {
	if p := l.ParentNode(); p != nil {
		return p.(Layout)
	}
	return nil
}

func (l *layout) ParentObject() render.Object {
	if p := l.ParentNode(); p != nil {
		return p.(render.Object)
	}

	return nil
}

func (l *layout) Next() Widget {
	if n := l.NextNode(); n != nil {
		return n.(Widget)
	}

	return nil
}

func (l *layout) Previous() Widget {
	if p := l.PreviousNode(); p != nil {
		return p.(Widget)
	}

	return nil
}

func (l *layout) FirstChild() Widget {
	if fc := l.FirstChildNode(); fc != nil {
		return fc.(Widget)
	}

	return nil
}

func (l *layout) LastChild() Widget {
	if lc := l.LastChildNode(); lc != nil {
		return lc.(Widget)
	}

	return nil
}

func (l *layout) AppendChild(child Widget) error {
	return l.AppendChildNode(child)
}

func (l *layout) InsertBefore(reference, child Widget) error {
	return l.InsertBeforeNode(reference, child)
}

func (l *layout) RemoveChild(child Widget) {
	_ = l.RemoveChildNode(child)
}
