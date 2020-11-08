package widgets

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/tree"
)

// Layout is a Widget that can contain child Widget.
type Layout interface {
	tree.ParentNode
	Widget

	FirstChildW() Widget
	LastChildW() Widget

	AppendChildW(child Widget) error
	InsertBeforeW(reference, child Widget) error
	RemoveChildW(child Widget)
}

type layout struct {
	tree.ParentNode
	events.Target
	*render.Layer
}

func NewLayout(name string, opts ...Option) Layout {
	l := &layout{
		ParentNode: tree.NewParent(name),
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

func (l *layout) ParentL() Layout {
	if p := l.ParentNode.Parent(); p != nil {
		return l.ParentNode.Parent().(Layout)
	}
	return nil
}

func (l *layout) NextW() Widget {
	if n := l.Next(); n != nil {
		return n.(Widget)
	}

	return nil
}

func (l *layout) PreviousW() Widget {
	if p := l.Previous(); p != nil {
		return p.(Widget)
	}

	return nil
}

func (l *layout) layer() *render.Layer {
	return l.Layer
}

func (l *layout) FirstChildW() Widget {
	if fc := l.FirstChild(); fc != nil {
		return fc.(Widget)
	}

	return nil
}

func (l *layout) LastChildW() Widget {
	if lc := l.LastChild(); lc != nil {
		return lc.(Widget)
	}

	return nil
}

func (l *layout) AppendChildW(child Widget) error {
	return l.AppendChild(child)
}

func (l *layout) InsertBeforeW(reference, child Widget) error {
	return l.InsertBefore(reference, child)
}

func (l *layout) RemoveChildW(child Widget) {
	l.RemoveChildW(child)
}
