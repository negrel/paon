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
	assert.True(renderer.IsValidRendererID(value))

	return property.MakeInt(property.IDDisplay, value)
}
