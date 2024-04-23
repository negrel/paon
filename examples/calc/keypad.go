package main

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

type KeyPad struct {
	widgets.BaseLayout
}

func NewKeyPad(m *Model) *KeyPad {
	keyPad := &KeyPad{}
	keyPad.BaseLayout = widgets.NewBaseLayout(keyPad)

	keyPad.AppendChild(NewNumPad(m))
	keyPad.AppendChild(NewOpPad(m))

	return keyPad
}

// Layout implements layout.Layout.
func (kp *KeyPad) Layout(co layout.Constraint) geometry.Size {
	kp.BaseLayout.ChildrenLayout.Reset()

	numPad := kp.FirstChild()
	numPadCo := co.ForceSize(co.MaxSize.WithWidth(co.MaxSize.Width / 5 * 4))

	kp.BaseLayout.ChildrenLayout.Append(widgets.ChildLayout{
		Widget: numPad,
		Bounds: geometry.Rectangle{
			Origin:   geometry.Vec2D{},
			RectSize: numPadCo.ApplyOnSize(numPad.Layout(numPadCo)),
		},
	})

	opPad := kp.LastChild()
	opPadCo := co.ForceSize(co.MaxSize.WithWidth(co.MaxSize.Width - co.MaxSize.Width/5*4))

	kp.BaseLayout.ChildrenLayout.Append(widgets.ChildLayout{
		Widget: opPad,
		Bounds: geometry.Rectangle{
			Origin:   geometry.Vec2D{X: co.MaxSize.Width / 5 * 4, Y: 0},
			RectSize: numPadCo.ApplyOnSize(opPad.Layout(opPadCo)),
		},
	})

	return co.MaxSize
}
