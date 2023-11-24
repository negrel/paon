package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/render"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/tree"
)

var _ Widget = &BaseLayout{}

// BaseLayout is a basic layout widget.
// It can either be used alone or embedded within a struct.
// BaseLayout
// - Dispatch mouse events to children
// - Drawing children
type BaseLayout struct {
	Widget
	LayoutRenderable
}

// NewBaseLayout returns a BaseLayout that embed the given widget.
func NewBaseLayout(
	widget Widget,
	renderable LayoutRenderable,
) *BaseLayout {
	bl := &BaseLayout{
		Widget:           widget,
		LayoutRenderable: renderable,
	}

	// Event handler that forwards mouse position to children.
	dispatchClickEvent := func(event mouse.ClickEvent) {
		for _, childLayout := range bl.LayoutRenderable.ChildrenRects() {
			if childLayout.Bounds.Contains(event.MousePress.RelPosition) {
				event.MousePress.RelPosition = event.MousePress.RelPosition.Sub(childLayout.Bounds.TopLeft())
				childLayout.Node.Unwrap().(events.Target).DispatchEvent(event)
			}
		}
	}
	dispatchMouseEvent := func(event mouse.Event) {
		for _, childLayout := range bl.LayoutRenderable.ChildrenRects() {
			if childLayout.Bounds.Contains(event.RelPosition) {
				event.RelPosition = event.RelPosition.Sub(childLayout.Bounds.TopLeft())
				childLayout.Node.Unwrap().(events.Target).DispatchEvent(event)
			}
		}
	}

	// Dispatch mouse event to child.
	bl.Widget.AddEventListener(mouse.PressListener(dispatchMouseEvent))
	bl.Widget.AddEventListener(mouse.UpListener(dispatchMouseEvent))
	bl.Widget.AddEventListener(mouse.ClickListener(dispatchClickEvent))

	return bl
}

// Renderable implements render.RenderableAccessor.
func (bl *BaseLayout) Renderable() render.Renderable {
	return bl.LayoutRenderable
}

// Style implements styles.Styled.
func (bl *BaseLayout) Style() styles.Style {
	if styled, isStyled := bl.LayoutRenderable.(styles.Styled); isStyled {
		return styled.Style()
	}
	return bl.Widget.Style()
}

// ChildLayout define position and size of children Node.
type ChildLayout struct {
	// Node associated to rectangle.
	Node *tree.Node
	// Bounds relative to layout origin.
	Bounds geometry.Rectangle
}

// LayoutChildren define any type that can layout its children.
type LayoutChildren interface {
	// Determine position and size of each children and return it along its own size.
	// ChildLayout slice argument can be used for appending element instead of allocating
	// a slice on each layout.
	LayoutChildren(_ layout.Constraint, _ []ChildLayout) ([]ChildLayout, geometry.Size)
}

// LayoutChildrenFunc is function type that implements LayoutChildren interface.
type LayoutChildrenFunc func(_ layout.Constraint, _ []ChildLayout) ([]ChildLayout, geometry.Size)

// LayoutChildren implements LayoutChildren.
func (lcf LayoutChildrenFunc) LayoutChildren(co layout.Constraint, r []ChildLayout) ([]ChildLayout, geometry.Size) {
	return lcf(co, r)
}

// LayoutRenderable define render.Renderable for layouts Widget.
type LayoutRenderable interface {
	render.Renderable
	// Returns children bounding rectangle relative to layout origin
	ChildrenRects() []ChildLayout
}

var _ LayoutRenderable = &BaseLayoutRenderable{}

// BaseLayoutRenderable is a wrapper of render.VoidRenderable that implements LayoutRenderable.
type BaseLayoutRenderable struct {
	render.VoidRenderable

	layout LayoutChildren
	// Children positions relative to layout origin.
	childrenLayout []ChildLayout
}

// NewBaseLayoutRenderable returns a new LayoutRenderable for the given NodeAccessor.
func NewBaseLayoutRenderable(nodeAccessor tree.NodeAccessor, layoutChildren LayoutChildren) *BaseLayoutRenderable {
	return &BaseLayoutRenderable{
		VoidRenderable: render.NewVoidRenderable(nodeAccessor),
		layout:         layoutChildren,
		childrenLayout: []ChildLayout{},
	}
}

// ChildrenRects implements LayoutRenderable.
func (lr *BaseLayoutRenderable) ChildrenRects() []ChildLayout {
	return lr.childrenLayout
}

// Layout implements layout.Layout.
func (lr *BaseLayoutRenderable) Layout(co layout.Constraint) geometry.Size {
	var size geometry.Size
	lr.childrenLayout, size = lr.layout.LayoutChildren(co, lr.childrenLayout[:0])
	return size
}

// Draw implements layout.Layout.
// BaseLayoutRenderable draw
func (lr *BaseLayoutRenderable) Draw(surface draw.Surface) {
	lr.VoidRenderable.Draw(surface)

	for _, childLayout := range lr.childrenLayout {
		childDrawer := childLayout.Node.Unwrap().(render.RenderableAccessor).Renderable()
		subsurface := draw.NewSubSurface(surface, childLayout.Bounds)

		childDrawer.Draw(subsurface)
	}
}

// LayoutRenderableCache define a LayoutRenderable with cache.
type LayoutRenderableCache struct {
	render.Cache[*BaseLayoutRenderable]
}

// NewLayoutRenderableCache returns a new LayoutRenderable.
func NewLayoutRenderableCache(nodeAccessor tree.NodeAccessor, layout LayoutChildren) LayoutRenderableCache {
	cache := render.NewCache(NewBaseLayoutRenderable(
		nodeAccessor,
		layout,
	))

	return LayoutRenderableCache{
		Cache: cache,
	}
}

// ChildrenRects implements LayoutRenderable.
func (lrc LayoutRenderableCache) ChildrenRects() []ChildLayout {
	return lrc.Cache.Unwrap().ChildrenRects()
}

// StyledLayoutRenderable define a LayoutRenderable wrapper that adds styling
// to it.
type StyledLayoutRenderable struct {
	LayoutRenderable
	styled styles.Styled
}

// NewStyledLayoutRenderable returns a new StyledLayoutRenderable that wraps
// the given LayoutRenderable.
func NewStyledLayoutRenderable(styled styles.Styled, layoutRenderable LayoutRenderable) StyledLayoutRenderable {
	return StyledLayoutRenderable{
		LayoutRenderable: layoutRenderable,
		styled:           styled,
	}
}

// Style implements styled.Styled.
func (slr StyledLayoutRenderable) Style() styles.Style {
	return slr.styled.Style()
}

// Layout implements layout.Layout.
func (slr StyledLayoutRenderable) Layout(co layout.Constraint) geometry.Size {
	style := slr.styled.Style().Compute()
	origin := styles.LayoutContentBoxOrigin(style)

	size := styles.Layout(
		slr.styled.Style().Compute(),
		co,
		layout.LayoutFunc(func(co layout.Constraint) geometry.Size {
			return slr.LayoutRenderable.Layout(co)
		}),
	)

	// Translate widgets inside content box.
	if origin.X() != 0 || origin.Y() != 0 {
		childrenLayout := slr.LayoutRenderable.ChildrenRects()
		for i := range childrenLayout {
			childLayout := &childrenLayout[i]
			childLayout.Bounds = childLayout.Bounds.MoveBy(origin)
		}
	}

	return size
}

// Draw implements draw.Drawer
func (slr StyledLayoutRenderable) Draw(surface draw.Surface) {
	_ = styles.Draw(slr.styled.Style().Compute(), surface)
	slr.LayoutRenderable.Draw(surface)
}
