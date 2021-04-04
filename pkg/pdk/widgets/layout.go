package widgets

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/pdk/id"
)

// Layout is a Widget that can contain child Widget.
type Layout interface {
	tree.ParentNode
	Widget

	// Returns the first child Widget of this Layout.
	FirstChild() Widget
	// Returns the last child Widget of this Layout.
	LastChild() Widget

	// Adds the given Widget at the end of the child list.
	// An error is returned if the given Widget is an ancestor
	// of this Layout.
	// This function panic if a nil value is given.
	AppendChild(child Widget) (Widget, error)

	// Inserts the given Widget before the reference Widget in the child list.
	// If the reference is nil the child is appended.
	// An error is returned if the given Widget is an ancestor
	// of this Layout.
	// This function panic if a nil child value is given.
	InsertBefore(reference, child Widget) (Widget, error)

	// Removes the given Widget of the child list of this Layout.
	// An error is returned if the Widget is not a direct child of
	// this Layout.
	RemoveChild(child Widget) error
}

type parentNode = tree.ParentNode

var _ Layout = &layout{}

type layout struct {
	*widget
	parentNode
}

func NewLayout(name string, ptr tree.ParentNode, opts ...Option) Layout {
	return newLayout(name, ptr, opts...)
}

func newLayout(name string, ptr tree.ParentNode, opts ...Option) *layout {
	parent := tree.NewCompositeParent(ptr)
	l := &layout{
		widget:     newWidget(name, parent, opts...),
		parentNode: parent,
	}

	return l
}

// ID implements the id.Identifiable interface.
func (l *layout) ID() id.ID {
	return l.widget.ID()
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
func (l *layout) AppendChild(child Widget) (Widget, error) {
	log.Debugln("appending", child, "in", l, "child list")
	err := l.AppendChildNode(child)

	if err == nil {
		l.needRender()
	}

	return child, err
}

// InsertBefore implements the Layout interface.
func (l *layout) InsertBefore(reference, child Widget) (Widget, error) {
	log.Debugln("inserting", child, "before", reference, "in", l)
	err := l.InsertBeforeNode(reference, child)

	if err == nil {
		l.needRender()
	}

	return child, err
}

// RemoveChild implements the Layout interface.
func (l *layout) RemoveChild(child Widget) error {
	log.Debugln("removing", child, "from", l)
	err := l.RemoveChildNode(child)

	if err == nil {
		l.needRender()
	}

	return err
}
