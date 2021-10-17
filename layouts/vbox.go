package layouts

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/layout"
	pdkmath "github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/pdk/render"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
)

var _ pdkwidgets.Layout = &VBox{}

// VBox is a layout that align contained widgets vertically.
type VBox struct {
	*pdkwidgets.BaseLayout

	cache *layout.Cache
}

// NewVBox returns a new VBox layout.
func NewVBox() *VBox {
	vbox := &VBox{
		// Cache is used as a store.
		cache: layout.NewCache(nil),
	}

	vbox.BaseLayout = pdkwidgets.NewBaseLayout(
		pdkwidgets.WidgetOptions(
			pdkwidgets.Wrap(vbox),
		),
	)

	return vbox
}

// Render implements the render.Renderable interface.
func (vb *VBox) Render(ctx render.Context) {
	needRedraw := pdkwidgets.IsDirty(vb.BaseLayout.BaseWidget) || ctx.Flags >= render.DrawFlag
	needReflow := !vb.cache.IsValid(ctx.Constraint) || ctx.Flags >= render.FullRenderFlag
	if !needRedraw && !needReflow {
		return
	}

	renderObj := vboxRenderObject{
		VBox:  vb,
		ctx:   ctx,
		child: vb.FirstChild(),
		freeSpace: geometry.Rectangle{
			Min: geometry.Vec2D{},
			Max: geometry.NewVec2D(ctx.Constraint.MaxSize.Width(), ctx.Constraint.MaxSize.Height()),
		},
	}

	vb.render(&renderObj)
}

type vboxRenderObject struct {
	*VBox
	ctx render.Context

	child         pdkwidgets.Widget
	width, height int
	freeSpace     geometry.Rectangle
}

func (vbro *vboxRenderObject) Layout(s geometry.Sized) geometry.Rectangle {
	child := vbro.child

	size := s.Size()
	childBounds := geometry.Rectangle{
		Min: vbro.freeSpace.Min,
		Max: vbro.freeSpace.Min.Add(geometry.NewVec2D(size.Width(), size.Height())),
	}
	vbro.freeSpace.Min = geometry.NewVec2D(vbro.freeSpace.Min.X(), childBounds.Max.Y())
	vbro.width = pdkmath.Max(size.Width(), vbro.width)
	vbro.height += size.Height()

	// Render the next child before returning the rectangle
	// to know the position of vbox.
	vbro.child = child.NextSibling()
	vbro.render(vbro)

	// Place children rectangle at the right place.
	return childBounds.MoveBy(vbro.VBox.cache.Position())
}

func (vb *VBox) render(renderObj *vboxRenderObject) {
	child := renderObj.child

	// All child are rendered
	// The size of the vbox is known
	if child == nil || renderObj.freeSpace.Empty() {
		// Layout the vbox
		rect := renderObj.ctx.Layout.Layout(
			geometry.NewSize(renderObj.width, renderObj.height),
		)

		box := layout.NewBox(rect.Size())
		renderObj.VBox.cache.Store(renderObj.ctx.Constraint, box)
		renderObj.VBox.cache.SetPosition(rect.Min)
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
	// the vboxLayout.Layout method will be called.
	child.Render(childCtx)

	// if child didn't call Layout and isn't the last child
	// we force the rendering of the next sibling
	if child.IsSame(renderObj.child) && !child.IsSame(vb.LastChild()) {
		renderObj.child = child.NextSibling()
		vb.render(renderObj)
	}
}
