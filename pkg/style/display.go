package style

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/painter"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

const (
	DisplayHidden = painter.Hidden
	DisplayBlock  = painter.Block
	DisplayFlex   = painter.Flex
)

func Display(value int) property.Int {
	assert.True(painter.IsValidRendererID(value))

	return property.MakeInt(property.IDDisplay, value)
}
