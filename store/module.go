package store

// Module is contain it's own state
type Module struct {
	s        State
	mutators map[string]func(s State)
	getters  map[string]func(s State) interface{}
}

// Commit an action
func (m *Module) Commit(action string) {
	m.mutators[action](m.s)
}
