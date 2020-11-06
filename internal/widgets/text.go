package widgets

import (
	"strings"

	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/utils"
)

type TextWidget struct {
	*widget

	content string
}

func Text(content string, opts ...Option) *TextWidget {
	opts = append(opts, Opt(renderTextWidget(content), 0))

	return &TextWidget{
		widget:  newWidget("text", opts...),
		content: content,
	}
}

func renderTextWidget(content string) func(render.Surface) {
	return func(buffer render.Surface) {
		content := utils.WordWrap(content, buffer.Width())
		cols := strings.Split(content, "\n")

		for i, row := range cols {
			for j, cell := range row {
				buffer.Draw(geometry.Pt(j, i), render.Cell{
					Content: cell,
				})
			}
		}
	}
}
