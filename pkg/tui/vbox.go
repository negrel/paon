package tui

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/flows"
	"github.com/negrel/paon/pkg/pdk/math"
	"github.com/negrel/paon/pkg/pdk/widgets"
)

type VBoxLayout struct {
	widgets.Layout
}

func VBox(children []widgets.Widget, opts ...widgets.Option) *VBoxLayout {
	vbox := &VBoxLayout{}
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

func (vbl *VBoxLayout) flow(constraint flows.Constraint) flows.BoxModel {
	childrenOk := true

	result := flows.Block(vbl.Style(), constraint, func(constraint flows.Constraint) flows.BoxModel {
		childrenOk = false
		childrenBoxes := vbl.flowChildren(constraint)

		width := 0
		height := 0
		for _, childBox := range childrenBoxes {
			width = math.Max(width, childBox.Width())
			height += childBox.Height()
		}

		width = math.Constrain(width, constraint.Min.Width(), constraint.Max.Width())
		height = math.Constrain(width, constraint.Min.Height(), constraint.Max.Height())

		return flows.NewBox(geometry.Rectangle{
			Min: constraint.Min.Min,
			Max: constraint.Min.Min.Add(geometry.Pt(width, height)),
		})
	})

	if childrenOk {
		vbl.flowChildren(constraint)
	}

	return result
}

func (vbl *VBoxLayout) flowChildren(constraint flows.Constraint) []flows.BoxModel {
	result := make([]flows.BoxModel, 0, 8)

	child := vbl.FirstChild()
	for child != nil {
		childBox := child.Flow(constraint)
		result = append(result, childBox)

		constraint.Min.Min = geometry.Pt(0, childBox.Height()).Add(constraint.Min.Min)
		child = child.Next()
	}

	return result
}

// draw implements the widgets.Widget interface.
func (vbl *VBoxLayout) draw(ctx draw.Context) {
	canvas := ctx.Canvas()

	child := vbl.FirstChild()
	for child != nil {
		childBox := child.Box()
		childCtx := canvas.NewContext(childBox.MarginBox())
		widgets.DrawBoxOf(child, childCtx)

		childCtx = childCtx.SubContext(childBox.ContentBox())
		child.Draw(childCtx)

		child = child.Next()
	}
}
