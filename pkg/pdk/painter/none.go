package painter

import "github.com/negrel/paon/internal/draw"

var _ draw.Painter = none{}

type none struct{}

func makeHidden() draw.Painter {
	return none{}
}

func (n none) Layout(ctx draw.Context) {}

func (n none) Draw(ctx draw.Context) {}
