package border

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
)

func New(border *Border, style *draw.CellStyle) (layout.Pipe, draw.Pipe) {
	return func(co layout.Constraint, next layout.Layout) geometry.Size {
			co.MaxSize = geometry.NewSize(co.MaxSize.Width()-2, co.MaxSize.Height()-2)
			size := next.Layout(co)
			return geometry.NewSize(size.Width()+2, size.Height()+2)
		}, func(surface draw.Surface, next draw.Drawer) {
			if style == nil {
				DrawBorder(surface, *border, draw.CellStyle{})
			} else {
				DrawBorder(surface, *border, *style)
			}

			surfaceSize := surface.Size()
			surface = draw.NewSubSurface(surface, geometry.Rect(1, 1, surfaceSize.Width()-1, surfaceSize.Height()-1))

			next.Draw(surface)
		}
}

func Normal(style *draw.CellStyle) (layout.Pipe, draw.Pipe) {
	return New(&NormalBorder, style)
}

func Rounded(style *draw.CellStyle) (layout.Pipe, draw.Pipe) {
	return New(&RoundedBorder, style)
}

func Block(style *draw.CellStyle) (layout.Pipe, draw.Pipe) {
	return New(&BlockBorder, style)
}

func OuterHalfBlock(style *draw.CellStyle) (layout.Pipe, draw.Pipe) {
	return New(&OuterHalfBlockBorder, style)
}

func InnerHalfBlock(style *draw.CellStyle) (layout.Pipe, draw.Pipe) {
	return New(&InnerHalfBlockBorder, style)
}

func Thick(style *draw.CellStyle) (layout.Pipe, draw.Pipe) {
	return New(&ThickBorder, style)
}

func Double(style *draw.CellStyle) (layout.Pipe, draw.Pipe) {
	return New(&DoubleBorder, style)
}

func Hidden(style *draw.CellStyle) (layout.Pipe, draw.Pipe) {
	return New(&HiddenBorder, style)
}
