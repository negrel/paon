package tui

import (
	"strings"

	"github.com/gdamore/tcell"
	"github.com/mitchellh/go-wordwrap"

	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/utils"
	"github.com/negrel/paon/internal/widgets"
)

var _ widgets.Widget = &Text{}

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
	patch.Origin = rect.Min
	patch.Frame = make([][]render.Cell, 1)

	wrapped := wordwrap.WrapString(t.content, uint(rect.Width()))
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
