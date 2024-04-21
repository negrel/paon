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
		childrenLayout := bl.LayoutRenderable.ChildrenLayout()
		for i := 0; i < childrenLayout.Size(); i++ {
			childLayout := childrenLayout.Get(i)

			if childLayout.Bounds.Contains(event.MousePress.RelPosition) {
				event.MousePress.RelPosition = event.MousePress.RelPosition.Sub(childLayout.Bounds.TopLeft())
				childLayout.Node.Unwrap().(events.Target).DispatchEvent(event)
			}
		}
	}
	dispatchMouseEvent := func(event mouse.Event) {
		childrenLayout := bl.LayoutRenderable.ChildrenLayout()
		for i := 0; i < childrenLayout.Size(); i++ {
			childLayout := childrenLayout.Get(i)

			if childLayout.Bounds.Contains(event.RelPosition) {
				event.RelPosition = event.RelPosition.Sub(childLayout.Bounds.TopLeft())
				childLayout.Node.Unwrap().(events.Target).DispatchEvent(event)
			}
		}
	}
	dispatchScrollEvent := func(event mouse.ScrollEvent) {
		childrenLayout := bl.LayoutRenderable.ChildrenLayout()
		for i := 0; i < childrenLayout.Size(); i++ {
			childLayout := childrenLayout.Get(i)

			if childLayout.Bounds.Contains(event.RelPosition) {
				event.RelPosition = event.RelPosition.Sub(childLayout.Bounds.TopLeft())
				childLayout.Node.Unwrap().(events.Target).DispatchEvent(event)
			}
		}
	}

	// Dispatch mouse event to child.
	bl.Widget.AddEventListener(mouse.EventListener(dispatchMouseEvent))
	bl.Widget.AddEventListener(mouse.PressListener(dispatchMouseEvent))
	bl.Widget.AddEventListener(mouse.UpListener(dispatchMouseEvent))
	bl.Widget.AddEventListener(mouse.ClickListener(dispatchClickEvent))
	bl.Widget.AddEventListener(mouse.ScrollListener(dispatchScrollEvent))

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

// LayoutChildren define any type that can layout its children.
type LayoutChildren interface {
	// LayoutChildren layout each children and append them in the given
	// ChildrenLayout.
	LayoutChildren(_ layout.Constraint, _ *ChildrenLayout) geometry.Size
}

// LayoutChildrenFunc is function type that implements LayoutChildren interface.
type LayoutChildrenFunc func(_ layout.Constraint, _ *ChildrenLayout) geometry.Size

// LayoutChildren implements LayoutChildren.
func (lcf LayoutChildrenFunc) LayoutChildren(co layout.Constraint, r *ChildrenLayout) geometry.Size {
	return lcf(co, r)
}

// LayoutRenderable define render.Renderable for layouts Widget.
type LayoutRenderable interface {
	render.Renderable
	// Returns children bounding rectangle relative to layout origin
	ChildrenLayout() *ChildrenLayout
}

var _ LayoutRenderable = BaseLayoutRenderable{}

// BaseLayoutRenderable is a wrapper of render.VoidRenderable that implements LayoutRenderable.
type BaseLayoutRenderable struct {
	*render.VoidRenderable

	layout         LayoutChildren
	childrenLayout *ChildrenLayout
}

// NewBaseLayoutRenderable returns a new LayoutRenderable for the given NodeAccessor.
func NewBaseLayoutRenderable(nodeAccessor tree.NodeAccessor, layoutChildren LayoutChildren) *BaseLayoutRenderable {
	vr := render.NewVoidRenderable(nodeAccessor)

	return &BaseLayoutRenderable{
		VoidRenderable: &vr,
		layout:         layoutChildren,
		childrenLayout: &ChildrenLayout{},
	}
}

// ChildrenRects implements LayoutRenderable.
func (lr BaseLayoutRenderable) ChildrenLayout() *ChildrenLayout {
	return lr.childrenLayout
}

// Layout implements layout.Layout.
func (lr BaseLayoutRenderable) Layout(co layout.Constraint) geometry.Size {
	lr.childrenLayout.reset()
	return lr.layout.LayoutChildren(co, lr.childrenLayout)
}

// Draw implements layout.Layout.
// BaseLayoutRenderable draw
func (lr BaseLayoutRenderable) Draw(surface draw.Surface) {
	lr.VoidRenderable.Draw(surface)

	for i := 0; i < lr.childrenLayout.Size(); i++ {
		childLayout := lr.childrenLayout.Get(i)

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

// ChildrenLayout implements LayoutRenderable.
func (lrc LayoutRenderableCache) ChildrenLayout() *ChildrenLayout {
	return lrc.Cache.Unwrap().ChildrenLayout()
}

// StyledLayoutRenderable define a LayoutRenderable wrapper that adds styling
// to it.
type StyledLayoutRenderable struct {
	styles.Renderable[LayoutRenderable]
}

// NewStyledLayoutRenderable returns a new StyledLayoutRenderable that wraps
// the given LayoutRenderable.
func NewStyledLayoutRenderable(styled styles.Styled, layoutRenderable LayoutRenderable) StyledLayoutRenderable {
	return StyledLayoutRenderable{
		Renderable: styles.Renderable[LayoutRenderable]{
			Renderable: layoutRenderable,
			Styled:     styled,
		},
	}
}

// Style implements styled.Styled.
func (slr StyledLayoutRenderable) Style() styles.Style {
	return slr.Renderable.Styled.Style()
}

// ChildrenLayout implements LayoutRenderable.
func (slr StyledLayoutRenderable) ChildrenLayout() *ChildrenLayout {
	return slr.Renderable.Renderable.ChildrenLayout()
}

// Layout implements layout.Layout.
func (slr StyledLayoutRenderable) Layout(co layout.Constraint) geometry.Size {
	style := slr.Styled.Style()
	if style == nil {
		return slr.Renderable.Layout(co)
	}

	size := slr.Renderable.Layout(co)

	computedStyle := style.Compute()
	origin := styles.LayoutContentBoxOrigin(computedStyle)
	slr.ChildrenLayout().origin = origin

	return size
}

// Draw implements draw.Drawer.
func (slr StyledLayoutRenderable) Draw(surface draw.Surface) {
	// Alter origin so renderable.Draw uses surface origin.
	childrenLayout := slr.ChildrenLayout()
	origin := childrenLayout.origin
	childrenLayout.origin = geometry.Vec2D{}

	slr.Renderable.Draw(surface)

	// Restore origin.
	childrenLayout.origin = origin
}
