package store

import "github.com/negrel/paon/pdk/id"

// Uint define a generic interface for uint store.
type Uint interface {
	Set(id.ID, uint)
	Get(id.ID) uint
	Swap(id.ID, uint) uint
}

// UintSlice us.an implementation of Uint store.
type UintSlice []uint

// NewUintSlice returns a new uint store of the given size
func NewUintSlice(size int) UintSlice {
	return make(UintSlice, size)
}

// Set implements the Uint interface.
func (us UintSlice) Set(id id.ID, value uint) {
	us[id] = value
}

// Get implements the Uint interface.
func (us UintSlice) Get(id id.ID) uint {
	return us[id]
}

// Swap implements the Uint interface.
func (us UintSlice) Swap(id id.ID, value uint) uint {
	old := us.Get(id)
	us.Set(id, value)

	return old
}

// UintMap us.an implementation of Uint store.
type UintMap map[id.ID]uint

// NewUintMap returns a new UintMap store.
func NewUintMap() UintMap {
	return make(UintMap)
}

// Set implements the Uint interface.
func (im UintMap) Set(id id.ID, value uint) {
	im[id] = value
}

// Get implements the Uint interface.
func (im UintMap) Get(id id.ID) uint {
	return im[id]
}

// Swap implements the Uint interface.
func (im UintMap) Swap(id id.ID, value uint) uint {
	old := im.Get(id)
	im.Set(id, value)

	return old
}

// Uint8 define a generic interface for uint8 store.
type Uint8 interface {
	Set(id.ID, uint8)
	Get(id.ID) uint8
	Swap(id.ID, uint8) uint8
}

// Uint8Slice us.an implementation of Uint8 store.
type Uint8Slice []uint8

// NewUint8Slice returns a new Uint8Slice store of the given size
func NewUint8Slice(size int) Uint8Slice {
	return make(Uint8Slice, size)
}

// Set implements the Uint8 interface.
func (us Uint8Slice) Set(id id.ID, value uint8) {
	us[id] = value
}

// Get implements the Uint8 interface.
func (us Uint8Slice) Get(id id.ID) uint8 {
	return us[id]
}

// Swap implements the Uint8 interface.
func (us Uint8Slice) Swap(id id.ID, value uint8) uint8 {
	old := us.Get(id)
	us.Set(id, value)

	return old
}

// Uint8Map us.an implementation of Uint8 store.
type Uint8Map map[id.ID]uint8

// NewUint8Map returns a new Uint8Map store.
func NewUint8Map() Uint8Map {
	return make(Uint8Map)
}

// Set implements the Uint8 interface.
func (im Uint8Map) Set(id id.ID, value uint8) {
	im[id] = value
}

// Get implements the Uint8 interface.
func (im Uint8Map) Get(id id.ID) uint8 {
	return im[id]
}

// Swap implements the Uint8 interface.
func (im Uint8Map) Swap(id id.ID, value uint8) uint8 {
	old := im.Get(id)
	im.Set(id, value)

	return old
}

// Uint16 define a generic interface for uint16 store.
type Uint16 interface {
	Set(id.ID, uint16)
	Get(id.ID) uint16
	Swap(id.ID, uint16) uint16
}

// Uint16Slice us.an implementation of Uint16 store.
type Uint16Slice []uint16

// NewUint16Slice returns a new uint16 store of the given size
func NewUint16Slice(size int) Uint16Slice {
	return make(Uint16Slice, size)
}

// Set implements the Uint16 interface.
func (us Uint16Slice) Set(id id.ID, value uint16) {
	us[id] = value
}

// Get implements the Uint16 interface.
func (us Uint16Slice) Get(id id.ID) uint16 {
	return us[id]
}

// Swap implements the Uint16 interface.
func (us Uint16Slice) Swap(id id.ID, value uint16) uint16 {
	old := us.Get(id)
	us.Set(id, value)

	return old
}

// Uint16Map us.an implementation of Uint16 store.
type Uint16Map map[id.ID]uint16

// NewUint16Map returns a new Uint16Map store.
func NewUint16Map() Uint16Map {
	return make(Uint16Map)
}

// Set implements the Uint16 interface.
func (im Uint16Map) Set(id id.ID, value uint16) {
	im[id] = value
}

// Get implements the Uint16Map interface.
func (im Uint16Map) Get(id id.ID) uint16 {
	return im[id]
}

// Swap implements the Uint16 interface.
func (im Uint16Map) Swap(id id.ID, value uint16) uint16 {
	old := im.Get(id)
	im.Set(id, value)

	return old
}

// Uint32 define a generic interface for uint32 store.
type Uint32 interface {
	Set(id.ID, uint32)
	Get(id.ID) uint32
	Swap(id.ID, uint32) uint32
}

// Uint32Slice us.an implementation of Uint32 store.
type Uint32Slice []uint32

// NewUint32Slice returns a new Uint32Slice store of the given size
func NewUint32Slice(size int) Uint32Slice {
	return make(Uint32Slice, size)
}

// Set implements the Uint32 interface.
func (us Uint32Slice) Set(id id.ID, value uint32) {
	us[id] = value
}

// Get implements the Uint32 interface.
func (us Uint32Slice) Get(id id.ID) uint32 {
	return us[id]
}

// Swap implements the Uint32 interface.
func (us Uint32Slice) Swap(id id.ID, value uint32) uint32 {
	old := us.Get(id)
	us.Set(id, value)

	return old
}

// Uint32Map us.an implementation of Uint32 store.
type Uint32Map map[id.ID]uint32

// NewUint32Map returns a new Uint32Slice store.
func NewUint32Map() Uint32Map {
	return make(Uint32Map)
}

// Set implements the Uint32 interface.
func (im Uint32Map) Set(id id.ID, value uint32) {
	im[id] = value
}

// Get implements the Uint32 interface.
func (im Uint32Map) Get(id id.ID) uint32 {
	return im[id]
}

// Swap implements the Uint32 interface.
func (im Uint32Map) Swap(id id.ID, value uint32) uint32 {
	old := im.Get(id)
	im.Set(id, value)

	return old
}

// Uint64 define a generic interface for uint64 store.
type Uint64 interface {
	Set(id.ID, uint64)
	Get(id.ID) uint64
	Swap(id.ID, uint64) uint64
}

// Uint64Slice us.an implementation of Uint store.
type Uint64Slice []uint64

// NewUint64Slice returns a new Uint64Slice store of the given size
func NewUint64Slice(size int) Uint64Slice {
	return make(Uint64Slice, size)
}

// Set implements the Uint64 interface.
func (us Uint64Slice) Set(id id.ID, value uint64) {
	us[id] = value
}

// Get implements the Uint64 interface.
func (us Uint64Slice) Get(id id.ID) uint64 {
	return us[id]
}

// Swap implements the Uint64 interface.
func (us Uint64Slice) Swap(id id.ID, value uint64) uint64 {
	old := us.Get(id)
	us.Set(id, value)

	return old
}

// Uint64Map us.an implementation of Uint64 store.
type Uint64Map map[id.ID]uint64

// NewUint64Map returns a new Uint64Map store.
func NewUint64Map() Uint64Map {
	return make(Uint64Map)
}

// Set implements the Uint64 interface.
func (im Uint64Map) Set(id id.ID, value uint64) {
	im[id] = value
}

// Get implements the Uint64 interface.
func (im Uint64Map) Get(id id.ID) uint64 {
	return im[id]
}

// Swap implements the Uint64 interface.
func (im Uint64Map) Swap(id id.ID, value uint64) uint64 {
	old := im.Get(id)
	im.Set(id, value)

	return old
}
