package style

import (
	"github.com/gdamore/tcell"
)

// Theme is used to style the content of a widget.
// It define the following property :
//
// Foreground color (text color)
//
// Background color
//
// Bold text
//
// Blink text
//
// Reverse text
//
// Underlin text
//
// Dim text
type Theme struct {
	foreground Color
	background Color
	bold       bool
	blink      bool
	reverse    bool
	underline  bool
	dim        bool
}

// DefaultTheme leave the terminal theme/style unchanged
// from whatever your terminal default theme is.
var DefaultTheme = Theme{
	foreground: DefaultColor,
	background: DefaultColor,
	bold:       false,
	blink:      false,
	reverse:    false,
	underline:  false,
	dim:        false,
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Foreground set the foreground color and return
// the new theme.
func (t Theme) Foreground(c Color) Theme {
	t.foreground = c
	return t
}

// Background set the background color and return
// the new theme.
func (t Theme) Background(c Color) Theme {
	t.background = c
	return t
}

// Bold returns a new style based on s, with the bold
// attribute set as requested.
func (t Theme) Bold(on bool) Theme {
	t.bold = on
	return t
}

// Blink returns a new style based on s, with the blink
// attribute set as requested.
func (t Theme) Blink(on bool) Theme {
	t.blink = on
	return t
}

// Reverse returns a new style based on s, with the reverse
// attribute set as requested.
func (t Theme) Reverse(on bool) Theme {
	t.reverse = on
	return t
}

// Underline returns a new style based on s, with the underline
// attribute set as requested.
func (t Theme) Underline(on bool) Theme {
	t.underline = on
	return t
}

// Dim returns a new style based on s, with the dim
// attribute set as requested.
func (t Theme) Dim(on bool) Theme {
	t.dim = on
	return t
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Compute implements Rendable interface.
func (t Theme) Compute() tcell.Style {
	var style tcell.Style = tcell.StyleDefault

	if t.foreground != DefaultColor {
		style = style.Foreground(tcell.NewHexColor(t.foreground))
	}

	if t.background != DefaultColor {
		style = style.Background(tcell.NewHexColor(t.background))
	}

	style = style.Bold(t.bold).Blink(t.blink).Reverse(t.reverse).Underline(t.underline).Dim(t.dim)

	return style
}
