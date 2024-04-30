package main

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/minmax"
	"github.com/negrel/paon/widgets"
)

// Scroll is our scrolling component that display file content.
type Scroll struct {
	widgets.BaseWidget

	lines   []string
	topLine int
}

// NewScroll creates a new scroll widget.
func NewScroll(lines []string) *Scroll {
	s := &Scroll{lines: lines}
	s.BaseWidget = widgets.NewBaseWidget(s)
	return s
}

func (s *Scroll) ScrollDown() {
	s.topLine = minmax.Min(len(s.lines), s.topLine+2)
	// Trigger rendering.
	s.NeedRender()
}

func (s *Scroll) ScrollUp() {
	s.topLine = minmax.Max(0, s.topLine-2)
	// Trigger rendering.
	s.NeedRender()
}

// Layout implements widgets.Widget.
func (s *Scroll) Layout(co layout.Constraint, ctx widgets.LayoutContext) layout.SizeHint {
	// Use all available space.
	return layout.SizeHint{
		MinSize:          co.MaxSize,
		Size:             co.MaxSize,
		VerticalPolicy:   layout.ExpandingSizePolicy,
		HorizontalPolicy: layout.ExpandingSizePolicy,
	}
}

// Draw implements widgets.Widget.
func (s *Scroll) Draw(srf draw.Surface) {
	// For each lines...
	for i := s.topLine; i < minmax.Min(srf.Size().Height, len(s.lines)); i++ {
		line := s.lines[i]

		// For each character.
		for x, char := range line {
			// Set terminal cell.
			srf.Set(geometry.Vec2D{X: x, Y: i - s.topLine}, draw.Cell{
				Style:   draw.CellStyle{},
				Content: char,
			})
		}
	}
}
