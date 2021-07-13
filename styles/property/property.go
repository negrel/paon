package property

// Property define a generic interface for styling properties.
type Property interface {
	ID() ID
}

// Prop define an empty proprety with no value.
// Prop should be used in composite struct for custom properties.
type Prop struct {
	id ID
}

// NewProp returns a new Prop object with the given property ID.
func NewProp(id ID) Prop {
	return Prop{
		id: id,
	}
}

// ID returns the property ID of this prop.
func (p Prop) ID() ID {
	return p.id
}
