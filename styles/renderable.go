package styles

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/render"
)

var _ render.Renderable = Renderable[*render.VoidRenderable]{}

type Renderable[R render.Renderable] struct {
	Renderable R
	Styled     Styled
}

// Style implements Styled.
func (r Renderable[R]) Style() Style {
	return r.Styled.Style()
}

// IsDirty implements render.Renderable.
func (r Renderable[R]) IsDirty() bool {
	return r.Renderable.IsDirty()
}

// MarkDirty implements render.Renderable.
func (r Renderable[R]) MarkDirty() {
	r.Renderable.MarkDirty()
}

// Layout implements layout.Layout.
func (r Renderable[R]) Layout(co layout.Constraint) geometry.Size {
	style := r.Styled.Style()
	if style == nil {
		return r.Renderable.Layout(co)
	}
	computedStyle := style.Compute()

	diffWidth := 0
	diffHeight := 0

	if (computedStyle.MarginStyle != MarginPaddingStyle{}) {
		diffWidth += computedStyle.MarginStyle.Left + computedStyle.MarginStyle.Right
		diffHeight += computedStyle.MarginStyle.Top + computedStyle.MarginStyle.Bottom
	}

	if computedStyle.BordersStyle.Left.Size != 0 {
		diffWidth++
	}
	if computedStyle.BordersStyle.Right.Size != 0 {
		diffWidth++
	}
	if computedStyle.BordersStyle.Top.Size != 0 {
		diffHeight++
	}
	if computedStyle.BordersStyle.Bottom.Size != 0 {
		diffHeight++
	}

	if (computedStyle.PaddingStyle != MarginPaddingStyle{}) {
		diffWidth += computedStyle.PaddingStyle.Left + computedStyle.PaddingStyle.Right
		diffHeight += computedStyle.PaddingStyle.Top + computedStyle.PaddingStyle.Bottom
	}

	co.MaxSize = geometry.NewSize(
		co.MaxSize.Width()-diffWidth,
		co.MaxSize.Height()-diffHeight,
	)

	size := r.Renderable.Layout(co)

	return geometry.NewSize(
		size.Width()+diffWidth,
		size.Height()+diffHeight,
	)
}

// Draw implements draw.Drawer.
func (r Renderable[R]) Draw(surface draw.Surface) {
	style := r.Styled.Style()
	if style == nil {
		r.Renderable.Draw(surface)
		return
	}
	computedStyle := style.Compute()

	surfaceSize := surface.Size()
	borderBox := geometry.Rect(0, 0, surfaceSize.Width(), surfaceSize.Height())
	if (computedStyle.MarginStyle != MarginPaddingStyle{}) {
		borderBox = borderBox.GrowLeft(-computedStyle.MarginStyle.Left)
		borderBox = borderBox.GrowTop(-computedStyle.MarginStyle.Top)
		borderBox = borderBox.GrowRight(-computedStyle.MarginStyle.Right)
		borderBox = borderBox.GrowBottom(-computedStyle.MarginStyle.Bottom)
	}

	subsurface := draw.NewSubSurface(surface, borderBox)
	if (computedStyle.CellStyle != draw.CellStyle{}) {
		for w := 0; w < subsurface.Size().Width(); w++ {
			for h := 0; h < subsurface.Size().Height(); h++ {
				subsurface.Set(geometry.NewVec2D(w, h), draw.Cell{
					Style: computedStyle.CellStyle,
				})
			}
		}
	}

	paddingBox := borderBox

	if (computedStyle.BordersStyle != BordersStyle{}) {
		if leftBorderStyle := computedStyle.BordersStyle.Left; leftBorderStyle.Size != 0 {
			subsurface = draw.NewSubSurface(
				surface,
				geometry.Rect(
					borderBox.TopLeft().X(),
					borderBox.TopLeft().Y(),
					borderBox.BottomLeft().X()+leftBorderStyle.Size,
					borderBox.BottomLeft().Y(),
				),
			)
			BorderDrawers[leftBorderStyle.Style].Left(computedStyle.BordersStyle, subsurface)

			paddingBox = paddingBox.GrowLeft(-1)
		}
		if topBorderStyle := computedStyle.BordersStyle.Top; topBorderStyle.Size != 0 {
			subsurface = draw.NewSubSurface(
				surface,
				geometry.Rect(
					borderBox.TopLeft().X(),
					borderBox.TopLeft().Y(),
					borderBox.TopRight().Y(),
					borderBox.TopRight().Y()+topBorderStyle.Size,
				),
			)
			BorderDrawers[topBorderStyle.Style].Top(computedStyle.BordersStyle, subsurface)

			paddingBox = paddingBox.GrowTop(-1)
		}
		if rightBorderStyle := computedStyle.BordersStyle.Right; rightBorderStyle.Size != 0 {
			subsurface = draw.NewSubSurface(
				surface,
				geometry.Rect(
					borderBox.TopRight().X()-rightBorderStyle.Size,
					borderBox.TopRight().Y(),
					borderBox.BottomRight().X(),
					borderBox.BottomRight().Y(),
				),
			)
			BorderDrawers[rightBorderStyle.Style].Right(computedStyle.BordersStyle, subsurface)

			paddingBox = paddingBox.GrowRight(-1)
		}
		if bottomBorderStyle := computedStyle.BordersStyle.Bottom; bottomBorderStyle.Size != 0 {
			subsurface = draw.NewSubSurface(
				surface,
				geometry.Rect(
					borderBox.BottomLeft().X(),
					borderBox.BottomLeft().Y()-bottomBorderStyle.Size,
					borderBox.BottomRight().X(),
					borderBox.BottomRight().Y(),
				),
			)
			BorderDrawers[bottomBorderStyle.Style].Bottom(computedStyle.BordersStyle, subsurface)

			paddingBox = paddingBox.GrowBottom(-1)
		}
	}

	subsurface = draw.NewSubSurface(surface, paddingBox)

	fillSurfaceWithCellStyle(r.Styled.Style().Compute().CellStyle, subsurface)

	contentBox := paddingBox
	if (computedStyle.PaddingStyle != MarginPaddingStyle{}) {
		contentBox = contentBox.GrowLeft(-computedStyle.PaddingStyle.Left)
		contentBox = contentBox.GrowTop(-computedStyle.PaddingStyle.Top)
		contentBox = contentBox.GrowRight(-computedStyle.PaddingStyle.Right)
		contentBox = contentBox.GrowBottom(-computedStyle.PaddingStyle.Bottom)
	}

	subsurface = draw.NewSubSurface(surface, contentBox)

	r.Renderable.Draw(subsurface)
}

func LayoutContentBoxOrigin(s ComputedStyle) geometry.Vec2D {
	originX := 0
	originY := 0

	if (s.MarginStyle != MarginPaddingStyle{}) {
		originX += s.MarginStyle.Left
		originY += s.MarginStyle.Top
	}

	if s.BordersStyle.Left.Size != 0 {
		originX++
	}
	if s.BordersStyle.Top.Size != 0 {
		originY++
	}

	if (s.PaddingStyle != MarginPaddingStyle{}) {
		originX += s.PaddingStyle.Left
		originY += s.PaddingStyle.Top
	}

	return geometry.NewVec2D(originX, originY)
}
