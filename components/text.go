package component

import "github.com/negrel/ginger/tree"

// Text is a text component
type Text struct {
	*tree.Widget

	txt string
}

// NewText return a Text component
func NewText(txt string) *Text {
	return &Text{
		Widget: &tree.Widget{},
		txt:    txt,
	}
}

// Draw the text component
func (t *Text) Draw() {

}
