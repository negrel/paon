package store

// Store is an application store
type Store struct {
	*Module
	modules map[string]*Module
}

// Get return the named module
func (s *Store) Get(m string) *Module {
	return s.modules[m]
}
