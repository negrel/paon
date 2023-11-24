package styles

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/widgets/border"
)

// Draw draws margin, border and padding onto the given surface and return a subsurface
// for content box.
func Draw(s ComputedStyle, surface draw.Surface) draw.Surface {
	surfaceSize := surface.Size()
	borderBox := geometry.Rect(0, 0, surfaceSize.Width(), surfaceSize.Height())
	if (s.MarginStyle != MarginPaddingStyle{}) {
		borderBox = borderBox.GrowLeft(-s.MarginStyle.Left)
		borderBox = borderBox.GrowTop(-s.MarginStyle.Top)
		borderBox = borderBox.GrowRight(-s.MarginStyle.Right)
		borderBox = borderBox.GrowBottom(-s.MarginStyle.Bottom)
	}

	subsurface := draw.NewSubSurface(surface, borderBox)
	if (s.CellStyle != draw.CellStyle{}) {
		for w := 0; w < subsurface.Size().Width(); w++ {
			for h := 0; h < subsurface.Size().Height(); h++ {
				subsurface.Set(geometry.NewVec2D(w, h), draw.Cell{
					Style: s.CellStyle,
				})
			}
		}
	}

	paddingBox := borderBox

	if (s.BorderStyle.Border != border.Border{}) {
		border.DrawBorder(subsurface, s.BorderStyle.Border, s.BorderStyle.CellStyle)

		if s.BorderStyle.Left != "" {
			paddingBox = paddingBox.GrowLeft(-1)
		}
		if s.BorderStyle.Right != "" {
			paddingBox = paddingBox.GrowRight(-1)
		}

		if s.BorderStyle.Top != "" {
			paddingBox = paddingBox.GrowTop(-1)
		}
		if s.BorderStyle.Bottom != "" {
			paddingBox = paddingBox.GrowBottom(-1)
		}
	}

	subsurface = draw.NewSubSurface(surface, paddingBox)

	contentBox := paddingBox
	if (s.PaddingStyle != MarginPaddingStyle{}) {
		// TODO: set background and foreground colors.
		contentBox = contentBox.GrowLeft(-s.PaddingStyle.Left)
		contentBox = contentBox.GrowTop(-s.PaddingStyle.Top)
		contentBox = contentBox.GrowRight(-s.PaddingStyle.Right)
		contentBox = contentBox.GrowBottom(-s.PaddingStyle.Bottom)
	}

	subsurface = draw.NewSubSurface(surface, contentBox)

	return subsurface
}
