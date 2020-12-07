package render

type StepType int

const (
	LayoutStepType StepType = iota
	DrawStepType
	EndStepType
)

type Step interface {
	Name() string
	Type() StepType
}

type LayoutStep interface {
	Step

	Layout(ctx Context)
}

type DrawStep interface {
	Step

	Draw(ctx Context)
}
