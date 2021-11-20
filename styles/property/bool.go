package property

import (
	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/id/store"
)

// BoolID define a unique ID accross all Bool properties.
type BoolID id.ID

var (
	boolRegistry = id.Registry{}
	boolMap      = store.NewStringMap()
)

// NewBoolID returns a new unique Bool property ID.
func NewBoolID(name string) BoolID {
	id := boolRegistry.New()
	boolMap.Set(id, name)

	return BoolID(id)
}

// String implements the fmt.Stringer interface.
func (bi BoolID) String() string {
	return boolMap.Get(id.ID(bi))
}

// BoolIDCount returns the number of BoolID generated.
func BoolIDCount() uint32 {
	return uint32(boolRegistry.Last())
}

// Bool defines a boolean.
type Bool bool

// NewBool returns a new Bool with the given value.
func NewBool(value bool) Bool {
	return Bool(value)
}

// Value returns the bool value.
func (b *Bool) Value() bool {
	return bool(*b)
}

var _true = NewBool(true)
var _false = NewBool(false)
var _inherit = NewBool(false)

// True returns a preallocated true Bool.
func True() *Bool {
	return &_true
}

// False returns a preallocated false Bool.
func False() *Bool {
	return &_false
}

// Inherit returns the inherit Bool.
func Inherit() *Bool {
	return &_inherit
}

var (
	_IDBold          = NewBoolID("bold")
	_IDBlink         = NewBoolID("blink")
	_IDReverse       = NewBoolID("reverse")
	_IDUnderline     = NewBoolID("underline")
	_IDim            = NewBoolID("dim")
	_IDItalic        = NewBoolID("italic")
	_IDStrikeThrough = NewBoolID("strike-through")
)

// Bold returns the BoolID of the "bold" property.
func Bold() BoolID {
	return _IDBold
}

// Blink returns the BoolID of the "blink" property.
func Blink() BoolID {
	return _IDBlink
}

// Reverse returns the BoolID of the "reverse" property.
func Reverse() BoolID {
	return _IDReverse
}

// Underline returns the BoolID of the "underline" property.
func Underline() BoolID {
	return _IDUnderline
}

// Dim returns the BoolID of the "dim" property.
func Dim() BoolID {
	return _IDim
}

// Italic returns the BoolID of the "italic" property.
func Italic() BoolID {
	return _IDItalic
}

// StrikeThrough returns the BoolID of the "strike-through" property.
func StrikeThrough() BoolID {
	return _IDStrikeThrough
}
