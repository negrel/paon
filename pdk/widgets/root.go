package widgets

import (
	"errors"

	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/tree"
	"github.com/negrel/paon/styles/value"
)

// Root define the root of a widget tree.
type Root struct {
	*BaseWidget

	child                  Widget
	needReflow, needRedraw bool
	reflowQ                []layout.Manager
	redrawQ                []draw.Drawer
}

var _ Layout = &Root{}

// NewRoot returns a new Widget that can be used as a root.
func NewRoot() *Root {
	root := &Root{}
	root.enqueueReflow()

	root.BaseWidget = newBaseWidget(
		initialLCS(LCSMounted),
		Wrap(root),
		NodeConstructor(func(data interface{}) tree.Node { return tree.NewRoot(data) }),
		LayoutManager(layout.ManagerFn(func(c layout.Constraint) layout.BoxModel {
			return root.child.Layout(c)
		})),
		Drawer(draw.DrawerFn(func(c draw.Context) {
			root.Draw(c)
		})),
		Listeners(
			ReflowListener(func(re ReflowEvent) {
				if root.needReflow {
					// Event is from direct child, we can redraw the entire tree.
					if re.ResourceID == root.ID() {
						root.enqueueReflow()
						return
					}
					root.reflowQ = append(root.reflowQ, re.Manager)
				}
			}),
			RedrawListener(func(re RedrawEvent) {
				if root.needRedraw {
					if re.ResourceID == root.ID() {
						root.enqueueRedraw()
						return
					}
					root.redrawQ = append(root.redrawQ, re.Drawer)
				}
			}),
		),
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

// Draw implements the draw.Drawable interface.
func (r *Root) Draw(ctx draw.Context) {
	ctx.SetFillColor(value.Color{})
	ctx.FillRectangle(ctx.Bounds())
	r.child.Draw(ctx)
	ctx.Commit()
}

func (r *Root) enqueueReflow() {
	r.needReflow = false
	r.reflowQ = []layout.Manager{r}
	r.enqueueRedraw()
}

func (r *Root) enqueueRedraw() {
	r.needRedraw = false
	r.redrawQ = []draw.Drawer{r}
}

// Update flush the reflow & redraw queue and trigger a redraw and a reflow if needed.
func (r *Root) Update(canvas draw.Canvas) {
	lenReflowQ := len(r.reflowQ)
	for i := 0; i < lenReflowQ; i++ {
		r.reflowQ[i].Layout(layout.Constraint{})
	}
	r.reflowQ = r.reflowQ[lenReflowQ:]
	r.needReflow = true

	lenRedrawQ := len(r.redrawQ)
	for i := 0; i < lenRedrawQ; i++ {
		r.redrawQ[i].Draw(canvas.NewContext(canvas.Bounds()))
	}
	r.redrawQ = r.redrawQ[lenRedrawQ:]
	r.needRedraw = true
}
