package flows

import "github.com/negrel/paon/internal/geometry"

// Constraint define the constraint an Flow must respect.
type Constraint struct {
	Min, Max geometry.Rectangle
}
