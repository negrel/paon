package color

// Color is an alias for int32. It should be use
// to write hexadecimal color.
// Red: 0xFF0000
// Green: 0x00FF00
// Blue: 0x0000FF
type Color = int32

// Default is used to leave the Color unchanged
// from whatever system or teminal default may exist.
var Default Color = -1
