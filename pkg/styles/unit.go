package styles

type Unit int8

const (
	CellUnit Unit = iota
	// Percent is a relative unit based on the parent of the widget
	PercentUnit
	// ViewportWidth is a percentage based on the terminal window width
	ViewportWidthUnit
	// ViewportHeight is a percentage based on the terminal height width
	ViewportHeightUnit

	NotDefined = iota
)
