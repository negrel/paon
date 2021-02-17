package flows

import "github.com/negrel/paon/pkg/pdk/styles"

// Flow is a generic interface for flow algorithm.
type Flow interface {
	Apply(styles.Stylised, Constraint) Box
}

var _ Flow = FlowFunc(nil)

// FlowFunc is a function that implements the Flow interface.
type FlowFunc func(styles.Stylised, Constraint) Box

// Apply implements the Flow interface.
func (af FlowFunc) Apply(s styles.Stylised, c Constraint) Box {
	return af(s, c)
}
