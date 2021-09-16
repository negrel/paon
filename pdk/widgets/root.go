package widgets

import (
	"errors"

	"github.com/negrel/paon/internal/metrics"
	"github.com/negrel/paon/pdk/render"
	"github.com/negrel/paon/pdk/tree"
)

// Root define the root of a widget tree.
type Root struct {
	*BaseWidget

	child Widget
}

var _ Layout = &Root{}

// NewRoot returns a new Widget that can be used as a root.
func NewRoot() *Root {
	root := &Root{}

	root.BaseWidget = newBaseWidget(
		initialLCS(LCSMounted),
		Wrap(root),
		NodeConstructor(func(data interface{}) tree.Node { return tree.NewRoot(data) }),
	)
	root.BaseWidget.root = root

	return root
}

// AppendChild implements the Layout interface.
func (r *Root) AppendChild(child Widget) error {
	if r.child == nil {
		r.SetChild(child)
		return nil
	}

	return errors.New("can't append child, root can only have one child")
}

// InsertBefore implements the Layout interface.
func (r *Root) InsertBefore(reference, newChild Widget) error {
	if reference != nil {
		return errors.New("can't insert child, the given reference must be nil on a root node")
	}

	return r.AppendChild(newChild)
}

// RemoveChild implements the Layout interface.
func (r *Root) RemoveChild(child Widget) error {
	if r.child.IsSame(child) {
		r.SetChild(nil)
		return nil
	}

	return errors.New("can't remove child, the widget is not a child of the root")
}

// FirstChild implements the Layout interface.
func (r *Root) FirstChild() Widget {
	return r.child
}

// LastChild implements the Layout interface.
func (r *Root) LastChild() Widget {
	return r.child
}

// Root returns itself to implements the Widget interface.
func (r *Root) Root() *Root {
	return r
}

// SetChild sets the direct child of the root.
// If a child is already present, it is unmounted.
func (r *Root) SetChild(child Widget) {
	if oldChild := r.child; oldChild != nil {
		oldChild.Node().SetParent(nil)
		oldChild.DispatchEvent(NewLifeCycleEvent(oldChild, LCSUnmounted))
	}

	childNode := child.Node()
	childNode.SetParent(r.Node())
	childNode.SetPrevious(nil)
	childNode.SetNext(nil)

	r.child = child
	child.DispatchEvent(NewLifeCycleEvent(child, LCSMounted))
}

// Render implements the render.Renderable interface.
func (r *Root) Render(ctx render.Context) {
	metrics.StartRenderTimer()
	defer metrics.StopRenderTimer()
}
