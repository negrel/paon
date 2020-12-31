package runtime

import "github.com/negrel/paon/internal/draw"

// Window define the terminal window.
var Window, _ = draw.NewTcellScreen()
