package id

type Identifiable interface {
	ID() ID
	IsSame(other Identifiable) bool
}
