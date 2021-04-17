package widgets

import (
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/flows"
	"github.com/negrel/paon/pkg/pdk/id"
	"github.com/negrel/paon/pkg/pdk/render"
	pdkstyles "github.com/negrel/paon/pkg/pdk/styles"
)

// Widget define any object part of the Widget tree
// that can be rendered in the screen.
type Widget interface {
	fmt.Stringer
	id.Identifiable
	tree.Node
	events.Target
	pdkstyles.Stylable
	render.Renderable

	// Drawer returns the drawer of this Widget.
	Drawer() draw.Drawer

	FlowAlgo() flows.Algorithm

	// Box returns the current flows.BoxModel of this Widget.
	Box() flows.BoxModel

	// Root returns the Root Widget in the tree.
	Root() *Root

	// Parent returns the Layout that contain this Widget in the tree.
	Parent() Layout

	// Next returns the next sibling of this Widget.
	Next() Widget

	// Previous returns the previous sibling of this Widget.
	Previous() Widget
}

var _ Widget = &widget{}

type widget struct {
	tree.Node
	events.Target

	*flows.Cache
	drawer                 draw.Drawer
	name                   string
	needReflow, needRedraw bool
	theme                  pdkstyles.Theme
}

// NewWidget returns a new Widget object customized with the given Option.
func NewWidget(name string, opts ...Option) Widget {
	return newWidget(name, tree.NewNode(), opts...)
}

func newWidget(name string, node tree.Node, opts ...Option) *widget {
	w := &widget{
		Node:       node,
		Target:     events.MakeTarget(),
		Cache:      flows.NewCache(),
		name:       name,
		needRedraw: true,
		needReflow: true,
	}

	for _, opt := range opts {
		opt(w)
	}

	if w.theme == nil {
		w.theme = pdkstyles.NewTheme(pdkstyles.NewStyle())
	}

	return w
}

// String implements fmt.Stringer interface.
func (w *widget) String() string {
	return fmt.Sprintf("#%s-%d", w.name, w.ID())
}

// Root implements the Widget interface.
func (w *widget) Root() *Root {
	if r := w.RootNode(); r != nil {
		return r.(*Root)
	}

	return nil
}

// Parent implements the Widget interface.
func (w *widget) Parent() Layout {
	if p := w.ParentNode(); p != nil {
		return p.(Layout)
	}
	return nil
}

// Next implements the Widget interface.
func (w *widget) Next() Widget {
	if n := w.NextNode(); n != nil {
		return n.(Widget)
	}

	return nil
}

// Previous implements the Widget interface.
func (w *widget) Previous() Widget {
	if p := w.PreviousNode(); p != nil {
		return p.(Widget)
	}

	return nil
}

// ParentStyle implements the styles.Stylable interface.
func (w *widget) ParentStyle() pdkstyles.Style {
	if parent := w.Parent(); parent != nil {
		return parent.Style()
	}

	return nil
}

// Style implements the styles.Stylable interface.
func (w *widget) Style() pdkstyles.Style {
	return w.theme
}

func (w *widget) Drawer() draw.Drawer {
	return w
}

// Draw implements the Widget interface.
func (w *widget) Draw(ctx draw.Context) {
	assert.NotNil(w.drawer)

	DrawBoxOf(w, ctx)
	w.drawer.Draw(
		ctx.SubContext(
			w.Box().ContentBox(),
		),
	)

	w.needRedraw = false
}

func (w *widget) FlowAlgo() flows.Algorithm {
	return w.flowAlgo
}

func (w *widget) flowAlgo(constraint flows.Constraint) flows.BoxModel {
	assert.NotNil(w.Cache.Algorithm)

	w.needReflow = false
	return w.Cache.Flow(constraint)
}

func (w *widget) enqueue() {
	if root := w.Root(); root != nil {
		root.engine.Enqueue(w)
	}
}

// NeedRendering implements the render.Renderable interface.
func (w *widget) NeedRendering() bool {
	return w.needRedraw || w.needReflow
}

// Render implements the render.Renderable interface.
func (w *widget) Render() render.Patch {
	if w.needReflow {
		return w.Parent().Render()
	}

	if w.needRedraw {

		ctx := w.Root().screen.Canvas().NewContext(w.Box().MarginBox())
		w.Draw(ctx)
		ctx.Commit()

		return ctx.Canvas().Patch()
	}

	return nil
}

func (w *widget) needRender() {
	w.needReflow = true
	w.needRedraw = true
	w.Cache.Invalidate()
	w.enqueue()
}
