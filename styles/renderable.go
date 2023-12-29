package styles

import (
	"github.com/negrel/paon/colors"
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
		diffWidth += computedStyle.BordersStyle.Left.Size
	}
	if computedStyle.BordersStyle.Right.Size != 0 {
		diffWidth += computedStyle.BordersStyle.Right.Size
	}
	if computedStyle.BordersStyle.Top.Size != 0 {
		diffHeight += computedStyle.BordersStyle.Top.Size
	}
	if computedStyle.BordersStyle.Bottom.Size != 0 {
		diffHeight += computedStyle.BordersStyle.Bottom.Size
	}

	if (computedStyle.PaddingStyle != MarginPaddingStyle{}) {
		diffWidth += computedStyle.PaddingStyle.Left + computedStyle.PaddingStyle.Right
		diffHeight += computedStyle.PaddingStyle.Top + computedStyle.PaddingStyle.Bottom
	}

	co.MaxSize = geometry.Size{
		Width:  co.MaxSize.Width - diffWidth,
		Height: co.MaxSize.Height - diffHeight,
	}

	size := r.Renderable.Layout(co)

	return geometry.Size{
		Width:  size.Width + diffWidth,
		Height: size.Height + diffHeight,
	}
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
	borderBox := geometry.Rectangle{
		Origin:   geometry.Vec2D{},
		RectSize: surfaceSize,
	}

	if (computedStyle.MarginStyle != MarginPaddingStyle{}) {
		borderBox = borderBox.GrowLeft(-computedStyle.MarginStyle.Left)
		borderBox = borderBox.GrowTop(-computedStyle.MarginStyle.Top)
		borderBox = borderBox.GrowRight(-computedStyle.MarginStyle.Right)
		borderBox = borderBox.GrowBottom(-computedStyle.MarginStyle.Bottom)
	}

	subsurface := draw.NewSubSurface(surface, borderBox)
	fillSurfaceWithCellStyle(draw.CellStyle{
		Background: colors.ColorGreen,
	}, subsurface)

	paddingBox := borderBox

	if (computedStyle.BordersStyle != BordersStyle{}) {
		if leftBorderStyle := computedStyle.BordersStyle.Left; leftBorderStyle.Size != 0 {
			subsurface = draw.NewSubSurface(
				surface,
				borderBox.ShrinkRight(borderBox.Size().Width-leftBorderStyle.Size),
			)
			BorderDrawers[leftBorderStyle.Style].Left(computedStyle.BordersStyle, subsurface)

			paddingBox = paddingBox.ShrinkLeft(leftBorderStyle.Size)
		}
		if topBorderStyle := computedStyle.BordersStyle.Top; topBorderStyle.Size != 0 {
			subsurface = draw.NewSubSurface(
				surface,
				borderBox.ShrinkBottom(borderBox.Size().Height-topBorderStyle.Size),
			)
			BorderDrawers[topBorderStyle.Style].Top(computedStyle.BordersStyle, subsurface)

			paddingBox = paddingBox.ShrinkTop(topBorderStyle.Size)
		}
		if rightBorderStyle := computedStyle.BordersStyle.Right; rightBorderStyle.Size != 0 {
			subsurface = draw.NewSubSurface(
				surface,
				borderBox.ShrinkLeft(borderBox.Size().Width-rightBorderStyle.Size),
			)
			BorderDrawers[rightBorderStyle.Style].Right(computedStyle.BordersStyle, subsurface)

			paddingBox = paddingBox.ShrinkRight(rightBorderStyle.Size)
		}
		if bottomBorderStyle := computedStyle.BordersStyle.Bottom; bottomBorderStyle.Size != 0 {
			subsurface = draw.NewSubSurface(
				surface,
				borderBox.ShrinkTop(borderBox.Size().Height-bottomBorderStyle.Size),
			)
			BorderDrawers[bottomBorderStyle.Style].Bottom(computedStyle.BordersStyle, subsurface)

			paddingBox = paddingBox.GrowBottom(-bottomBorderStyle.Size)
		}
	}

	subsurface = draw.NewSubSurface(surface, paddingBox)
	fillSurfaceWithCellStyle(computedStyle.CellStyle, subsurface)

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
		originX += s.BordersStyle.Left.Size
	}
	if s.BordersStyle.Top.Size != 0 {
		originY += s.BordersStyle.Top.Size
	}

	if (s.PaddingStyle != MarginPaddingStyle{}) {
		originX += s.PaddingStyle.Left
		originY += s.PaddingStyle.Top
	}

	return geometry.Vec2D{X: originX, Y: originY}
}
