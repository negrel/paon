package styles

import "github.com/negrel/paon/pkg/runtime"

type Unit uint8

const (
	CellUnit Unit = iota
	PercentageUnit
	WindowWidthUnit
	WindowHeightUnit
)

type UnitValue struct {
	Value int
	Unit  Unit
}

func (uv UnitValue) CellUnit() UnitValue {
	switch uv.Unit {
	case CellUnit:
		return uv

	case PercentageUnit:

	case WindowWidthUnit:
		return UnitValue{
			Value: runtime.Window.Width() / 100 * uv.Value,
			Unit:  CellUnit,
		}
	case WindowHeightUnit:
		return UnitValue{
			Value: runtime.Window.Height() / 100 * uv.Value,
			Unit:  CellUnit,
		}

	default:
		panic("can't convert unknown unit value to cell unit")
	}

	return UnitValue{}
}
