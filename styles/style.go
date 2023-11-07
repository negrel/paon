package styles

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets/border"
)

// Style define common widget styling properties.
type Style struct {
	Margin  *MarginStyle
	Border  *BorderStyle
	Padding *PaddingStyle
	Content *draw.CellStyle
	Extras  any
}

// New creates a new style configured with the given options.
func New(options ...Option) *Style {
	style := &Style{}

	for _, applyOption := range options {
		applyOption(style)
	}

	return style
}

// Style implements Styled.
func (s *Style) Style() *Style {
	return s
}

func Layout(s *Style, co layout.Constraint, next layout.Layout) geometry.Size {
	diffWidth := 0
	diffHeight := 0

	if s.Margin != nil {
		diffWidth += s.Margin.Left + s.Margin.Right
		diffHeight += s.Margin.Top + s.Margin.Bottom
	}

	if s.Border != nil {
		if s.Border.Left != "" {
			diffWidth++
		}
		if s.Border.Right != "" {
			diffWidth++
		}

		if s.Border.Top != "" {
			diffHeight++
		}
		if s.Border.Bottom != "" {
			diffHeight++
		}
	}

	if s.Padding != nil {
		diffWidth += s.Padding.Left + s.Padding.Right
		diffHeight += s.Padding.Top + s.Padding.Bottom
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

func Draw(s *Style, surface draw.Surface, next draw.Drawer) {
	surfaceSize := surface.Size()
	borderBox := geometry.Rect(0, 0, surfaceSize.Width(), surfaceSize.Height())
	if s.Margin != nil {
		borderBox = borderBox.GrowLeft(-s.Margin.Left)
		borderBox = borderBox.GrowTop(-s.Margin.Top)
		borderBox = borderBox.GrowRight(-s.Margin.Right)
		borderBox = borderBox.GrowBottom(-s.Margin.Bottom)
	}

	subsurface := draw.NewSubSurface(surface, borderBox)

	paddingBox := borderBox
	if s.Border != nil {
		border.DrawBorder(subsurface, s.Border.Border, s.Border.CellStyle)

		if s.Border.Left != "" {
			paddingBox = paddingBox.GrowLeft(-1)
		}
		if s.Border.Right != "" {
			paddingBox = paddingBox.GrowRight(-1)
		}

		if s.Border.Top != "" {
			paddingBox = paddingBox.GrowTop(-1)
		}
		if s.Border.Bottom != "" {
			paddingBox = paddingBox.GrowBottom(-1)
		}
	}

	subsurface = draw.NewSubSurface(surface, paddingBox)

	contentBox := paddingBox
	if s.Padding != nil {
		// TODO: set background and foreground colors.
		contentBox = contentBox.GrowLeft(-s.Padding.Left)
		contentBox = contentBox.GrowTop(-s.Padding.Top)
		contentBox = contentBox.GrowRight(-s.Padding.Right)
		contentBox = contentBox.GrowBottom(-s.Padding.Bottom)
	}

	subsurface = draw.NewSubSurface(surface, contentBox)

	next.Draw(subsurface)
}
