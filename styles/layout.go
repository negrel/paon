package styles

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets/border"
)

func LayoutFunc(s ComputedStyle, next layout.LayoutFunc) layout.Layout {
	return layout.LayoutFunc(func(co layout.Constraint) geometry.Size {
		return Layout(s, co, next)
	})
}

func Layout(s ComputedStyle, co layout.Constraint, next layout.Layout) geometry.Size {
	diffWidth := 0
	diffHeight := 0

	if (s.MarginStyle != MarginPaddingStyle{}) {
		diffWidth += s.MarginStyle.Left + s.MarginStyle.Right
		diffHeight += s.MarginStyle.Top + s.MarginStyle.Bottom
	}

	if (s.BorderStyle.Border != border.Border{}) {
		if s.BorderStyle.Left != "" {
			diffWidth++
		}
		if s.BorderStyle.Right != "" {
			diffWidth++
		}

		if s.BorderStyle.Top != "" {
			diffHeight++
		}
		if s.BorderStyle.Bottom != "" {
			diffHeight++
		}
	}

	if (s.PaddingStyle != MarginPaddingStyle{}) {
		diffWidth += s.PaddingStyle.Left + s.PaddingStyle.Right
		diffHeight += s.PaddingStyle.Top + s.PaddingStyle.Bottom
	}

	co.MaxSize = geometry.NewSize(
		co.MaxSize.Width()-diffWidth,
		co.MaxSize.Height()-diffHeight,
	)

	size := next.Layout(co)

	return geometry.NewSize(
		size.Width()+diffWidth,
		size.Height()+diffHeight,
	)
}

func LayoutContentBoxOrigin(s ComputedStyle) geometry.Vec2D {
	originX := 0
	originY := 0

	if (s.MarginStyle != MarginPaddingStyle{}) {
		originX += s.MarginStyle.Left
		originY += s.MarginStyle.Top
	}

	if (s.BorderStyle.Border != border.Border{}) {
		if s.BorderStyle.Left != "" {
			originX++
		}
		if s.BorderStyle.Top != "" {
			originY++
		}
	}

	if (s.PaddingStyle != MarginPaddingStyle{}) {
		originX += s.PaddingStyle.Left
		originY += s.PaddingStyle.Top
	}

	return geometry.NewVec2D(originX, originY)
}
