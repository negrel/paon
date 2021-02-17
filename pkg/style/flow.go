package style

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/flows"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

const (
	FlowHidden = flows.Hidden
	FlowBlock  = flows.Block
	FlowFlex   = flows.Flex
)

func Flow(value int) property.Int {
	assert.True(flows.IsValidFlowID(value))

	return property.MakeInt(property.IDFlow, value)
}
