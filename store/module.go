package store

// Module is contain it's own state
type Module struct {
	s       State
	actions []func(s State){}
}

// Commit an action
func (m *Module) Commit(action string) {
	m.actions[action](s)
}
