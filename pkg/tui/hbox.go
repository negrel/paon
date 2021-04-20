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
	hbox := &HBoxLayout{}
	hbox.Layout = widgets.NewLayout(
		"hbox",
		widgets.PrependOptions(opts,
			widgets.Algo(hbox.flow), widgets.DrawerFn(hbox.draw),
		)...,
	)

	for _, child := range children {
		err := hbox.AppendChild(child)
		if err != nil {
			panic(err)
		}
	}

	return hbox
}

func (hbl *HBoxLayout) flow(constraint flows.Constraint) flows.BoxModel {
	childrenNeedFlow := true

	result := flows.Block(hbl.Style(), constraint, func(constraint flows.Constraint) flows.BoxModel {
		childrenNeedFlow = false
		childrenBoxes := hbl.flowChildren(constraint)

		width := 0
		height := 0
		for _, childBox := range childrenBoxes {
			width += childBox.Width()
			height = math.Max(height, childBox.Height())
		}

		width = math.Constrain(width, constraint.Min.Width(), constraint.Max.Width())
		height = math.Constrain(height, constraint.Min.Height(), constraint.Max.Height())

		return flows.NewBox(geometry.Rectangle{
			Min: constraint.Min.Min,
			Max: constraint.Min.Min.Add(geometry.Pt(width, height)),
		})
	})

	if childrenNeedFlow {
		hbl.flowChildren(constraint)
	}

	return result
}

func (hbl *HBoxLayout) flowChildren(constraint flows.Constraint) []flows.BoxModel {
	result := make([]flows.BoxModel, 0, 8)

	child := hbl.FirstChild()
	for child != nil {
		childBox := child.FlowAlgo()(constraint)
		result = append(result, childBox)

		constraint.Min.Min = geometry.Pt(childBox.Width(), 0).Add(constraint.Min.Min)
		child = child.Next()
	}

	return result
}

// draw implements the widgets.Widget interface.
func (hbl *HBoxLayout) draw(ctx draw.Context) {
	child := hbl.FirstChild()
	for child != nil {
		childBox := child.Box()

		childCtx := ctx.SubContext(childBox.MarginBox())
		child.Drawer().Draw(childCtx)

		child = child.Next()
	}
}
