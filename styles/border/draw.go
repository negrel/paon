package border

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
)

// BorderOld contains a series of values which comprise the various parts of a
// border.
type BorderOld struct {
	Top          string
	Bottom       string
	Left         string
	Right        string
	TopLeft      rune
	TopRight     rune
	BottomLeft   rune
	BottomRight  rune
	MiddleLeft   rune
	MiddleRight  rune
	Middle       rune
	MiddleTop    rune
	MiddleBottom rune
}

var (
	NoBorder = BorderOld{}

	NormalBorder = BorderOld{
		Top:          "─",
		Bottom:       "─",
		Left:         "│",
		Right:        "│",
		TopLeft:      '┌',
		TopRight:     '┐',
		BottomLeft:   '└',
		BottomRight:  '┘',
		MiddleLeft:   '├',
		MiddleRight:  '┤',
		Middle:       '┼',
		MiddleTop:    '┬',
		MiddleBottom: '┴',
	}

	RoundedBorder = BorderOld{
		Top:          "─",
		Bottom:       "─",
		Left:         "│",
		Right:        "│",
		TopLeft:      '╭',
		TopRight:     '╮',
		BottomLeft:   '╰',
		BottomRight:  '╯',
		MiddleLeft:   '├',
		MiddleRight:  '┤',
		Middle:       '┼',
		MiddleTop:    '┬',
		MiddleBottom: '┴',
	}

	BlockBorder = BorderOld{
		Top:         "█",
		Bottom:      "█",
		Left:        "█",
		Right:       "█",
		TopLeft:     '█',
		TopRight:    '█',
		BottomLeft:  '█',
		BottomRight: '█',
	}

	OuterHalfBlockBorder = BorderOld{
		Top:         "▀",
		Bottom:      "▄",
		Left:        "▌",
		Right:       "▐",
		TopLeft:     '▛',
		TopRight:    '▜',
		BottomLeft:  '▙',
		BottomRight: '▟',
	}

	InnerHalfBlockBorder = BorderOld{
		Top:         "▄",
		Bottom:      "▀",
		Left:        "▐",
		Right:       "▌",
		TopLeft:     '▗',
		TopRight:    '▖',
		BottomLeft:  '▝',
		BottomRight: '▘',
	}

	ThickBorder = BorderOld{
		Top:          "━",
		Bottom:       "━",
		Left:         "┃",
		Right:        "┃",
		TopLeft:      '┏',
		TopRight:     '┓',
		BottomLeft:   '┗',
		BottomRight:  '┛',
		MiddleLeft:   '┣',
		MiddleRight:  '┫',
		Middle:       '╋',
		MiddleTop:    '┳',
		MiddleBottom: '┻',
	}

	DoubleBorder = BorderOld{
		Top:          "═",
		Bottom:       "═",
		Left:         "║",
		Right:        "║",
		TopLeft:      '╔',
		TopRight:     '╗',
		BottomLeft:   '╚',
		BottomRight:  '╝',
		MiddleLeft:   '╠',
		MiddleRight:  '╣',
		Middle:       '╬',
		MiddleTop:    '╦',
		MiddleBottom: '╩',
	}

	HiddenBorder = BorderOld{
		Top:          " ",
		Bottom:       " ",
		Left:         " ",
		Right:        " ",
		TopLeft:      ' ',
		TopRight:     ' ',
		BottomLeft:   ' ',
		BottomRight:  ' ',
		MiddleLeft:   ' ',
		MiddleRight:  ' ',
		Middle:       ' ',
		MiddleTop:    ' ',
		MiddleBottom: ' ',
	}
)

// DrawBorder draw borders onto the given surface.
func DrawBorder(surface draw.Surface, border BorderOld, style draw.CellStyle) {
	borderTop := []rune(border.Top)
	borderBottom := []rune(border.Bottom)
	borderLeft := []rune(border.Left)
	borderRight := []rune(border.Right)

	bottom := surface.Size().Height() - 1
	right := surface.Size().Width() - 1

	// Top and bottom borders.
	for i := 0; i < surface.Size().Width(); i++ {
		if border.Top != "" {
			r := borderTop[i%len(borderTop)]
			surface.Set(geometry.NewVec2D(i, 0), draw.Cell{
				Style:   style,
				Content: r,
			})
		}

		if border.Bottom != "" {
			r := borderBottom[i%len(borderBottom)]
			surface.Set(geometry.NewVec2D(i, bottom), draw.Cell{
				Style:   style,
				Content: r,
			})
		}
	}

	// Left and right borders.
	for i := 0; i < surface.Size().Height(); i++ {
		if border.Left != "" {
			r := borderLeft[i%len(borderLeft)]
			surface.Set(geometry.NewVec2D(0, i), draw.Cell{
				Style:   style,
				Content: r,
			})
		}

		if border.Right != "" {
			r := borderRight[i%len(borderRight)]
			surface.Set(geometry.NewVec2D(right, i), draw.Cell{
				Style:   style,
				Content: r,
			})
		}
	}

	// Corners
	surface.Set(geometry.NewVec2D(0, 0), draw.Cell{
		Style:   style,
		Content: border.TopLeft,
	})
	surface.Set(geometry.NewVec2D(right, 0), draw.Cell{
		Style:   style,
		Content: border.TopRight,
	})
	surface.Set(geometry.NewVec2D(right, bottom), draw.Cell{
		Style:   style,
		Content: border.BottomRight,
	})
	surface.Set(geometry.NewVec2D(0, bottom), draw.Cell{
		Style:   style,
		Content: border.BottomLeft,
	})
}
