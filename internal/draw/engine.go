package draw

type Engine struct {
	set map[Object]struct{}
}

func (e Engine) clear() {
	e.set = make(map[Object]struct{})
}

// Add the given object to the list of object to render for the next frame.
func (e Engine) Add(obj Object) {
	e.set[obj] = struct{}{}
}
