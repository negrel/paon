package styles

type NumericValue struct {
	Defined bool
	value   int
	unit    Unit
}

func NValue(value int, unit Unit) NumericValue {
	return NumericValue{
		Defined: true,
		value:   value,
		unit:    unit,
	}
}

func (nv NumericValue) toCellUnit() NumericValue {
	return NumericValue{
		Defined: nv.Defined,
		value:   nv.value,
		unit:    CellUnit,
	}
}
