package paon

import (
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/pkg/pdk/events"
)

type Window interface {
	geometry.Sized
	events.Target
}

var _ Window = &window{}

type window struct {
	events.Target
	render.Screen
}

func newWindow() *window {
	return &window{
		Target: events.MakeTarget(),
		Screen: nil,
	}
}
