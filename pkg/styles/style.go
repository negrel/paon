package styles

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/render"
	"sort"
)

var _ sort.Interface = &propList{}

type propList struct {
	props  map[string]Property
	sorted []string
}

func (p *propList) Len() int {
	return len(p.props)
}

func (p *propList) Less(i, j int) bool {
	return p.props[p.sorted[i]].Priority() < p.props[p.sorted[j]].Priority()
}

func (p *propList) Swap(i, j int) {
	p.sorted[i], p.sorted[j] = p.sorted[j], p.sorted[i]
}

type Style struct {
	*propList
	DrawStep render.DrawStep
}

func New() *Style {
	return &Style{
		propList: &propList{
			props:  make(map[string]Property, 8),
			sorted: make([]string, 0, 16),
		},
		DrawStep: nil,
	}
}

func (s *Style) Layout(ctx render.Context) {
	log.Debug("style", s)

	for _, prop := range s.props {
		if layoutStep, isLayoutStep := prop.(render.LayoutStep); isLayoutStep {
			layoutStep.Layout(ctx)
		}
	}
}

func (s *Style) Draw(ctx render.Context) {
	for _, prop := range s.props {
		if drawStep, isDrawStep := prop.(render.DrawStep); isDrawStep {
			drawStep.Draw(ctx)
		}
	}

	s.DrawStep.Draw(ctx)
	log.Debug("draw step done")
}

func (s *Style) Steps() []string {
	return s.propList.sorted
}

func (s *Style) Step(name string) render.Step {
	return s.props[name]
}

func (s *Style) SetStep(step render.Step) {
	if property, isProperty := step.(Property); isProperty {
		s.props[step.Name()] = property
		s.sorted = append(s.sorted, step.Name())
		sort.Sort(s.propList)
		return
	}

	s.DrawStep = step.(render.DrawStep)
}
