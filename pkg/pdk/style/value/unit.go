package value

import "github.com/negrel/paon/pkg/runtime"

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

func (uv Unit) CellUnit() Unit {
	switch uv.ID {
	case CellUnit:
		return uv

	case PercentageUnit:

	case WindowWidthUnit:
		return Unit{
			Value: runtime.Window.Width() / 100 * uv.Value,
			ID:    CellUnit,
		}
	case WindowHeightUnit:
		return Unit{
			Value: runtime.Window.Height() / 100 * uv.Value,
			ID:    CellUnit,
		}

	default:
		panic("can't convert unknown unit value to cell unit")
	}

	return Unit{}
}
