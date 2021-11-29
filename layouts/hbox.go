package layouts

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/layout"
	pdkmath "github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/pdk/render"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
)

var _ pdkwidgets.Layout = &HBox{}

// HBox is a layout that align contained widgets horizontally.
type HBox struct {
	*pdkwidgets.BaseLayout
}

// NewHBox returns a new HBox layout.
func NewHBox() *HBox {
	hbox := &HBox{}

	hbox.BaseLayout = pdkwidgets.NewBaseLayout(
		pdkwidgets.WidgetOptions(
			pdkwidgets.Wrap(hbox),
		),
	)

	return hbox
}

// AppendChild implements the pdkwidgets.Layout interface.
func (hb *HBox) AppendChild(child pdkwidgets.Widget) error {
	err := hb.BaseLayout.AppendChild(child)
	if err != nil {
		pdkwidgets.NeedRender(hb.BaseLayout.BaseWidget, render.NeedDrawState.Merge(render.NeedLayoutState))
	}

	return err
}

// InsertBefore implements the pdkwidgets.Layout interface.
func (hb *HBox) InsertBefore(reference, child pdkwidgets.Widget) error {
	err := hb.BaseLayout.InsertBefore(reference, child)
	if err != nil {
		pdkwidgets.NeedRender(hb.BaseLayout.BaseWidget, render.NeedDrawState.Merge(render.NeedLayoutState))
	}

	return err
}

// RemoveChild implements the pdkwidgets.Layout interface.
func (hb *HBox) RemoveChild(child pdkwidgets.Widget) error {
	err := hb.BaseLayout.RemoveChild(child)
	if err != nil {
		pdkwidgets.NeedRender(hb.BaseLayout.BaseWidget, render.NeedDrawState.Merge(render.NeedLayoutState))
	}

	return err
}

// Render implements the render.Renderable interface.
func (hb *HBox) Render(ctx render.Context) {
	state := ctx.StateModifier.Merge(pdkwidgets.RenderState(hb.BaseLayout.BaseWidget))
	if state.IsClean() {
		return
	}

	renderObj := hboxRenderObject{
		HBox:  hb,
		ctx:   ctx,
		child: hb.FirstChild(),
		freeSpace: geometry.Rectangle{
			Min: geometry.Vec2D{},
			Max: geometry.NewVec2D(ctx.Constraint.MaxSize.Width(), ctx.Constraint.MaxSize.Height()),
		},
	}

	hb.render(&renderObj)
}

type hboxRenderObject struct {
	*HBox
	ctx render.Context

	child         pdkwidgets.Widget
	width, height int
	freeSpace     geometry.Rectangle
}

func (hbro *hboxRenderObject) Layout(s geometry.Sized) geometry.Rectangle {
	child := hbro.child

	size := s.Size()
	childBounds := geometry.Rectangle{
		Min: hbro.freeSpace.Min,
		Max: hbro.freeSpace.Min.Add(geometry.NewVec2D(size.Width(), size.Height())),
	}
	hbro.freeSpace.Min = geometry.NewVec2D(childBounds.Max.X(), hbro.freeSpace.Min.Y())
	hbro.width += size.Width()
	hbro.height = pdkmath.Max(hbro.height, size.Height())

	// Render the next child before returning the rectangle
	// to know the position of hbox.
	hbro.child = child.NextSibling()
	hbro.render(hbro)

	// Place children rectangle at the right place.
	return childBounds.MoveBy(hbro.HBox.cache.Position())
}

func (hb *HBox) render(renderObj *hboxRenderObject) {
	child := renderObj.child

	// All child are rendered
	// The size of the hbox is known
	if child == nil {
		// Layout the hbox
		rect := renderObj.ctx.Layout.Layout(
			geometry.NewSize(renderObj.width, renderObj.height),
		)

		box := layout.NewBox(rect.Size())
		renderObj.HBox.cache.Store(renderObj.ctx.Constraint, &box)
		renderObj.HBox.cache.SetPosition(rect.Min)
		return
	}

	childCtx := render.Context{
		Flags:  renderObj.ctx.Flags,
		Layer:  renderObj.ctx.Layer,
		Layout: renderObj,
		Constraint: layout.Constraint{
			MinSize:    geometry.Size{},
			MaxSize:    renderObj.freeSpace.Size(),
			ParentSize: geometry.Size{},
			RootSize:   renderObj.ctx.RootSize,
		},
	}

	// Render the children
	// If it call layout and occupy some space
	// the hboxLayout.Layout method will be called.
	child.Render(childCtx)

	// if child didn't call Layout and isn't the last child
	if child.IsSame(renderObj.child) && !child.IsSame(hb.LastChild()) {
		renderObj.child = child.NextSibling()
		hb.render(renderObj)
	}
}
