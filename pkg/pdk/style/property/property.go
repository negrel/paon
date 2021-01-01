package property

type Property interface {
	ID() ID
	IsInherited() bool
}

type Prop struct {
	id        ID
	inherited bool
}

func MakeProp(id ID) Prop {
	return Prop{
		id:        id,
		inherited: false,
	}
}

func (p Prop) ID() ID {
	return p.id
}

func (p Prop) IsInherited() bool {
	return p.inherited
}
