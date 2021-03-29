package tui

import (
	"github.com/mitchellh/go-wordwrap"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/flows"
	"github.com/negrel/paon/pkg/pdk/math"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/styles/value"
	"github.com/negrel/paon/pkg/pdk/widgets"
	"github.com/negrel/paon/pkg/style"
	"strings"
)

var _ widgets.Widget = &TextWidget{}

type TextWidget struct {
	widgets.Widget

	content    string
	fmtContent []string
}

// Text returns a new TextWidget with the given content.
func Text(content string, options ...widgets.Option) *TextWidget {
	tw := &TextWidget{
		content: content,
	}

	defaultStyle := styles.NewStyle()
	defaultStyle.Set(style.FgColor(value.ColorFromHex(0xFFFFFF)))

	tw.Widget = widgets.NewWidget(
		"text",
		widgets.PrependOptions(
			options,
			widgets.Algo(tw.flow),
			widgets.Script(tw.draw),
			widgets.DefaultStyle(defaultStyle),
		)...,
	)

	return tw
}

func (tw *TextWidget) flowText(width uint) int {
	tw.fmtContent = strings.Split(wordwrap.WrapString(tw.content, width), "\n")

	return len(tw.fmtContent)
}

func (tw *TextWidget) flow(constraint flows.Constraint) flows.BoxModel {
	return flows.Block(tw.Style(), constraint, func(constraint flows.Constraint) flows.BoxModel {
		height := tw.flowText(uint(constraint.Max.Width()))

		width := math.Constrain(len(tw.content), constraint.Min.Width(), constraint.Max.Width())
		height = math.Constrain(height, constraint.Min.Height(), constraint.Max.Height())

		box := flows.NewBox(geometry.Rectangle{
			Min: constraint.Min.Min,
			Max: constraint.Min.Min.Add(geometry.Pt(width, height)),
		})
		box.ApplyMargin(tw.Style())
		box.ApplyBorder(tw.Style())
		box.ApplyPadding(tw.Style())

		return box
	})
}

func (tw *TextWidget) draw(ctx draw.Context) {
	color, ok := tw.Style().Get(property.ForegroundColor()).(property.Color)
	if ok {
		ctx.SetFillColor(color.Color)
	}

	for i, line := range tw.fmtContent {
		ctx.FillTextH(geometry.Pt(0, i), line)
	}

	ctx.Commit()
}
