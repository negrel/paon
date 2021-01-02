package property

type Int struct {
	Prop
	Value int
}

func MakeInt(id ID, value int) Int {
	return Int{
		Prop:  MakeProp(id),
		Value: value,
	}
}
