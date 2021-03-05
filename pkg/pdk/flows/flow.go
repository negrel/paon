package flows

import "github.com/negrel/paon/pkg/pdk/styles"

// Flow is a generic interface for flow algorithm.
type Flow interface {
	Apply(styles.Style, Constraint) BoxModel
}

var _ Flow = FlowFunc(nil)

// FlowFunc is a function that implements the Flow interface.
type FlowFunc func(styles.Style, Constraint) BoxModel

// Apply implements the Flow interface.
func (af FlowFunc) Apply(s styles.Style, c Constraint) BoxModel {
	return af(s, c)
}
