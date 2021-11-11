package layout

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

// Constraint define layout constraint.
type Constraint struct {
	MinSize, MaxSize geometry.Size
	ParentSize       geometry.Size
	RootSize         geometry.Size
}

// SetMin sets the constraint minimum rectangle.
func (c Constraint) SetMin(min geometry.Size) Constraint {
	c.MinSize = min
	return c
}

// SetMax sets the constraint maximum rectangle.
func (c Constraint) SetMax(max geometry.Size) Constraint {
	c.MaxSize = max
	return c
}

// Equals returns true if the given Constraint is equal to this Constraint.
func (c Constraint) Equals(other Constraint) bool {
	return c.MinSize.Equals(other.MinSize) && c.MaxSize.Equals(other.MaxSize) &&
		c.ParentSize.Equals(other.ParentSize) && c.RootSize.Equals(other.RootSize)
}

// ToCellUnit converts the given value to a Cell based value.Unit and returns it.
func (c Constraint) ToCellUnit(unit *property.IntUnit) int {
	switch unit.Unit() {
	case property.CellUnit:
		return unit.Value()

	case property.PercentageWidthUnit:
		return c.ParentSize.Width() * unit.Value() / 100

	case property.PercentageHeightUnit:
		return c.ParentSize.Height() * unit.Value() / 100

	case property.WindowWidthUnit:
		return c.RootSize.Width() * unit.Value() / 100

	case property.WindowHeightUnit:
		return c.RootSize.Height() * unit.Value() / 100

	default:
		panic("invalid unit")
	}
}

// NewFromStyle returns a stricter Constraint based on the given style.
// The new constraint will either use MinWidth/MinHeight property as minimum
// size or the current constraint size. The same is done for maximum size.
func (c Constraint) NewFromStyle(style styles.Style) Constraint {
	newc := Constraint{
		ParentSize: c.ParentSize,
		RootSize:   c.RootSize,
	}

	minWidth := c.MinSize.Width()
	maxWidth := c.MaxSize.Width()

	if minWidthProp := style.IntUnit(property.MinWidth()); minWidthProp != nil {
		minWidth = c.ApplyOnWidthProp(minWidthProp)
	}

	if maxWidthProp := style.IntUnit(property.MaxWidth()); maxWidthProp != nil {
		maxWidth = c.ApplyOnWidthProp(maxWidthProp)
	}

	minHeight := c.MinSize.Height()
	maxHeight := c.MaxSize.Height()
	if minHeightProp := style.IntUnit(property.MinHeight()); minHeightProp != nil {
		minHeight = c.ApplyOnHeightProp(minHeightProp)
	}

	maxHeightProp := style.IntUnit(property.MaxHeight())
	if maxHeightProp != nil {
		maxHeight = c.ApplyOnHeightProp(maxHeightProp)
	}

	newc.MinSize = geometry.NewSize(minWidth, minHeight)
	newc.MaxSize = geometry.NewSize(maxWidth, maxHeight)

	return newc
}

// ApplyOnSize applies the size constraints on the given size and return
// the constrained size.
func (c Constraint) ApplyOnSize(size geometry.Size) geometry.Size {
	return geometry.NewSize(
		c.ApplyOnWidth(size.Width()),
		c.ApplyOnHeight(size.Height()),
	)
}

// ApplyOnWidthProp applies the width constraints on the given width property and
// return the constrained width as an int (CellUnit).
func (c Constraint) ApplyOnWidthProp(widthProp *property.IntUnit) int {
	width := c.ToCellUnit(widthProp)
	return c.ApplyOnWidth(width)
}

// ApplyOnWidth applies the width constraints on the given width and return the
// constrained width.
func (c Constraint) ApplyOnWidth(width int) int {
	width = math.Max(width, c.MinSize.Width())
	return math.Min(width, c.MaxSize.Width())
}

// ApplyOnHeightProp applies the height constraints on the given height property and
// return the constrained height as an int (CellUnit).
func (c Constraint) ApplyOnHeightProp(heightProp *property.IntUnit) int {
	height := c.ToCellUnit(heightProp)
	return c.ApplyOnHeight(height)
}

// ApplyOnHeight applies the height constraints on the given height and return the
// constrained height.
func (c Constraint) ApplyOnHeight(height int) int {
	height = math.Max(height, c.MinSize.Height())
	return math.Min(height, c.MaxSize.Height())
}
