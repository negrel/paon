package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/negrel/paon"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/minmax"
	"github.com/negrel/paon/render"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/button"
	"github.com/negrel/paon/widgets/span"
	"github.com/negrel/paon/widgets/vbox"
)

func main() {
	// Parse args.
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "only one file argument expected")
		return
	}
	fpath := args[1]

	rawContent, err := os.ReadFile(fpath)
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "failed to readt file: %v", err.Error())
		return
	}
	lines := strings.Split(string(rawContent), "\n")

	// Create application.
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start application.
	err = app.Start(ctx, widgets.NewRoot(
		vbox.New(
			vbox.WithChildren(
				button.New("Click to exit", button.OnClick(func(_ mouse.ClickEvent) {
					cancel()
				})),
				NewPager(lines),
			),
		),
	))
	if err != nil {
		panic(err)
	}
}

type Pager struct {
	*widgets.PanicWidget
	pagerRenderable
}

func NewPager(lines []string) *Pager {
	p := &Pager{}

	p.PanicWidget = widgets.NewPanicWidget(p)
	p.pagerRenderable = pagerRenderable{
		VoidRenderable: render.NewVoidRenderable(p),
		lines:          lines,
	}

	lastScrollDirection := mouse.ScrollDown
	p.AddEventListener(mouse.ScrollListener(func(ev mouse.ScrollEvent) {
		if ev.Vertical() {
			if lastScrollDirection != ev.ScrollDirection {
				p.pagerRenderable.scrollFrame = 0
			}
			lastScrollDirection = ev.ScrollDirection

			if ev.ScrollDirection == mouse.ScrollUp {
				p.pagerRenderable.scrollFrame -= 4
			} else {
				p.pagerRenderable.scrollFrame += 4
			}
			p.pagerRenderable.MarkDirty()
		}
	}))

	return p
}

// Renderable implements render.Renderable.
func (p *Pager) Renderable() render.Renderable {
	return &p.pagerRenderable
}

type pagerRenderable struct {
	render.VoidRenderable
	lines        []string
	scrollOffset int
	scrollFrame  int
}

func (pr *pagerRenderable) addScrollOffset(offset int) {
	pr.scrollOffset += offset
	pr.scrollOffset = minmax.Constrain(pr.scrollOffset, 0, len(pr.lines)-1)
}

// Layout implements layout.Layout.
func (pr *pagerRenderable) Layout(co layout.Constraint) geometry.Size {
	return co.MaxSize
}

// Draw implements draw.Drawer.
func (pr *pagerRenderable) Draw(surface draw.Surface) {
	pr.VoidRenderable.Draw(surface)
	scrollSpeed := 3
	if pr.scrollFrame != 0 {
		offset := pr.scrollFrame / scrollSpeed
		pr.addScrollOffset(offset)
		pr.scrollFrame -= offset
		if (pr.scrollFrame < 0 && pr.scrollFrame > -scrollSpeed) ||
			(pr.scrollFrame > 0 && pr.scrollFrame < scrollSpeed) {
			pr.scrollFrame = 0
		}
		pr.MarkDirty()
	}

	totalLine := minmax.Min(surface.Size().Height, len(pr.lines[pr.scrollOffset:]))
	for i := 0; i < totalLine; i++ {
		subsurface := draw.NewSubSurface(surface, geometry.Rectangle{Origin: geometry.Vec2D{X: 0, Y: i}, RectSize: surface.Size()})
		span.Draw(subsurface, pr.lines[pr.scrollOffset+i], draw.CellStyle{})
	}
}
