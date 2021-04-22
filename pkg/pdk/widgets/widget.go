package widgets

import (
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
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
	events.Target
	pdkstyles.Stylable
	render.Renderable

	// Drawer returns the drawer of this Widget.
	Drawer() draw.Drawer

	// FlowAlgo returns the flow algorithm of this Widget
	FlowAlgo() flows.Algorithm

	// Box returns the current flows.BoxModel of this Widget.
	Box() flows.BoxModel

	// IsSame returns true if the given widget is the same as this Widget.
	IsSame(Widget) bool

	// Root returns the Root Widget in the tree.
	Root() *Root

	// Parent returns the Layout that contain this Widget in the tree.
	Parent() Layout
	setParent(Layout)

	// Next returns the next sibling of this Widget.
	Next() Widget
	setNext(Widget)

	// Previous returns the previous sibling of this Widget.
	Previous() Widget
	setPrevious(Widget)

	// IsDescendantOf return true if this Widget is a descendant of the given Layout.
	IsDescendantOf(layout Layout) bool

	markAsNeedRedraw()
	markAsNeedReflow()
}

var _ Widget = &widget{}

type widget struct {
	events.Target
	*flows.Cache

	nextWidget Widget
	prevWidget Widget
	parent     Layout

	id                     id.ID
	drawer                 draw.Drawer
	name                   string
	needReflow, needRedraw bool
	theme                  pdkstyles.Theme
}

// NewWidget returns a new Widget object customized with the given Option.
func NewWidget(name string, opts ...Option) Widget {
	return newWidget(name, opts...)
}

func newWidget(name string, opts ...Option) *widget {
	assert.Greater(len(name), 0)
	w := &widget{
		Target: events.MakeTarget(),
		Cache:  flows.NewCache(),
		id:     id.Make(),

		name:       name,
		needRedraw: true,
		needReflow: true,
	}

	for _, opt := range opts {
		opt(w)
	}

	w.DispatchEvent(
		MakeLifeCycleEvent(beforeCreateLifeCycleStep),
	)
	defer w.DispatchEvent(
		MakeLifeCycleEvent(createdLifeCycleStep),
	)

	if w.theme == nil {
		w.theme = pdkstyles.NewTheme(pdkstyles.NewStyle())
	}

	return w
}

// String implements fmt.Stringer interface.
func (w *widget) String() string {
	return fmt.Sprintf("#%s-%d", w.name, w.ID())
}

func (w *widget) ID() id.ID {
	return w.id
}

func (w *widget) IsSame(other Widget) bool {
	return w.ID() == other.ID()
}

// Root implements the Widget interface.
func (w *widget) Root() *Root {
	if parent := w.Parent(); parent != nil {
		return parent.Root()
	}

	return nil
}

// Parent implements the Widget interface.
func (w *widget) Parent() Layout {
	return w.parent
}

func (w *widget) setParent(layout Layout) {
	w.parent = layout
}

func (w *widget) IsDescendantOf(layout Layout) bool {
	if layout == nil {
		return false
	}

	var widget Widget = w
	for widget != nil {
		if widget.IsSame(layout) {
			return true
		}

		widget = widget.Parent()
	}

	return false
}

// Next implements the Widget interface.
func (w *widget) Next() Widget {
	return w.nextWidget
}

func (w *widget) setNext(widget Widget) {
	w.nextWidget = widget
}

// Previous implements the Widget interface.
func (w *widget) Previous() Widget {
	return w.prevWidget
}

func (w *widget) setPrevious(widget Widget) {
	w.prevWidget = widget
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
	// if already in the rendering queue
	if w.NeedRendering() {
		return
	}

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
	w.DispatchEvent(MakeLifeCycleEvent(beforeUpdateLifeCycleStep))
	defer w.DispatchEvent(MakeLifeCycleEvent(updatedLifeCycleStep))

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

func (w *widget) markAsNeedReflow() {
	w.Cache.Invalidate()
	w.enqueue()
	w.needReflow = true
	w.needRedraw = true
}

func (w *widget) markAsNeedRedraw() {
	w.enqueue()
	w.needRedraw = true
}
