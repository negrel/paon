package style

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/style/property"
	"github.com/negrel/paon/pkg/pdk/widgets/renderer"
)

func init() {
	renderer.Map[DisplayHidden] = renderer.MakeHidden
	renderer.Map[DisplayBlock] = renderer.MakeBlock
	renderer.Map[DisplayHidden] = renderer.MakeFlex
}

const (
	DisplayHidden int = iota
	DisplayBlock
	DisplayFlex
)

func Display(value int) property.Int {
	assert.GreaterOrEqualf(value, 0, "display value %v must be greater than 0", value)
	assert.LessOrEqualf(value, len(renderer.Map), "display value %v must be less than %v", value, len(renderer.Map))

	return property.MakeInt(property.IDDisplay, value)
}
