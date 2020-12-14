package render

// Object define any object that can be rendered
type Object interface {
	ParentObject() Object

	Layout(ctx Context)
	Draw(ctx Context)

	Steps() []string
	Step(name string) Step
	SetStep(Step)

	RenderStep() StepType
}
