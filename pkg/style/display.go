package style

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/renderer"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

const (
	DisplayHidden = renderer.Hidden
	DisplayBlock  = renderer.Block
	DisplayFlex   = renderer.Flex
)

func Display(value int) property.Int {
	assert.GreaterOrEqualf(value, 0, "display value %v must be greater than 0", value)
	assert.LessOrEqualf(value, renderer.Map.Len(), "display value %v must be less than %v", value, renderer.Map.Len())

	return property.MakeInt(property.IDDisplay, value)
}
