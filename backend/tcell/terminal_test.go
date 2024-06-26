package tcell

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/golang/mock/gomock"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/stretchr/testify/require"
)

func TestTerminalHappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	screenMock := NewMockScreen(ctrl)
	screenMock.EXPECT().Init().Return(nil).Times(1)

	ter, err := NewTerminal(Screen(screenMock))
	require.NoError(t, err)

	screenMock.EXPECT().EnableMouse(tcell.MouseMotionEvents).Times(1)
	screenMock.EXPECT().EnablePaste().Times(1)

	evch := make(chan events.Event)
	err = ter.Start(evch)
	require.NoError(t, err)

	screenMock.EXPECT().GetContent(1, 2).
		Return(' ', []rune{}, tcell.StyleDefault.Reverse(true), 1).
		Times(1)

	cell := ter.Get(geometry.Vec2D{X: 1, Y: 2})
	require.Equal(t, cell, draw.Cell{
		Style: draw.CellStyle{
			Foreground:    0,
			Background:    0,
			Bold:          false,
			Blink:         false,
			Reverse:       true,
			Underline:     false,
			Dim:           false,
			Italic:        false,
			StrikeThrough: false,
		},
		Content: ' ',
	})

	screenMock.EXPECT().SetContent(3, 1, 'a', []rune{}, tcell.StyleDefault.Bold(true)).
		Times(1)
	ter.Set(geometry.Vec2D{X: 3, Y: 1}, draw.Cell{
		Style: draw.CellStyle{
			Foreground:    0,
			Background:    0,
			Bold:          true,
			Blink:         false,
			Reverse:       false,
			Underline:     false,
			Dim:           false,
			Italic:        false,
			StrikeThrough: false,
		},
		Content: 'a',
	})

	screenMock.EXPECT().Show().Times(1)
	ter.Flush()

	screenMock.EXPECT().Clear().Times(1)
	ter.Clear()

	screenMock.EXPECT().Fini().Times(1)
	ter.Stop()
}
