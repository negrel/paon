package layout

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// AddCustomRenderer adds the given draw.Painter constructor to the render map
// and return its ID. This ID should be used for the Display property.
func AddCustomAlgorithm(constructor func() Algorithm) int {
	return algoMap.add(constructor)
}

// GetFor returns a new draw.Painter for the given draw.Object.
func GetFor(object styles.Stylised) Algorithm {
	prop := object.Style().Get(property.IDDisplay)
	assert.NotNil(prop)
	algoID := prop.(property.Int)

	return algoMap.algorithms[algoID.Value]()
}

type algorithmMap struct {
	algorithms map[int]func() Algorithm
}

const (
	Hidden = iota
	Block
	Inline
	Flex
)

var algoMap = algorithmMap{
	algorithms: map[int]func() Algorithm{
		Hidden: makeHidden,
		Block:  makeBlock,
		Inline: makeInline,
		Flex:   makeFlex,
	},
}

func (rm algorithmMap) add(renderer func() Algorithm) int {
	index := rm.len()
	rm.algorithms[index] = renderer

	return index
}

func (rm algorithmMap) len() int {
	return len(rm.algorithms)
}
