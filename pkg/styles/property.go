package styles

import "github.com/negrel/paon/internal/render"

const (
	DisplayPriority = iota * 100

	WidthPriority
	HeightPriority
	FlexPriority
	FlexDirectionPriority
	FlexWrapPriority
	MarginPriority
	BorderPriority

	PaddingPriority
	BackgroundColorPriority

	innerRenderingPriority

	TextColorPriority
	TextBoldPriority
	TextUnderlinePriority
	TextBlinkPriority

	OverflowPriority
)

type Property interface {
	render.Step

	Priority() int
}

var _ Property = &property{}

type property struct {
	step     render.StepType
	name     string
	priority int
}

func (p property) Name() string {
	return p.name
}

func (p property) Type() render.StepType {
	return p.step
}

func (p property) Priority() int {
	return p.priority
}
