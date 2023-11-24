package widgets

import (
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/tree"
)

// InheritStyle define a Style wrapper that adds inheritance capability to it.
// Inheritance happens when styles.Style.Compute is called.
type InheritStyle struct {
	// NodeAccessor provide node associated with this style. It is used to retrieve
	// parent node and thus parent style.
	NodeAccessor tree.NodeAccessor
	// Style associated with this node. Nil property are inherited on Compute.
	InnerStyle Style
}

// Style implements styles.Styled.
func (i InheritStyle) Style() styles.Style {
	return i
}

// Compute implements styles.Style.
func (i InheritStyle) Compute() styles.ComputedStyle {
	parentComputedStyle := ParentStyle(i.NodeAccessor.Node().Parent())
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
	if i.InnerStyle.borderColor == nil {
		computedStyle.BorderStyle.Foreground = parentComputedStyle.BorderStyle.Foreground
	}
	computedStyle.BorderStyle.Background = computedStyle.Background
	if i.InnerStyle.marginStyle == nil {
		computedStyle.MarginStyle = parentComputedStyle.MarginStyle
	}

	return computedStyle
}

// ParentStyle finds the first ancestor of node that is styled and returns its
// computed style.
func ParentStyle(parent *tree.Node) styles.ComputedStyle {
	if parent == nil {
		return styles.ComputedStyle{}
	}

	styled := parent.Unwrap().(styles.Styled)
	if s := styled.Style(); s != nil {
		return s.Compute()
	}

	return ParentStyle(parent.Parent())
}
