package style

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/layout"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

const (
	DisplayHidden = layout.Hidden
	DisplayBlock  = layout.Block
	DisplayFlex   = layout.Flex
)

func Display(value int) property.Int {
	assert.True(layout.IsValidAlgorithmID(value))

	return property.MakeInt(property.IDDisplay, value)
}
