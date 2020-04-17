package style

import "github.com/gdamore/tcell"

// type property struct {
// 	value bool
// 	unset bool
// }

// // Theme define the style property of a widget
// type Theme struct {
// 	Foreground Color
// 	Background Color
// 	Bold       property
// 	Blink      property
// 	Reverse    property
// 	Underline  property
// 	Dim        property
// }

type Theme = tcell.Style

var DefaultTheme Theme = tcell.StyleDefault

// // DefaultTheme is the default theme with default
// // property value.
// var DefaultTheme = Theme{
// 	Foreground: DefaultColor,
// 	Background: DefaultColor,
// 	Bold:       property{false, true},
// 	Blink:      property{false, true},
// 	Reverse:    property{false, true},
// 	Underline:  property{false, true},
// 	Dim:        property{false, true},
// }

// // Compute the style for rendering/painting.
// // <7b flags><1b><24b fgcolor><7b attr><1b><24b bgcolor>
// func (t Theme) Compute() tcell.Style {
// 	var style tcell.Style

// 	if t.Bold.value {
// 		style |= 1 << 25
// 	}

// 	if t.Blink.value {
// 		style |= 1 << 26
// 	}

// 	if t.Reverse.value {
// 		style |= 1 << 27
// 	}

// 	if t.Underline.value {
// 		style |= 1 << 28
// 	}

// 	return style
// }
