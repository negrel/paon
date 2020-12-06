package widgets

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/styles"
)

// Layout is a Widget that can contain child Widget.
type Layout interface {
	tree.ParentNode
	Widget

	FirstChild() Widget
	LastChild() Widget

	AppendChild(child Widget) error
	InsertBefore(reference, child Widget) error
	RemoveChild(child Widget)
}

var _ Layout = &layout{}

type layout struct {
	tree.ParentNode
	events.Target

	styles.Style
}

func NewLayout(name string, opts ...Option) Layout {
	l := &layout{
		ParentNode: tree.NewParent(name),
		Style:      styles.New(),
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

func (l *layout) Parent() Layout {
	if p := l.ParentNode.ParentNode(); p != nil {
		return l.ParentNode.ParentNode().(Layout)
	}
	return nil
}

func (l *layout) ParentObject() render.Object {
	panic("implement me")
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
	l.RemoveChild(child)
}

func (l *layout) Layout(ctx render.Context) {
	l.Style.Layout(ctx)
}

func (l *layout) Draw(ctx render.Context) {
	l.Style.Draw(ctx)
}
