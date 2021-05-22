package value

type UnitID uint8

const (
	CellUnit UnitID = iota
	PercentageWidthUnit
	PercentageHeightUnit
	WindowWidthUnit
	WindowHeightUnit
)

type Unit struct {
	Value int
	UnitID
}
