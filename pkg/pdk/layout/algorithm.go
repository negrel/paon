package layout

import "github.com/negrel/paon/pkg/pdk/styles"

// Algorithm is a generic interface for layout algorithm.
type Algorithm interface {
	Apply(styles.Stylised, Constraint) Result
}
