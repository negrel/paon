package tui

import (
	"github.com/negrel/paon/internal/render"
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

func (t Text) Render() render.Patch {
	panic("implement me")
}
