package widgets

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/pkg/styles"
)

type Option func(widget Widget)

func Listener(eventType events.EventType, listener events.Listener) Option {
	return func(widget Widget) {
		widget.AddEventListener(eventType, &listener)
	}
}

func InnerDrawStep(step render.DrawStep) Option {
	return func(widget Widget) {
		widget.setDrawStep(step)
	}
}

var _ styles.Property = renderFunc{}

type renderFunc struct {
	name     string
	priority int
	fn       func(ctx render.Context)
}

func (rf renderFunc) Priority() int {
	return styles.InnerRenderingPriority
}

func (rf renderFunc) Name() string {
	return rf.name
}

func (rf renderFunc) Type() render.StepType {
	return render.DrawStepType
}

func (rf renderFunc) Draw(ctx render.Context) {
	rf.fn(ctx)
}

func InnerDrawFunc(fn func(ctx render.Context)) Option {
	return func(widget Widget) {
		rs := renderFunc{
			fn:   fn,
			name: "render-" + widget.String()}

		widget.setDrawStep(rs)
	}
}

func RenderSteps(steps ...render.Step) Option {
	return func(widget Widget) {
		for _, step := range steps {
			widget.SetStep(step)
		}
	}
}

//
//func Rendering(steps ...RenderStep) Option {
//	return func(widget Widget) {
//		for _, step := range steps {
//			widget.
//		}
//	}
//}
