package flows

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// AddCustomRenderer adds the given draw.Painter constructor to the render map
// and return its ID. This ID should be used for the Display property.
func AddCustomAlgorithm(constructor func() Flow) int {
	return fMap.add(constructor)
}

// GetFor returns a new draw.Painter for the given draw.Object.
func GetFor(object styles.Stylised) Flow {
	prop := object.Style().Get(property.IDFlow)
	assert.NotNil(prop)
	algoID := prop.(property.Int)

	makeFlow, ok := fMap.algorithms[algoID.Value]
	if !ok {
		return fMap.algorithms[Unset]()
	}

	return makeFlow()
}

// IsValidFlowID returns true if the given rendererID is valid.
func IsValidFlowID(algorithmID int) bool {
	return algorithmID > 0 && algorithmID < fMap.len()
}

type flowMap struct {
	algorithms map[int]func() Flow
}

const (
	Unset = iota
	Hidden
	Block
	Inline
	Flex
)

var fMap = flowMap{
	algorithms: map[int]func() Flow{
		Unset:  func() Flow { return nil },
		Hidden: makeHidden,
		Block:  makeBlock,
		Inline: makeInline,
		Flex:   makeFlex,
	},
}

func (fm flowMap) add(renderer func() Flow) int {
	index := fm.len()
	fm.algorithms[index] = renderer

	return index
}

func (fm flowMap) len() int {
	return len(fm.algorithms)
}
