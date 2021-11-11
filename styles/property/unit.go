package property

type Unit uint8

const (
	CellUnit Unit = iota
	PercentageWidthUnit
	PercentageHeightUnit
	WindowWidthUnit
	WindowHeightUnit
)
