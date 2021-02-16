package layout

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/styles"
)

var _ Algorithm = block{}

type block struct {
}

func makeBlock() Algorithm {
	return block{}
}

func (b block) Apply(obj styles.Stylised, constraint Constraint) Result {
	style := obj.Style()

	result := newBoxModel(constraint.Max)
	result.applyMargin(marginOf(style))
	result.applyBorder(borderOf(style))
	result.applyPadding(paddingOf(style))

	contentBoxConstraint := Constraint{
		Min: constraint.Min,
		Max: result.ContentBox(),
	}
	height := computeConstrainedHeight(style, contentBoxConstraint)
	heightDiff := height - result.BorderBox().Height()
	if heightDiff < 0 {
		result.resize(
			geometry.MakeSize(result.Width(), result.Height()+heightDiff),
		)
	}

	return result
}
