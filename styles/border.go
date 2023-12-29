package styles

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
)

// BorderStyle enumerates border styles.
type BorderStyle int8

const (
	BorderHidden BorderStyle = iota
)

// Border define style properties of a single border side.
type BorderSide struct {
	Size  int
	Style BorderStyle
	draw.CellStyle
}

type BorderDrawer interface {
	Left(BordersStyle, draw.Surface)
	Top(BordersStyle, draw.Surface)
	Right(BordersStyle, draw.Surface)
	Bottom(BordersStyle, draw.Surface)
}

var (
	BorderDrawers = []BorderDrawer{
		BorderHidden: borderDrawer{
			left:   func(bs BordersStyle, surface draw.Surface) { fillSurfaceWithCellStyle(bs.Left.CellStyle, surface) },
			top:    func(bs BordersStyle, surface draw.Surface) { fillSurfaceWithCellStyle(bs.Top.CellStyle, surface) },
			right:  func(bs BordersStyle, surface draw.Surface) { fillSurfaceWithCellStyle(bs.Right.CellStyle, surface) },
			bottom: func(bs BordersStyle, surface draw.Surface) { fillSurfaceWithCellStyle(bs.Bottom.CellStyle, surface) },
		},
	}
)

type borderDrawer struct {
	left   func(BordersStyle, draw.Surface)
	top    func(BordersStyle, draw.Surface)
	right  func(BordersStyle, draw.Surface)
	bottom func(BordersStyle, draw.Surface)
}

func (bd borderDrawer) Left(style BordersStyle, surface draw.Surface) {
	bd.left(style, surface)
}

func (bd borderDrawer) Top(style BordersStyle, surface draw.Surface) {
	bd.top(style, surface)
}

func (bd borderDrawer) Right(style BordersStyle, surface draw.Surface) {
	bd.right(style, surface)
}

func (bd borderDrawer) Bottom(style BordersStyle, surface draw.Surface) {
	bd.bottom(style, surface)
}

func fillSurfaceWithCellStyle(cellstyle draw.CellStyle, surface draw.Surface) {
	for i := 0; i < surface.Size().Width; i++ {
		for j := 0; j < surface.Size().Height; j++ {
			surface.Set(geometry.Vec2D{X: i, Y: j}, draw.Cell{
				Style:   cellstyle,
				Content: 0,
			})
		}
	}
}
