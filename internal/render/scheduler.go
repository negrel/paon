package render

type Scheduler struct {
	set map[Object]struct{}
}

func (s Scheduler) clear() {
	s.set = make(map[Object]struct{})
}

// Add the given object to the list of object to render for the next frame.
func (s Scheduler) Add(obj Object) {
	s.set[obj] = struct{}{}
}
