package property

// Unit define units of measure that can be used.
type Unit uint8

// Unit that can be used for IntUnit properties.
const (
	CellUnit Unit = iota
	PercentageWidthUnit
	PercentageHeightUnit
	WindowWidthUnit
	WindowHeightUnit
)
