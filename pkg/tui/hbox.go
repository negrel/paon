package tui

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/flows"
	"github.com/negrel/paon/pkg/pdk/math"
	"github.com/negrel/paon/pkg/pdk/widgets"
)

type HBoxLayout struct {
	widgets.Layout
}

func HBox(children []widgets.Widget, opts ...widgets.Option) *HBoxLayout {
	vbox := &HBoxLayout{}
	vbox.Layout = widgets.NewLayout(
		widgets.PrependOptions(opts,
			widgets.Algo(vbox.flow), widgets.Script(vbox.draw),
		)...,
	)

	for _, child := range children {
		_, err := vbox.AppendChild(child)
		if err != nil {
			panic(err)
		}
	}

	return vbox
}

func (hbl *HBoxLayout) flow(constraint flows.Constraint) flows.BoxModel {
	childrenOk := true

	result := flows.Block(hbl.Style(), constraint, func(constraint flows.Constraint) flows.BoxModel {
		childrenOk = false
		childrenBoxes := hbl.flowChildren(constraint)

		width := 0
		height := 0
		for _, childBox := range childrenBoxes {
			width += childBox.Width()
			height = math.Max(height, childBox.Height())
		}

		width = math.Constrain(width, constraint.Min.Width(), constraint.Max.Width())
		height = math.Constrain(width, constraint.Min.Height(), constraint.Max.Height())

		return flows.NewBox(geometry.Rectangle{
			Min: constraint.Min.Min,
			Max: constraint.Min.Min.Add(geometry.Pt(width, height)),
		})
	})

	if childrenOk {
		hbl.flowChildren(constraint)
	}

	return result
}

func (hbl *HBoxLayout) flowChildren(constraint flows.Constraint) []flows.BoxModel {
	result := make([]flows.BoxModel, 0, 8)

	child := hbl.FirstChild()
	for child != nil {
		childBox := child.Flow(constraint)
		result = append(result, childBox)

		constraint.Min.Min = geometry.Pt(childBox.Width(), 0).Add(constraint.Min.Min)
		child = child.Next()
	}

	return result
}

// draw implements the widgets.Widget interface.
func (hbl *HBoxLayout) draw(ctx draw.Context) {
	canvas := ctx.Canvas()

	child := hbl.FirstChild()
	for child != nil {
		childBox := child.Box()
		childCtx := canvas.NewContext(childBox.MarginBox())
		widgets.DrawBoxOf(child, childCtx)

		childCtx = childCtx.SubContext(childBox.ContentBox())
		child.Draw(childCtx)

		child = child.Next()
	}
}