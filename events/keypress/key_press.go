package keypress

import (
	"fmt"

	"github.com/negrel/paon/events"
)

var EventType = events.NewType("KeyPress")

// Listener returns an events.Listener that will call the given handler
// on key press events.
func Listener(handler func(Event)) (events.Type, events.Handler) {
	return EventType, events.HandlerFunc(func(event events.Event) {
		handler(event.(Event))
	})
}

// Event define a key press event.
type Event struct {
	events.Event

	// Modifiers is the modifiers that were present with the key press. Note
	// that not all platforms and terminals support this equally well, and some
	// cases we will not not know for sure. Hence, applications should avoid
	// using this in most circumstances.
	Modifiers ModMask

	// Key is a virtual key code. We use this to identify specific key
	// codes, such as KeyEnter, etc. Most control and function keys are reported
	// with unique Key values. Normal alphanumeric and punctuation keys will
	// generally return KeyRune here; the specific key can be further decoded
	// using the Rune field.
	Key Key

	// Rune corresponding to the key press, if it makes sense.
	// The result is only defined if the value of Key is KeyRune.
	Rune rune
}

// New returns a new key press event.
func New(modifiers ModMask, key Key, r rune) Event {
	return Event{
		Event:     events.NewEvent(EventType),
		Modifiers: modifiers,
		Key:       key,
		Rune:      r,
	}
}

// String implements the fmt.Stringer interface.
func (kp Event) String() string {
	return fmt.Sprintf("%v{Modifiers: %v, Key: %v, Rune: %v}", kp.Type(), kp.Modifiers, kp.Key, string(kp.Rune))
}

// String implements the fmt.Stringer interface.
func (k Key) String() string {
	if keyName, hasKeyName := KeyNames[k]; hasKeyName {
		return keyName
	}

	return ""
}

// Copied from tcell:

// KeyNames holds the written names of special keys. Useful to echo back a key
// name, or to look up a key from a string value.
var KeyNames = map[Key]string{
	KeyEnter:          "Enter",
	KeyBackspace:      "Backspace",
	KeyTab:            "Tab",
	KeyBacktab:        "Backtab",
	KeyEsc:            "Esc",
	KeyBackspace2:     "Backspace2",
	KeyDelete:         "Delete",
	KeyInsert:         "Insert",
	KeyUp:             "Up",
	KeyDown:           "Down",
	KeyLeft:           "Left",
	KeyRight:          "Right",
	KeyHome:           "Home",
	KeyEnd:            "End",
	KeyUpLeft:         "UpLeft",
	KeyUpRight:        "UpRight",
	KeyDownLeft:       "DownLeft",
	KeyDownRight:      "DownRight",
	KeyCenter:         "Center",
	KeyPgDn:           "PgDn",
	KeyPgUp:           "PgUp",
	KeyClear:          "Clear",
	KeyExit:           "Exit",
	KeyCancel:         "Cancel",
	KeyPause:          "Pause",
	KeyPrint:          "Print",
	KeyF1:             "F1",
	KeyF2:             "F2",
	KeyF3:             "F3",
	KeyF4:             "F4",
	KeyF5:             "F5",
	KeyF6:             "F6",
	KeyF7:             "F7",
	KeyF8:             "F8",
	KeyF9:             "F9",
	KeyF10:            "F10",
	KeyF11:            "F11",
	KeyF12:            "F12",
	KeyF13:            "F13",
	KeyF14:            "F14",
	KeyF15:            "F15",
	KeyF16:            "F16",
	KeyF17:            "F17",
	KeyF18:            "F18",
	KeyF19:            "F19",
	KeyF20:            "F20",
	KeyF21:            "F21",
	KeyF22:            "F22",
	KeyF23:            "F23",
	KeyF24:            "F24",
	KeyF25:            "F25",
	KeyF26:            "F26",
	KeyF27:            "F27",
	KeyF28:            "F28",
	KeyF29:            "F29",
	KeyF30:            "F30",
	KeyF31:            "F31",
	KeyF32:            "F32",
	KeyF33:            "F33",
	KeyF34:            "F34",
	KeyF35:            "F35",
	KeyF36:            "F36",
	KeyF37:            "F37",
	KeyF38:            "F38",
	KeyF39:            "F39",
	KeyF40:            "F40",
	KeyF41:            "F41",
	KeyF42:            "F42",
	KeyF43:            "F43",
	KeyF44:            "F44",
	KeyF45:            "F45",
	KeyF46:            "F46",
	KeyF47:            "F47",
	KeyF48:            "F48",
	KeyF49:            "F49",
	KeyF50:            "F50",
	KeyF51:            "F51",
	KeyF52:            "F52",
	KeyF53:            "F53",
	KeyF54:            "F54",
	KeyF55:            "F55",
	KeyF56:            "F56",
	KeyF57:            "F57",
	KeyF58:            "F58",
	KeyF59:            "F59",
	KeyF60:            "F60",
	KeyF61:            "F61",
	KeyF62:            "F62",
	KeyF63:            "F63",
	KeyF64:            "F64",
	KeyCtrlA:          "Ctrl-A",
	KeyCtrlB:          "Ctrl-B",
	KeyCtrlC:          "Ctrl-C",
	KeyCtrlD:          "Ctrl-D",
	KeyCtrlE:          "Ctrl-E",
	KeyCtrlF:          "Ctrl-F",
	KeyCtrlG:          "Ctrl-G",
	KeyCtrlJ:          "Ctrl-J",
	KeyCtrlK:          "Ctrl-K",
	KeyCtrlL:          "Ctrl-L",
	KeyCtrlN:          "Ctrl-N",
	KeyCtrlO:          "Ctrl-O",
	KeyCtrlP:          "Ctrl-P",
	KeyCtrlQ:          "Ctrl-Q",
	KeyCtrlR:          "Ctrl-R",
	KeyCtrlS:          "Ctrl-S",
	KeyCtrlT:          "Ctrl-T",
	KeyCtrlU:          "Ctrl-U",
	KeyCtrlV:          "Ctrl-V",
	KeyCtrlW:          "Ctrl-W",
	KeyCtrlX:          "Ctrl-X",
	KeyCtrlY:          "Ctrl-Y",
	KeyCtrlZ:          "Ctrl-Z",
	KeyCtrlSpace:      "Ctrl-Space",
	KeyCtrlUnderscore: "Ctrl-_",
	KeyCtrlRightSq:    "Ctrl-]",
	KeyCtrlBackslash:  "Ctrl-\\",
	KeyCtrlCarat:      "Ctrl-^",
}

// ModMask is a mask of modifier keys.  Note that it will not always be
// possible to report modifier keys.
type ModMask int16

// These are the modifiers keys that can be sent either with a key press,
// or a mouse event.  Note that as of now, due to the confusion associated
// with Meta, and the lack of support for it on many/most platforms, the
// current implementations never use it.  Instead, they use ModAlt, even for
// events that could possibly have been distinguished from ModAlt.
const (
	ModShift ModMask = 1 << iota
	ModCtrl
	ModAlt
	ModMeta
	ModNone ModMask = 0
)

// Key is a generic value for representing keys, and especially special
// keys (function keys, cursor movement keys, etc.)  For normal keys, like
// ASCII letters, we use KeyRune, and then expect the application to
// inspect the Rune() member of the EventKey.
type Key int16

// This is the list of named keys.  KeyRune is special however, in that it is
// a place holder key indicating that a printable character was sent.  The
// actual value of the rune will be transported in the Rune of the associated
// EventKey.
const (
	KeyRune Key = iota + 256
	KeyUp
	KeyDown
	KeyRight
	KeyLeft
	KeyUpLeft
	KeyUpRight
	KeyDownLeft
	KeyDownRight
	KeyCenter
	KeyPgUp
	KeyPgDn
	KeyHome
	KeyEnd
	KeyInsert
	KeyDelete
	KeyHelp
	KeyExit
	KeyClear
	KeyCancel
	KeyPrint
	KeyPause
	KeyBacktab
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyF25
	KeyF26
	KeyF27
	KeyF28
	KeyF29
	KeyF30
	KeyF31
	KeyF32
	KeyF33
	KeyF34
	KeyF35
	KeyF36
	KeyF37
	KeyF38
	KeyF39
	KeyF40
	KeyF41
	KeyF42
	KeyF43
	KeyF44
	KeyF45
	KeyF46
	KeyF47
	KeyF48
	KeyF49
	KeyF50
	KeyF51
	KeyF52
	KeyF53
	KeyF54
	KeyF55
	KeyF56
	KeyF57
	KeyF58
	KeyF59
	KeyF60
	KeyF61
	KeyF62
	KeyF63
	KeyF64
)

const (
	// These key codes are used internally, and will never appear to applications.
	keyPasteStart Key = iota + 16384
	keyPasteEnd
)

// These are the control keys.  Note that they overlap with other keys,
// perhaps.  For example, KeyCtrlH is the same as KeyBackspace.
const (
	KeyCtrlSpace Key = iota
	KeyCtrlA
	KeyCtrlB
	KeyCtrlC
	KeyCtrlD
	KeyCtrlE
	KeyCtrlF
	KeyCtrlG
	KeyCtrlH
	KeyCtrlI
	KeyCtrlJ
	KeyCtrlK
	KeyCtrlL
	KeyCtrlM
	KeyCtrlN
	KeyCtrlO
	KeyCtrlP
	KeyCtrlQ
	KeyCtrlR
	KeyCtrlS
	KeyCtrlT
	KeyCtrlU
	KeyCtrlV
	KeyCtrlW
	KeyCtrlX
	KeyCtrlY
	KeyCtrlZ
	KeyCtrlLeftSq // Escape
	KeyCtrlBackslash
	KeyCtrlRightSq
	KeyCtrlCarat
	KeyCtrlUnderscore
)

// Special values - these are fixed in an attempt to make it more likely
// that aliases will encode the same way.

// These are the defined ASCII values for key codes.  They generally match
// with KeyCtrl values.
const (
	KeyNUL Key = iota
	KeySOH
	KeySTX
	KeyETX
	KeyEOT
	KeyENQ
	KeyACK
	KeyBEL
	KeyBS
	KeyTAB
	KeyLF
	KeyVT
	KeyFF
	KeyCR
	KeySO
	KeySI
	KeyDLE
	KeyDC1
	KeyDC2
	KeyDC3
	KeyDC4
	KeyNAK
	KeySYN
	KeyETB
	KeyCAN
	KeyEM
	KeySUB
	KeyESC
	KeyFS
	KeyGS
	KeyRS
	KeyUS
	KeyDEL Key = 0x7F
)

// These keys are aliases for other names.
const (
	KeyBackspace  = KeyBS
	KeyTab        = KeyTAB
	KeyEsc        = KeyESC
	KeyEscape     = KeyESC
	KeyEnter      = KeyCR
	KeyBackspace2 = KeyDEL
)
