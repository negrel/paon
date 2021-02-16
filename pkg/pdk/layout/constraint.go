package layout

import "github.com/negrel/paon/internal/geometry"

// Constraint define the constraint an Algorithm must respect.
type Constraint struct {
	Min, Max geometry.Rectangle
}
