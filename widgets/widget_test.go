package widgets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type panicsWrapper struct {
	*PanicWidget
}

func TestPanicsWidget(t *testing.T) {
	t.Run("Swap", func(t *testing.T) {
		widget1 := panicsWrapper{}
		widget1.PanicWidget = NewPanicWidget(&widget1)

		widget2 := panicsWrapper{}
		widget2.PanicWidget = NewPanicWidget(&widget2)

		preSwapWidget1Node := widget1.Node()
		preSwapWidget2Node := widget2.Node()

		require.NotSame(t, preSwapWidget1Node, preSwapWidget2Node)

		// Swap widgets.
		widget1.Swap(widget2)

		// Widgets node swapped.
		require.Same(t, preSwapWidget2Node, widget1.Node())
		require.Same(t, preSwapWidget1Node, widget2.Node())

		// Node underlying data was swapped also.
		require.Equal(t, &widget1, widget1.Node().Unwrap())
		require.Equal(t, &widget2, widget2.Node().Unwrap())
	})

	t.Run("Renderable/Unimplemented", func(t *testing.T) {
		widget := panicsWrapper{}
		widget.PanicWidget = NewPanicWidget(&widget)
		require.Panics(t, func() {
			widget.Renderable()
		})
	})

	t.Run("Style/Nil", func(t *testing.T) {
		widget := panicsWrapper{}
		widget.PanicWidget = NewPanicWidget(&widget)
		require.Nil(t, widget.Style())
	})
}
