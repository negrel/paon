package widgets

import (
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

	// Add the given Widget at the end of the children list.
	// An error is returned if the given Widget is an ancestor
	// of this Layout.
	// This function panic if a nil value is given.
	AppendChild(child Widget) (Widget, error)

	// Insert the given Widget before the reference Widget in the children list.
	// If the reference is nil the child is appended.
	// An error is returned if the given Widget is an ancestor
	// of this Layout.
	// This function panic if a nil child value is given.
	InsertBefore(reference, child Widget) (Widget, error)

	// Remove the given Widget of the children list of this Layout.
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
	err := l.AppendChildNode(child)

	return child, err
}

// InsertBefore implements the Layout interface.
func (l *layout) InsertBefore(reference, child Widget) (Widget, error) {
	err := l.InsertBeforeNode(reference, child)

	return child, err
}

// RemoveChild implements the Layout interface.
func (l *layout) RemoveChild(child Widget) error {
	return l.RemoveChildNode(child)
}

//var _ Widget = &childWidget{}
//
//type childWidget struct {
//	Widget
//	box        *flows.Box
//	constraint flows.Constraint
//}
//
//func wrapWidget(widget Widget) *childWidget {
//	cw, ok := widget.(*childWidget)
//	if ok {
//		return cw
//	}
//
//	return &childWidget{
//		Widget:     widget,
//		box:        nil,
//		constraint: flows.Constraint{},
//	}
//}
//
//// flow implements the Widget interface.
//func (cw *childWidget) flow(constraint flows.Constraint) *flows.Box {
//	if cw.constraint.Equals(constraint) && cw.box != nil {
//		return cw.box
//	}
//
//	cw.constraint = constraint
//	box := cw.Widget.flow(constraint)
//	cw.box = box
//
//	assert.True(box.Width() > constraint.Min.Width())
//	assert.True(box.Width() < constraint.Max.Width())
//	assert.True(box.Height() > constraint.Min.Height())
//	assert.True(box.Height() < constraint.Max.Height())
//
//	return box
//}
//
//// Box implements the Widget interface.
//func (cw *childWidget) Box() flows.BoxModel {
//	return cw.box
//}
