package flows

import "github.com/negrel/paon/internal/geometry"

type Box interface {
	geometry.Sized

	MarginBox() geometry.Rectangle
	BorderBox() geometry.Rectangle
	PaddingBox() geometry.Rectangle
	ContentBox() geometry.Rectangle
}

type boxModel struct {
	borderBox geometry.Rectangle
	marginBoxOffset,
	paddingBoxOffset,
	contentBoxOffset boxOffset
}

func newBoxModel(bounds geometry.Rectangle) *boxModel {
	return &boxModel{
		borderBox: bounds,
	}
}

// MarginBox implements the Box interface.
func (bm *boxModel) MarginBox() geometry.Rectangle {
	return bm.marginBoxOffset.applyOn(bm.borderBox)
}

// BorderBox implements the Box interface.
func (bm *boxModel) BorderBox() geometry.Rectangle {
	return bm.borderBox
}

// PaddingBox implements the Box interface.
func (bm *boxModel) PaddingBox() geometry.Rectangle {
	return bm.paddingBoxOffset.applyOn(bm.BorderBox())
}

// ContentBox implements the Box interface.
func (bm *boxModel) ContentBox() geometry.Rectangle {
	return bm.contentBoxOffset.applyOn(bm.PaddingBox())
}

// Size implements the geometry.Sized interface.
func (bm *boxModel) Size() geometry.Size {
	return bm.BorderBox().Size()
}

// Width implements the geometry.Sized interface.
func (bm *boxModel) Width() int {
	return bm.BorderBox().Width()
}

// Height implements the geometry.Sized interface.
func (bm *boxModel) Height() int {
	return bm.BorderBox().Height()
}

func (bm *boxModel) resize(size geometry.Size) {
	bm.borderBox.Max = bm.borderBox.Min.Add(
		geometry.Point(size),
	)
}

func (bm *boxModel) applyMargin(margin boxOffset) {
	bm.borderBox.Min = bm.borderBox.Min.Add(
		geometry.Pt(bm.marginBoxOffset.left(), bm.marginBoxOffset.top()),
	)
	bm.borderBox.Max = bm.borderBox.Max.Add(
		geometry.Pt(bm.marginBoxOffset.left(), bm.marginBoxOffset.top()),
	)

	bm.marginBoxOffset = margin.reverse()

	bm.borderBox.Min = bm.borderBox.Min.Add(
		geometry.Pt(margin.left(), margin.top()),
	)
	bm.borderBox.Max = bm.borderBox.Max.Add(
		geometry.Pt(margin.left(), margin.top()),
	)
}

func (bm *boxModel) applyBorder(border boxOffset) {
	bm.paddingBoxOffset = border
}

func (bm *boxModel) applyPadding(padding boxOffset) {
	bm.contentBoxOffset = padding
}
