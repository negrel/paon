package tui

import (
	"strings"

	"github.com/gdamore/tcell"

	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/utils"
	"github.com/negrel/paon/internal/widgets"
)

var _ widgets.Widget = &Text{}

// Text define a simple TUI text element.
type Text struct {
	*widgets.Node
	content string
}

// TextWidget return a new Text element.
func TextWidget(content string) *Text {
	return &Text{
		Node:    widgets.NewNodeWidget("text"),
		content: content,
	}
}

// Render implements the widgets.Widget interface.
func (t Text) Render(rect utils.Rectangle) (patch render.Patch) {
	maxWidth := rect.Width()
	wrapped := utils.WordWrap(t.content, maxWidth)

	patch.Origin = rect.Min
	patch.Frame = make([][]render.Cell, maxWidth)

	for i, line := range strings.Split(wrapped, "\n") {
		patch.Frame[i] = make([]render.Cell, len(line))

		for j, r := range line {
			patch.Frame[i][j] = render.Cell{
				Style:   tcell.StyleDefault,
				Content: r,
			}
		}
	}

	return patch
}
