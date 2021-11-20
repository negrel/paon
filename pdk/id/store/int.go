package store

import "github.com/negrel/paon/pdk/id"

// Int define a generic interface for int store.
type Int interface {
	Set(id.ID, int)
	Get(id.ID) int
	Swap(id.ID, int) int
}

// IntSlice is an implementation of Int store.
type IntSlice []int

// NewIntSlice returns a new int store of the given size
func NewIntSlice(size int) IntSlice {
	return make(IntSlice, size)
}

// Set implements the Int interface.
func (is IntSlice) Set(id id.ID, value int) {
	is[id] = value
}

// Get implements the Int interface.
func (is IntSlice) Get(id id.ID) int {
	return is[id]
}

// Swap implements the Int interface.
func (is IntSlice) Swap(id id.ID, value int) int {
	old := is.Get(id)
	is.Set(id, value)

	return old
}

// IntMap is an implementation of Int store.
type IntMap map[id.ID]int

// NewIntMap returns a new IntMap store.
func NewIntMap() IntMap {
	return make(IntMap)
}

// Set implements the Int interface.
func (im IntMap) Set(id id.ID, value int) {
	im[id] = value
}

// Get implements the Int interface.
func (im IntMap) Get(id id.ID) int {
	return im[id]
}

// Swap implements the Int interface.
func (im IntMap) Swap(id id.ID, value int) int {
	old := im.Get(id)
	im.Set(id, value)

	return old
}

// Int8 define a generic interface for int8 store.
type Int8 interface {
	Set(id.ID, int8)
	Get(id.ID) int8
	Swap(id.ID, int8) int8
}

// Int8Slice is an implementation of Int8 store.
type Int8Slice []int8

// NewInt8Slice returns a new Int8Slice store of the given size
func NewInt8Slice(size int) Int8Slice {
	return make(Int8Slice, size)
}

// Set implements the Int8 interface.
func (is Int8Slice) Set(id id.ID, value int8) {
	is[id] = value
}

// Get implements the Int8 interface.
func (is Int8Slice) Get(id id.ID) int8 {
	return is[id]
}

// Swap implements the Int8 interface.
func (is Int8Slice) Swap(id id.ID, value int8) int8 {
	old := is.Get(id)
	is.Set(id, value)

	return old
}

// Int8Map is an implementation of Int8 store.
type Int8Map map[id.ID]int8

// NewInt8Map returns a new Int8Map store.
func NewInt8Map() Int8Map {
	return make(Int8Map)
}

// Set implements the Int8 interface.
func (im Int8Map) Set(id id.ID, value int8) {
	im[id] = value
}

// Get implements the Int8 interface.
func (im Int8Map) Get(id id.ID) int8 {
	return im[id]
}

// Swap implements the Int8 interface.
func (im Int8Map) Swap(id id.ID, value int8) int8 {
	old := im.Get(id)
	im.Set(id, value)

	return old
}

// Int16 define a generic interface for int16 store.
type Int16 interface {
	Set(id.ID, int16)
	Get(id.ID) int16
	Swap(id.ID, int16) int16
}

// Int16Slice is an implementation of Int16 store.
type Int16Slice []int16

// NewInt16Slice returns a new int16 store of the given size
func NewInt16Slice(size int) Int16Slice {
	return make(Int16Slice, size)
}

// Set implements the Int16 interface.
func (is Int16Slice) Set(id id.ID, value int16) {
	is[id] = value
}

// Get implements the Int16 interface.
func (is Int16Slice) Get(id id.ID) int16 {
	return is[id]
}

// Swap implements the Int16 interface.
func (is Int16Slice) Swap(id id.ID, value int16) int16 {
	old := is.Get(id)
	is.Set(id, value)

	return old
}

// Int16Map is an implementation of Int16 store.
type Int16Map map[id.ID]int16

// NewInt16Map returns a new Int16Map store.
func NewInt16Map() Int16Map {
	return make(Int16Map)
}

// Set implements the Int16 interface.
func (im Int16Map) Set(id id.ID, value int16) {
	im[id] = value
}

// Get implements the Int16Map interface.
func (im Int16Map) Get(id id.ID) int16 {
	return im[id]
}

// Swap implements the Int16 interface.
func (im Int16Map) Swap(id id.ID, value int16) int16 {
	old := im.Get(id)
	im.Set(id, value)

	return old
}

// Int32 define a generic interface for int32 store.
type Int32 interface {
	Set(id.ID, int32)
	Get(id.ID) int32
	Swap(id.ID, int32) int32
}

// Int32Slice is an implementation of Int32 store.
type Int32Slice []int32

// NewInt32Slice returns a new Int32Slice store of the given size
func NewInt32Slice(size int) Int32Slice {
	return make(Int32Slice, size)
}

// Set implements the Int32 interface.
func (is Int32Slice) Set(id id.ID, value int32) {
	is[id] = value
}

// Get implements the Int32 interface.
func (is Int32Slice) Get(id id.ID) int32 {
	return is[id]
}

// Swap implements the Int32 interface.
func (is Int32Slice) Swap(id id.ID, value int32) int32 {
	old := is.Get(id)
	is.Set(id, value)

	return old
}

// Int32Map is an implementation of Int32 store.
type Int32Map map[id.ID]int32

// NewInt32Map returns a new Int32Slice store.
func NewInt32Map() Int32Map {
	return make(Int32Map)
}

// Set implements the Int32 interface.
func (im Int32Map) Set(id id.ID, value int32) {
	im[id] = value
}

// Get implements the Int32 interface.
func (im Int32Map) Get(id id.ID) int32 {
	return im[id]
}

// Swap implements the Int32 interface.
func (im Int32Map) Swap(id id.ID, value int32) int32 {
	old := im.Get(id)
	im.Set(id, value)

	return old
}

// Int64 define a generic interface for int64 store.
type Int64 interface {
	Set(id.ID, int64)
	Get(id.ID) int64
	Swap(id.ID, int64) int64
}

// Int64Slice is an implementation of Int store.
type Int64Slice []int64

// NewInt64Slice returns a new Int64Slice store of the given size
func NewInt64Slice(size int) Int64Slice {
	return make(Int64Slice, size)
}

// Set implements the Int64 interface.
func (is Int64Slice) Set(id id.ID, value int64) {
	is[id] = value
}

// Get implements the Int64 interface.
func (is Int64Slice) Get(id id.ID) int64 {
	return is[id]
}

// Swap implements the Int64 interface.
func (is Int64Slice) Swap(id id.ID, value int64) int64 {
	old := is.Get(id)
	is.Set(id, value)

	return old
}

// Int64Map is an implementation of Int64 store.
type Int64Map map[id.ID]int64

// NewInt64Map returns a new Int64Map store.
func NewInt64Map() Int64Map {
	return make(Int64Map)
}

// Set implements the Int64 interface.
func (im Int64Map) Set(id id.ID, value int64) {
	im[id] = value
}

// Get implements the Int64 interface.
func (im Int64Map) Get(id id.ID) int64 {
	return im[id]
}

// Swap implements the Int64 interface.
func (im Int64Map) Swap(id id.ID, value int64) int64 {
	old := im.Get(id)
	im.Set(id, value)

	return old
}
