package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/tree"
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

type parentNode = tree.ParentNode

var _ Layout = &layout{}

type layout struct {
	*widget
	parentNode
}

func NewLayout(opts ...Option) Layout {
	l := newLayout(tree.NewParent())

	for _, opt := range opts {
		opt(l.widget)
	}

	return l
}

func newLayout(node tree.ParentNode) *layout {
	return &layout{
		widget:     newWidget(node),
		parentNode: node,
	}
}

// Parent implements the Widget interface.
func (l *layout) Parent() Layout {
	if p := l.ParentNode(); p != nil {
		return p.(Layout)
	}
	return nil
}

// LastChild implements the Layout interface.
func (l *layout) FirstChild() Widget {
	if fc := l.FirstChildNode(); fc != nil {
		return fc.(Widget)
	}

	return nil
}

// LastChild implements the Layout interface.
func (l *layout) LastChild() Widget {
	if lc := l.LastChildNode(); lc != nil {
		return lc.(Widget)
	}

	return nil
}

// AppendChild implements the Layout interface.
func (l *layout) AppendChild(child Widget) error {
	return l.AppendChildNode(child)
}

// InsertBefore implements the Layout interface.
func (l *layout) InsertBefore(reference, child Widget) error {
	return l.InsertBeforeNode(reference, child)
}

// RemoveChild implements the Layout interface.
func (l *layout) RemoveChild(child Widget) {
	err := l.RemoveChildNode(child)
	assert.Nilf(err, "removing %v returned a non-nil error", child)
}
