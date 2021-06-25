package layout

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/styles"
)

// BoxModel define a sized box with margin, border and padding in a 2D geometric plane.
type BoxModel interface {
	// The size of the outer box (the margin box).
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

// NewBox return a new Box with the given content box.
func NewBox(rectangle geometry.Rectangle) *Box {
	return &Box{
		borderBox: rectangle,
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
	return b.MarginBox().Size()
}

// Width implements the geometry.Sized interface.
func (b *Box) Width() int {
	return b.MarginBox().Width()
}

// Height implements the geometry.Sized interface.
func (b *Box) Height() int {
	return b.MarginBox().Height()
}

// Resize change the size of the margin box.
func (b *Box) Resize(size geometry.Size) {
	diffW := size.Width() - b.Width()
	diffH := size.Height() - b.Height()

	b.borderBox.Max = b.borderBox.Max.Add(
		geometry.Pt(diffW, diffH),
	)
}

func (b *Box) ApplyMargin(style styles.Style) {
	b.applyMargin(marginOf(style))
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

func (b *Box) ApplyBorder(style styles.Style) {
	b.applyBorder(borderOf(style))
}

func (b *Box) applyBorder(border boxOffset) {
	b.paddingBoxOffset = border
}

func (b *Box) ApplyPadding(style styles.Style) {
	b.applyPadding(paddingOf(style))
}

func (b *Box) applyPadding(padding boxOffset) {
	b.contentBoxOffset = padding
}
