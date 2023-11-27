package mouse

// Copied from tcell

// ButtonMask is a mask of mouse buttons and wheel events. Mouse button presses
// are normally delivered as both press and release events. Mouse wheel events
// are normally just single impulse events. Windows supports up to eight
// separate buttons, but XTerm can only support mouse buttons 1-3.
// Its not unheard of for terminals to support only one or two buttons
// (think Macs). Old terminals, and true emulations (such as vt100) won't support
// mice at all, of course.
type ButtonMask int16

// These are the actual button values. Note that tcell version 1.x reversed buttons
// two and three on *nix based terminals. We use button 1 as the primary, and
// button 2 as the secondary, and button 3 (which is often missing) as the middle.
const (
	Button1 ButtonMask = 1 << iota // Usually the left (primary) mouse button.
	Button2                        // Usually the right (secondary) mouse button.
	Button3                        // Usually the middle mouse button.
	Button4                        // Often a side button (thumb/next).
	Button5                        // Often a side button (thumb/prev).
	Button6
	Button7
	Button8
	ButtonNone ButtonMask = 0 // No button.

	ButtonPrimary   = Button1
	ButtonSecondary = Button2
	ButtonMiddle    = Button3
)
