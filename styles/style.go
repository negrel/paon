package styles

import (
	"github.com/negrel/paon/draw"
)

// Styled define any object with a Style.
type Styled interface {
	Style() *Style
}

// Style holds styling properties.
type Style struct {
	draw.CellStyle
	Extras map[string]any
}
