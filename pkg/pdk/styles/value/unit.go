package value

type UnitID uint8

const (
	CellUnit UnitID = iota
	PercentageUnit
	WindowWidthUnit
	WindowHeightUnit
)

type Unit struct {
	Value int
	ID    UnitID
}
