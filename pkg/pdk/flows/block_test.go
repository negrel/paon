package flows

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testBlockApply(t *testing.T, constraint Constraint, s styles.Style) {
	width := constraint.Max.Width()
	height := ComputeConstrainedHeight(s, constraint)
	margin := marginOf(s)
	border := borderOf(s)
	padding := paddingOf(s)

	b := makeBlock()
	actual := b.Apply(s, constraint)

	expectedMarginBox := geometry.Rectangle{
		Min: constraint.Min.Min,
		Max: geometry.Pt(width+margin.x(), height+margin.y()),
	}
	assert.Equal(t, expectedMarginBox, actual.MarginBox())

	expectedBorderBox := geometry.Rectangle{
		Min: geometry.Pt(expectedMarginBox.Min.X()+margin.left(), expectedMarginBox.Min.Y()+margin.top()),
		Max: geometry.Pt(expectedMarginBox.Max.X()-margin.right(), expectedMarginBox.Max.Y()-margin.bottom()),
	}
	assert.Equal(t, expectedBorderBox, actual.BorderBox())

	expectedPaddingBox := geometry.Rectangle{
		Min: geometry.Pt(expectedBorderBox.Min.X()+border.left(), expectedBorderBox.Min.Y()+border.top()),
		Max: geometry.Pt(expectedBorderBox.Max.X()-border.right(), expectedBorderBox.Max.Y()-border.bottom()),
	}
	assert.Equal(t, expectedPaddingBox, actual.PaddingBox())

	expectedContentBox := geometry.Rectangle{
		Min: geometry.Pt(expectedPaddingBox.Min.X()+padding.left(), expectedPaddingBox.Min.Y()+padding.top()),
		Max: geometry.Pt(expectedPaddingBox.Max.X()-padding.right(), expectedPaddingBox.Max.Y()-padding.bottom()),
	}
	assert.Equal(t, expectedContentBox, actual.ContentBox())
}

func TestBlock_Apply_NoStyle(t *testing.T) {
	s := styles.MakeStyle()
	constraint := makeConstraint(5, 7, 15, 15)

	testBlockApply(t, constraint, s)
}

func TestBlock_Apply_MarginOnly(t *testing.T) {
	s := styles.MakeStyle()
	margin := makeBoxOffset(1, 5, 2, 7)
	applyMargin(s, margin)

	constraint := makeConstraint(5, 7, 15, 15)
	testBlockApply(t, constraint, s)
}

func TestBlock_Apply_BorderOnly(t *testing.T) {
	s := styles.MakeStyle()
	border := makeBoxOffset(1, 5, 2, 7)
	applyBorder(s, border)

	constraint := makeConstraint(5, 7, 25, 25)
	testBlockApply(t, constraint, s)
}

func TestBlock_Apply_PaddingOnly(t *testing.T) {
	s := styles.MakeStyle()
	padding := makeBoxOffset(1, 5, 2, 7)
	applyPadding(s, padding)

	constraint := makeConstraint(5, 7, 25, 25)
	testBlockApply(t, constraint, s)
}
