package styles

import "github.com/negrel/paon/pdk/events"

type Option func(*style)

// EventTarget returns an Option that sets the weight of a Style.
func Weight(weight int) Option {
	return func(s *style) {
		s.weight = weight
	}
}

// EventTarget returns an Option that sets the event target used by a Style.
func EventTarget(target events.Target) Option {
	return func(s *style) {
		s.Target = target
	}
}
