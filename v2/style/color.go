package style

// Color represents a color for a terminal cell
// (foreground or background).
type Color = int32

// DefaultColor leave the color unchanged from whatever
// your terminal default color is.
var DefaultColor Color = -1
