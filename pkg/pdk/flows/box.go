package flows

import "github.com/negrel/paon/internal/geometry"

// BoxModel define a sized box with margin, border and padding in a 2D geometric plane.
type BoxModel interface {
	geometry.Sized

	MarginBox() geometry.Rectangle
	BorderBox() geometry.Rectangle
	PaddingBox() geometry.Rectangle
	ContentBox() geometry.Rectangle
}

type Box struct {
	borderBox geometry.Rectangle
	marginBoxOffset,
	paddingBoxOffset,
	contentBoxOffset boxOffset
}

func NewBox(bounds geometry.Rectangle) *Box {
	return &Box{
		borderBox: bounds,
	}
}

// MarginBox implements the BoxModel interface.
func (b *Box) MarginBox() geometry.Rectangle {
	return b.marginBoxOffset.applyOn(b.borderBox)
}

// BorderBox implements the BoxModel interface.
func (b *Box) BorderBox() geometry.Rectangle {
	return b.borderBox
}

// PaddingBox implements the BoxModel interface.
func (b *Box) PaddingBox() geometry.Rectangle {
	return b.paddingBoxOffset.applyOn(b.BorderBox())
}

// ContentBox implements the BoxModel interface.
func (b *Box) ContentBox() geometry.Rectangle {
	return b.contentBoxOffset.applyOn(b.PaddingBox())
}

// Size implements the geometry.Sized interface.
func (b *Box) Size() geometry.Size {
	return b.BorderBox().Size()
}

// Width implements the geometry.Sized interface.
func (b *Box) Width() int {
	return b.BorderBox().Width()
}

// Height implements the geometry.Sized interface.
func (b *Box) Height() int {
	return b.BorderBox().Height()
}

func (b *Box) resize(size geometry.Size) {
	b.borderBox.Max = b.borderBox.Min.Add(
		geometry.Point(size),
	)
}

func (b *Box) applyMargin(margin boxOffset) {
	b.borderBox.Min = b.borderBox.Min.Add(
		geometry.Pt(b.marginBoxOffset.left(), b.marginBoxOffset.top()),
	)
	b.borderBox.Max = b.borderBox.Max.Add(
		geometry.Pt(b.marginBoxOffset.left(), b.marginBoxOffset.top()),
	)

	b.marginBoxOffset = margin.reverse()

	b.borderBox.Min = b.borderBox.Min.Add(
		geometry.Pt(margin.left(), margin.top()),
	)
	b.borderBox.Max = b.borderBox.Max.Add(
		geometry.Pt(margin.left(), margin.top()),
	)
}

func (b *Box) applyBorder(border boxOffset) {
	b.paddingBoxOffset = border
}

func (b *Box) applyPadding(padding boxOffset) {
	b.contentBoxOffset = padding
}
