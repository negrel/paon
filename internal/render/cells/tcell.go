package cells

import (
	"github.com/gdamore/tcell"
)

var _ Cell = &tcellCell{}

type tcellCell struct {
	style   *tcell.Style
	content rune
}

// Style implements the Cell interface.
func (tc tcellCell) Style() tcell.Style {
	return *tc.style
}

// Content implements the Cell interface.
func (tc tcellCell) Content() rune {
	return tc.content
}
