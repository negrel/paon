package property

type Int struct {
	Prop
	Value int
}

func NewInt(id ID, value int) Int {
	return Int{
		Prop:  NewProp(id),
		Value: value,
	}
}
