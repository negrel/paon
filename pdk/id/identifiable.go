package id

// Identifiable define any object containing a unique ID.
type Identifiable interface {
	ID() ID
	IsSame(other Identifiable) bool
}
