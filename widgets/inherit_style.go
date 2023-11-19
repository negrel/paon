package widgets

import (
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/tree"
)

// InheritStyle wraps a Widget and a Style to add inheritance capability to it.
type InheritStyle struct {
	Widget     Widget
	InnerStyle Style
}

// Style implements styles.Styled.
func (i InheritStyle) Style() styles.Style {
	return i
}

// Compute implements styles.Style.
func (i InheritStyle) Compute() styles.ComputedStyle {
	parentComputedStyle := ParentStyle(i.Widget.Node().Parent())
	_ = parentComputedStyle
	computedStyle := i.InnerStyle.Compute()

	if i.InnerStyle.textDecoration == nil {
		computedStyle.CellStyle.Dim = parentComputedStyle.CellStyle.Dim
		computedStyle.CellStyle.Bold = parentComputedStyle.CellStyle.Bold
		computedStyle.CellStyle.Blink = parentComputedStyle.CellStyle.Blink
		computedStyle.CellStyle.StrikeThrough = parentComputedStyle.CellStyle.StrikeThrough
		computedStyle.CellStyle.Underline = parentComputedStyle.CellStyle.Underline
		computedStyle.CellStyle.Reverse = parentComputedStyle.CellStyle.Reverse
		computedStyle.CellStyle.Italic = parentComputedStyle.CellStyle.Italic
	}
	if i.InnerStyle.background == nil {
		computedStyle.CellStyle.Background = parentComputedStyle.CellStyle.Background
	}
	if i.InnerStyle.foreground == nil {
		computedStyle.CellStyle.Foreground = parentComputedStyle.CellStyle.Foreground
	}
	if i.InnerStyle.paddingStyle == nil {
		computedStyle.PaddingStyle = parentComputedStyle.PaddingStyle
	}
	if i.InnerStyle.borderStyle == nil {
		computedStyle.BorderStyle = parentComputedStyle.BorderStyle
	}
	computedStyle.BorderStyle.Background = computedStyle.Background
	if i.InnerStyle.marginStyle == nil {
		computedStyle.MarginStyle = parentComputedStyle.MarginStyle
	}

	return computedStyle
}

func ParentStyle(parent *tree.Node[Widget]) styles.ComputedStyle {
	if parent == nil {
		return styles.ComputedStyle{}
	}

	styled := parent.Unwrap()
	if s := styled.Style(); s != nil {
		return s.Compute()
	}

	return ParentStyle(parent.Parent())
}
