package styles

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/widgets/border"
)

type Option func(*Style)

type MarginStyle struct {
	Left, Top, Right, Bottom int
}

func Margin(margins ...int) Option {
	switch len(margins) {
	case 0:
		panic("you must pass at least one argument to styles.Margin")

	case 1:
		return func(s *Style) {
			if s.Margin == nil {
				s.Margin = &MarginStyle{}
			}

			s.Margin.Left = margins[0]
			s.Margin.Top = margins[0]
			s.Margin.Right = margins[0]
			s.Margin.Bottom = margins[0]
		}
	case 2:
		return func(s *Style) {
			if s.Margin == nil {
				s.Margin = &MarginStyle{}
			}

			s.Margin.Left = margins[1]
			s.Margin.Top = margins[0]
			s.Margin.Right = margins[1]
			s.Margin.Bottom = margins[0]
		}

	case 3:
		return func(s *Style) {
			if s.Margin == nil {
				s.Margin = &MarginStyle{}
			}

			s.Margin.Left = margins[1]
			s.Margin.Top = margins[0]
			s.Margin.Right = margins[1]
			s.Margin.Bottom = margins[2]
		}

	default:
		return func(s *Style) {
			if s.Margin == nil {
				s.Margin = &MarginStyle{}
			}

			s.Margin.Left = margins[3]
			s.Margin.Top = margins[0]
			s.Margin.Right = margins[1]
			s.Margin.Bottom = margins[2]
		}
	}
}

func MarginLeft(m int) Option {
	return func(s *Style) {
		if s.Margin == nil {
			s.Margin = &MarginStyle{}
		}

		s.Margin.Left = m
	}
}

func MarginTop(m int) Option {
	return func(s *Style) {
		if s.Margin == nil {
			s.Margin = &MarginStyle{}
		}

		s.Margin.Top = m
	}
}

func MarginRight(m int) Option {
	return func(s *Style) {
		if s.Margin == nil {
			s.Margin = &MarginStyle{}
		}

		s.Margin.Right = m
	}
}

func MarginBottom(m int) Option {
	return func(s *Style) {
		if s.Margin == nil {
			s.Margin = &MarginStyle{}
		}

		s.Margin.Bottom = m
	}
}

func MarginX(m int) Option {
	return func(s *Style) {
		if s.Margin == nil {
			s.Margin = &MarginStyle{}
		}

		s.Margin.Left = m
		s.Margin.Right = m
	}
}

func MarginY(m int) Option {
	return func(s *Style) {
		if s.Margin == nil {
			s.Margin = &MarginStyle{}
		}

		s.Margin.Top = m
		s.Margin.Bottom = m
	}
}

type PaddingStyle struct {
	Left, Top, Right, Bottom int
}

func Padding(paddings ...int) Option {
	switch len(paddings) {
	case 0:
		panic("you must pass at least one argument to styles.Margin")

	case 1:
		return func(s *Style) {
			s.Padding.Left = paddings[0]
			s.Padding.Top = paddings[0]
			s.Padding.Right = paddings[0]
			s.Padding.Bottom = paddings[0]
		}
	case 2:
		return func(s *Style) {
			s.Padding.Left = paddings[1]
			s.Padding.Top = paddings[0]
			s.Padding.Right = paddings[1]
			s.Padding.Bottom = paddings[0]
		}

	case 3:
		return func(s *Style) {
			s.Padding.Left = paddings[1]
			s.Padding.Top = paddings[0]
			s.Padding.Right = paddings[1]
			s.Padding.Bottom = paddings[2]
		}

	default:
		return func(s *Style) {
			s.Padding.Left = paddings[3]
			s.Padding.Top = paddings[0]
			s.Padding.Right = paddings[1]
			s.Padding.Bottom = paddings[2]
		}
	}
}

func PaddingLeft(l int) Option {
	return func(s *Style) {
		if s.Padding == nil {
			s.Padding = &PaddingStyle{}
		}

		s.Padding.Left = l
	}
}

func PaddingTop(l int) Option {
	return func(s *Style) {
		if s.Padding == nil {
			s.Padding = &PaddingStyle{}
		}

		s.Padding.Top = l
	}
}

func PaddingRight(l int) Option {
	return func(s *Style) {
		if s.Padding == nil {
			s.Padding = &PaddingStyle{}
		}

		s.Padding.Right = l
	}
}

func PaddingBottom(l int) Option {
	return func(s *Style) {
		if s.Padding == nil {
			s.Padding = &PaddingStyle{}
		}

		s.Padding.Bottom = l
	}
}

func PaddingX(l int) Option {
	return func(s *Style) {
		if s.Padding == nil {
			s.Padding = &PaddingStyle{}
		}

		s.Padding.Left = l
		s.Padding.Right = l
	}
}

func PaddingY(l int) Option {
	return func(s *Style) {
		if s.Padding == nil {
			s.Padding = &PaddingStyle{}
		}

		s.Padding.Top = l
		s.Padding.Bottom = l
	}
}

type BorderStyle struct {
	border.Border
	draw.CellStyle
}

func Border(b border.Border) Option {
	return func(s *Style) {
		if s.Border == nil {
			s.Border = &BorderStyle{}
		}

		s.Border.Border = b
	}
}

func Extra(v any) Option {
	return func(s *Style) {
		s.Extras = v
	}
}
