package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/styles"
)

type privateStyle = styles.Style

type StyledLayout struct {
	*BaseLayout
	privateStyle
}

func NewStyledLayout(
	widget Widget,
	style styles.Style,
	algo LayoutAlgo,
) StyledLayout {
	sl := StyledLayout{
		privateStyle: style,
	}

	sl.BaseLayout = NewBaseLayout(widget, algo)

	return sl
}

// Layout implements layout.Layout.
func (sl StyledLayout) Layout(co layout.Constraint) (size geometry.Size) {
	sl.latestLayoutConstraint = co

	style := sl.Compute()
	origin := styles.LayoutContentBoxOrigin(style)

	size = styles.Layout(style, co, layout.LayoutFunc(func(co layout.Constraint) geometry.Size {
		var size geometry.Size
		sl.childrenRects, size = sl.layoutAlgo(co, sl.childrenRects[:0])
		return size
	}))

	// Translate widgets inside content box.
	if origin.X() != 0 || origin.Y() != 0 {
		for i, rect := range sl.childrenRects {
			sl.childrenRects[i] = rect.MoveBy(origin)
		}
	}

	return size
}

// Draw implements draw.Drawer.
func (sl StyledLayout) Draw(surface draw.Surface) {
	_ = styles.Draw(sl.Compute(), surface)

	child := sl.Node().FirstChild()
	for _, boundingRect := range sl.childrenRects {
		childDrawer := child.Unwrap().(draw.Drawer)
		subsurface := draw.NewSubSurface(surface, boundingRect)

		childDrawer.Draw(subsurface)

		child = child.Next()
	}
}

// Style implements styles.Styled
func (sl StyledLayout) Style() styles.Style {
	return sl.privateStyle
}

// SetStyle sets internal layout style to the given one.
func (sl *StyledLayout) SetStyle(s styles.Style) {
	sl.privateStyle = s
}
