package widgets

import (
	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/widgets/border"
)

// Style define a basic style type for widgets.
type Style struct {
	// nil values means that styling prop must be inherited
	*marginStyle
	borderStyle *border.Border
	borderColor *colors.Color
	*paddingStyle
	*textDecoration
	foreground *colors.Color
	background *colors.Color
}

// Style implements styles.Styled.
func (s Style) Style() styles.Style {
	return s
}

type marginStyle struct {
	// TODO: support relative unit
	left, top, right, bottom int
}

func (s Style) InheritMargin() Style {
	s.marginStyle = nil
	return s
}

func (s Style) Margin(margins ...int) Style {
	if s.marginStyle == nil {
		s.marginStyle = &marginStyle{}
	}

	switch len(margins) {
	case 1:
		s.marginStyle.left = margins[0]
		s.marginStyle.top = margins[0]
		s.marginStyle.right = margins[0]
		s.marginStyle.bottom = margins[0]

	case 2:
		s.marginStyle.left = margins[1]
		s.marginStyle.top = margins[0]
		s.marginStyle.right = margins[1]
		s.marginStyle.bottom = margins[0]

	case 3:
		s.marginStyle.left = margins[1]
		s.marginStyle.top = margins[0]
		s.marginStyle.right = margins[1]
		s.marginStyle.bottom = margins[2]

	case 4:
		s.marginStyle.left = margins[3]
		s.marginStyle.top = margins[0]
		s.marginStyle.right = margins[1]
		s.marginStyle.bottom = margins[2]
	}
	return s
}

func (s Style) MarginLeft(m int) Style {
	if s.marginStyle == nil {
		s.marginStyle = &marginStyle{}
	}

	s.marginStyle.left = m
	return s
}

func (s Style) MarginTop(m int) Style {
	if s.marginStyle == nil {
		s.marginStyle = &marginStyle{}
	}

	s.marginStyle.top = m
	return s
}

func (s Style) MarginRight(m int) Style {
	if s.marginStyle == nil {
		s.marginStyle = &marginStyle{}
	}

	s.marginStyle.right = m
	return s
}

func (s Style) MarginBottom(m int) Style {
	if s.marginStyle == nil {
		s.marginStyle = &marginStyle{}
	}

	s.marginStyle.bottom = m
	return s
}

func (s Style) MarginX(m int) Style {
	if s.marginStyle == nil {
		s.marginStyle = &marginStyle{}
	}

	s.marginStyle.left = m
	s.marginStyle.right = m

	return s
}

func (s Style) MarginY(m int) Style {
	if s.marginStyle == nil {
		s.marginStyle = &marginStyle{}
	}

	s.marginStyle.top = m
	s.marginStyle.bottom = m

	return s
}

func (s Style) InheritBorder() Style {
	s.borderStyle = nil
	s.borderColor = nil
	return s
}

func (s Style) Border(b border.Border) Style {
	s.borderStyle = &b

	return s
}

func (s Style) BorderColor(c colors.Color) Style {
	s.borderColor = &c

	return s
}

type paddingStyle struct {
	// TODO: support relative unit
	left, top, right, bottom int
}

func (s Style) InheritPadding() Style {
	s.paddingStyle = nil
	return s
}

func (s Style) Padding(paddings ...int) Style {
	if s.paddingStyle == nil {
		s.paddingStyle = &paddingStyle{}
	}

	switch len(paddings) {
	case 1:
		s.paddingStyle.left = paddings[0]
		s.paddingStyle.top = paddings[0]
		s.paddingStyle.right = paddings[0]
		s.paddingStyle.bottom = paddings[0]

	case 2:
		s.paddingStyle.left = paddings[1]
		s.paddingStyle.top = paddings[0]
		s.paddingStyle.right = paddings[1]
		s.paddingStyle.bottom = paddings[0]

	case 3:
		s.paddingStyle.left = paddings[1]
		s.paddingStyle.top = paddings[0]
		s.paddingStyle.right = paddings[1]
		s.paddingStyle.bottom = paddings[2]

	default:
		s.paddingStyle.left = paddings[3]
		s.paddingStyle.top = paddings[0]
		s.paddingStyle.right = paddings[1]
		s.paddingStyle.bottom = paddings[2]
	}

	return s
}

func (s Style) PaddingLeft(l int) Style {
	if s.paddingStyle == nil {
		s.paddingStyle = &paddingStyle{}
	}

	s.paddingStyle.left = l
	return s
}

func (s Style) PaddingTop(l int) Style {
	if s.paddingStyle == nil {
		s.paddingStyle = &paddingStyle{}
	}

	s.paddingStyle.top = l
	return s
}

func (s Style) PaddingRight(l int) Style {
	if s.paddingStyle == nil {
		s.paddingStyle = &paddingStyle{}
	}

	s.paddingStyle.right = l
	return s
}

func (s Style) PaddingBottom(l int) Style {
	if s.paddingStyle == nil {
		s.paddingStyle = &paddingStyle{}
	}

	s.paddingStyle.bottom = l
	return s
}

func (s Style) PaddingX(l int) Style {
	if s.paddingStyle == nil {
		s.paddingStyle = &paddingStyle{}
	}

	s.paddingStyle.left = l
	s.paddingStyle.right = l
	return s
}

func (s Style) PaddingY(l int) Style {
	if s.paddingStyle == nil {
		s.paddingStyle = &paddingStyle{}
	}

	s.paddingStyle.top = l
	s.paddingStyle.bottom = l
	return s
}

func (s Style) InheritBackground() Style {
	s.background = nil
	return s
}

func (s Style) Background(c colors.Color) Style {
	s.background = &c

	return s
}

func (s Style) InheritForeground() Style {
	s.foreground = nil
	return s
}

func (s Style) Foreground(c colors.Color) Style {
	s.foreground = &c

	return s
}

type textDecoration struct {
	bold, blink, reverse, underline,
	dim, italic, strikeThrough bool
}

func (s Style) InheritTextDecoration() Style {
	s.textDecoration = nil
	return s
}

func (s Style) Bold(b bool) Style {
	if s.textDecoration == nil {
		s.textDecoration = &textDecoration{}
	}

	s.textDecoration.bold = b
	return s
}

func (s Style) Blink(b bool) Style {
	if s.textDecoration == nil {
		s.textDecoration = &textDecoration{}
	}

	s.textDecoration.blink = b
	return s
}

func (s Style) Reverse(b bool) Style {
	if s.textDecoration == nil {
		s.textDecoration = &textDecoration{}
	}

	s.textDecoration.reverse = b
	return s
}

func (s Style) Underline(b bool) Style {
	if s.textDecoration == nil {
		s.textDecoration = &textDecoration{}
	}

	s.textDecoration.underline = b
	return s
}

func (s Style) Dim(b bool) Style {
	if s.textDecoration == nil {
		s.textDecoration = &textDecoration{}
	}

	s.textDecoration.dim = b
	return s
}

func (s Style) Italic(b bool) Style {
	if s.textDecoration == nil {
		s.textDecoration = &textDecoration{}
	}

	s.textDecoration.italic = b
	return s
}

func (s Style) StrikeThrough(b bool) Style {
	if s.textDecoration == nil {
		s.textDecoration = &textDecoration{}
	}

	s.textDecoration.strikeThrough = b
	return s
}

// Compute implements styles.Style.
func (s Style) Compute() styles.ComputedStyle {
	result := styles.ComputedStyle{}

	if s.background != nil {
		result.CellStyle.Background = *s.background
	}
	if s.foreground != nil {
		result.CellStyle.Foreground = *s.foreground
	}
	if s.marginStyle != nil {
		result.MarginStyle = styles.MarginPaddingStyle{
			Left:   s.marginStyle.left,
			Top:    s.marginStyle.top,
			Right:  s.marginStyle.right,
			Bottom: s.marginStyle.bottom,
		}
	}
	if s.borderStyle != nil {
		result.BorderStyle = styles.BorderStyle{
			Border: *s.borderStyle,
			CellStyle: draw.CellStyle{
				Background: result.Background,
			},
		}
	}
	if s.borderColor != nil {
		result.BorderStyle.CellStyle.Foreground = *s.foreground
	}
	if s.paddingStyle != nil {
		result.PaddingStyle = styles.MarginPaddingStyle{
			Left:   s.paddingStyle.left,
			Top:    s.paddingStyle.top,
			Right:  s.paddingStyle.right,
			Bottom: s.paddingStyle.bottom,
		}
	}

	if s.textDecoration != nil {
		result.CellStyle.Blink = s.textDecoration.blink
		result.CellStyle.Bold = s.textDecoration.bold
		result.CellStyle.Dim = s.textDecoration.dim
		result.CellStyle.Italic = s.textDecoration.italic
		result.CellStyle.Reverse = s.textDecoration.reverse
		result.CellStyle.Underline = s.textDecoration.underline
		result.CellStyle.StrikeThrough = s.textDecoration.strikeThrough
	}

	return result
}
