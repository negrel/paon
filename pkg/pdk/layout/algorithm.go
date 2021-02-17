package layout

import "github.com/negrel/paon/pkg/pdk/styles"

// Algorithm is a generic interface for layout algorithm.
type Algorithm interface {
	Apply(styles.Stylised, Constraint) Box
}

var _ Algorithm = AlgoFunc(nil)

// AlgoFunc is a function that implements the Algorithm interface.
type AlgoFunc func(styles.Stylised, Constraint) Box

// Apply implements the Algorithm interface.
func (af AlgoFunc) Apply(s styles.Stylised, c Constraint) Box {
	return af(s, c)
}
