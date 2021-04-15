package observables

import (
	"github.com/negrel/debuggo/pkg/assert"
	"unsafe"
)

// InvalidationHandler is called whenever an Observable becomes invalid.
type InvalidationHandler *func(Observable)

// Observable is a generic interface implemented by all Observable types.
type Observable interface {
	AddInvalidationHandler(...InvalidationHandler)
	RemoveInvalidationHandler(...InvalidationHandler)
}

var _ Observable = &observable{}

type observable struct {
	ptr                  Observable
	invalidationHandlers map[InvalidationHandler]struct{}
}

// NewObservable returns a new instantiated Observable.
func NewObservable() Observable {
	o := newCompositeObservable(nil)
	o.ptr = o

	return o
}

// NewCompositeObservable returns a new Observable object. Unlike NewObservable,
// this function returns an Observable object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this Observable will become invalid.
func NewCompositeObservable(ptr Observable) Observable {
	return newCompositeObservable(ptr)
}

func newCompositeObservable(ptr Observable) *observable {
	return &observable{
		ptr:                  ptr,
		invalidationHandlers: make(map[InvalidationHandler]struct{}),
	}
}

// AddInvalidationHandler implements the Observable interface.
func (o observable) AddInvalidationHandler(handlers ...InvalidationHandler) {
	assert.True(o.isValid())

	for _, handler := range handlers {
		o.invalidationHandlers[handler] = struct{}{}
	}
}

// RemoveInvalidationHandler implements the Observable interface.
func (o observable) RemoveInvalidationHandler(handlers ...InvalidationHandler) {
	assert.True(o.isValid())

	for _, handler := range handlers {
		_, ok := o.invalidationHandlers[handler]
		if ok {
			delete(o.invalidationHandlers, handler)
		}
	}
}

func (o observable) invalid() {
	for handler := range o.invalidationHandlers {
		(*handler)(o.ptr)
	}

	o.invalidationHandlers = nil
}

func (o observable) isValid() bool {
	return o.invalidationHandlers != nil
}

// IntChangeHandler is called whenever the value of an ObservableInt change.
type IntChangeHandler *func(observable ObservableInt, old, new int)

// ObservableInt is an entity that wraps a value and allows to observe value changes.
type ObservableInt interface {
	Observable

	Get() int
	Set(int)

	AddChangeHandler(...IntChangeHandler)
	RemoveChangeHandler(...IntChangeHandler)
}

var _ ObservableInt = &observableInt{}

type observableInt struct {
	*observable

	data     int
	handlers map[IntChangeHandler]struct{}
}

// NewObservableInt returns a new instantiated ObservableInt.
func NewObservableInt(data int) ObservableInt {
	oo := newCompositeObservableInt(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableInt returns a new ObservableInt object. Unlike NewObservableInt,
// this function returns an ObservableInt object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableInt will become invalid.
func NewCompositeObservableInt(data int) ObservableInt {
	oo := newCompositeObservableInt(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableInt(ptr ObservableInt, data int) *observableInt {
	return &observableInt{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[IntChangeHandler]struct{}),
	}
}

// Get implements the ObservableInt interface.
func (oo *observableInt) Get() int {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableInt interface.
func (oo *observableInt) Set(newValue int) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableInt interface.
func (oo *observableInt) AddChangeHandler(handlers ...IntChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableInt interface.
func (oo *observableInt) RemoveChangeHandler(handlers ...IntChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// IntBinding define the minimum required to implement a binding to int.
type IntBinding interface {
	ObservableInt

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableInt

	// Bind starts observing the dependencies for change.
	Bind(...ObservableInt)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableInt)
}

var _ IntBinding = &intBinding{}

type intBinding struct {
	*observableInt

	dependencies []ObservableInt
}

func NewIntBinding(data int) IntBinding {
	ob := newCompositeIntBinding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeIntBinding(ptr IntBinding, data int) IntBinding {
	return newCompositeIntBinding(ptr, data)
}

func newCompositeIntBinding(ptr IntBinding, data int) *intBinding {
	return &intBinding{
		observableInt: newCompositeObservableInt(ptr, data),
		dependencies:  make([]ObservableInt, 0, 8),
	}
}

func (ob *intBinding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *intBinding) GetDependencies() []ObservableInt {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *intBinding) Bind(observables ...ObservableInt) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableInt, oldValue, newValue int) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *intBinding) Unbind(observables ...ObservableInt) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Int8ChangeHandler is called whenever the value of an ObservableInt8 change.
type Int8ChangeHandler *func(observable ObservableInt8, old, new int8)

// ObservableInt8 is an entity that wraps a value and allows to observe value changes.
type ObservableInt8 interface {
	Observable

	Get() int8
	Set(int8)

	AddChangeHandler(...Int8ChangeHandler)
	RemoveChangeHandler(...Int8ChangeHandler)
}

var _ ObservableInt8 = &observableInt8{}

type observableInt8 struct {
	*observable

	data     int8
	handlers map[Int8ChangeHandler]struct{}
}

// NewObservableInt8 returns a new instantiated ObservableInt8.
func NewObservableInt8(data int8) ObservableInt8 {
	oo := newCompositeObservableInt8(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableInt8 returns a new ObservableInt8 object. Unlike NewObservableInt8,
// this function returns an ObservableInt8 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableInt8 will become invalid.
func NewCompositeObservableInt8(data int8) ObservableInt8 {
	oo := newCompositeObservableInt8(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableInt8(ptr ObservableInt8, data int8) *observableInt8 {
	return &observableInt8{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Int8ChangeHandler]struct{}),
	}
}

// Get implements the ObservableInt8 interface.
func (oo *observableInt8) Get() int8 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableInt8 interface.
func (oo *observableInt8) Set(newValue int8) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableInt8 interface.
func (oo *observableInt8) AddChangeHandler(handlers ...Int8ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableInt8 interface.
func (oo *observableInt8) RemoveChangeHandler(handlers ...Int8ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Int8Binding define the minimum required to implement a binding to int8.
type Int8Binding interface {
	ObservableInt8

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableInt8

	// Bind starts observing the dependencies for change.
	Bind(...ObservableInt8)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableInt8)
}

var _ Int8Binding = &int8Binding{}

type int8Binding struct {
	*observableInt8

	dependencies []ObservableInt8
}

func NewInt8Binding(data int8) Int8Binding {
	ob := newCompositeInt8Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeInt8Binding(ptr Int8Binding, data int8) Int8Binding {
	return newCompositeInt8Binding(ptr, data)
}

func newCompositeInt8Binding(ptr Int8Binding, data int8) *int8Binding {
	return &int8Binding{
		observableInt8: newCompositeObservableInt8(ptr, data),
		dependencies:   make([]ObservableInt8, 0, 8),
	}
}

func (ob *int8Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *int8Binding) GetDependencies() []ObservableInt8 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *int8Binding) Bind(observables ...ObservableInt8) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableInt8, oldValue, newValue int8) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *int8Binding) Unbind(observables ...ObservableInt8) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Int16ChangeHandler is called whenever the value of an ObservableInt16 change.
type Int16ChangeHandler *func(observable ObservableInt16, old, new int16)

// ObservableInt16 is an entity that wraps a value and allows to observe value changes.
type ObservableInt16 interface {
	Observable

	Get() int16
	Set(int16)

	AddChangeHandler(...Int16ChangeHandler)
	RemoveChangeHandler(...Int16ChangeHandler)
}

var _ ObservableInt16 = &observableInt16{}

type observableInt16 struct {
	*observable

	data     int16
	handlers map[Int16ChangeHandler]struct{}
}

// NewObservableInt16 returns a new instantiated ObservableInt16.
func NewObservableInt16(data int16) ObservableInt16 {
	oo := newCompositeObservableInt16(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableInt16 returns a new ObservableInt16 object. Unlike NewObservableInt16,
// this function returns an ObservableInt16 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableInt16 will become invalid.
func NewCompositeObservableInt16(data int16) ObservableInt16 {
	oo := newCompositeObservableInt16(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableInt16(ptr ObservableInt16, data int16) *observableInt16 {
	return &observableInt16{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Int16ChangeHandler]struct{}),
	}
}

// Get implements the ObservableInt16 interface.
func (oo *observableInt16) Get() int16 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableInt16 interface.
func (oo *observableInt16) Set(newValue int16) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableInt16 interface.
func (oo *observableInt16) AddChangeHandler(handlers ...Int16ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableInt16 interface.
func (oo *observableInt16) RemoveChangeHandler(handlers ...Int16ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Int16Binding define the minimum required to implement a binding to int16.
type Int16Binding interface {
	ObservableInt16

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableInt16

	// Bind starts observing the dependencies for change.
	Bind(...ObservableInt16)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableInt16)
}

var _ Int16Binding = &int16Binding{}

type int16Binding struct {
	*observableInt16

	dependencies []ObservableInt16
}

func NewInt16Binding(data int16) Int16Binding {
	ob := newCompositeInt16Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeInt16Binding(ptr Int16Binding, data int16) Int16Binding {
	return newCompositeInt16Binding(ptr, data)
}

func newCompositeInt16Binding(ptr Int16Binding, data int16) *int16Binding {
	return &int16Binding{
		observableInt16: newCompositeObservableInt16(ptr, data),
		dependencies:    make([]ObservableInt16, 0, 8),
	}
}

func (ob *int16Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *int16Binding) GetDependencies() []ObservableInt16 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *int16Binding) Bind(observables ...ObservableInt16) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableInt16, oldValue, newValue int16) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *int16Binding) Unbind(observables ...ObservableInt16) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Int32ChangeHandler is called whenever the value of an ObservableInt32 change.
type Int32ChangeHandler *func(observable ObservableInt32, old, new int32)

// ObservableInt32 is an entity that wraps a value and allows to observe value changes.
type ObservableInt32 interface {
	Observable

	Get() int32
	Set(int32)

	AddChangeHandler(...Int32ChangeHandler)
	RemoveChangeHandler(...Int32ChangeHandler)
}

var _ ObservableInt32 = &observableInt32{}

type observableInt32 struct {
	*observable

	data     int32
	handlers map[Int32ChangeHandler]struct{}
}

// NewObservableInt32 returns a new instantiated ObservableInt32.
func NewObservableInt32(data int32) ObservableInt32 {
	oo := newCompositeObservableInt32(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableInt32 returns a new ObservableInt32 object. Unlike NewObservableInt32,
// this function returns an ObservableInt32 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableInt32 will become invalid.
func NewCompositeObservableInt32(data int32) ObservableInt32 {
	oo := newCompositeObservableInt32(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableInt32(ptr ObservableInt32, data int32) *observableInt32 {
	return &observableInt32{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Int32ChangeHandler]struct{}),
	}
}

// Get implements the ObservableInt32 interface.
func (oo *observableInt32) Get() int32 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableInt32 interface.
func (oo *observableInt32) Set(newValue int32) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableInt32 interface.
func (oo *observableInt32) AddChangeHandler(handlers ...Int32ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableInt32 interface.
func (oo *observableInt32) RemoveChangeHandler(handlers ...Int32ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Int32Binding define the minimum required to implement a binding to int32.
type Int32Binding interface {
	ObservableInt32

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableInt32

	// Bind starts observing the dependencies for change.
	Bind(...ObservableInt32)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableInt32)
}

var _ Int32Binding = &int32Binding{}

type int32Binding struct {
	*observableInt32

	dependencies []ObservableInt32
}

func NewInt32Binding(data int32) Int32Binding {
	ob := newCompositeInt32Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeInt32Binding(ptr Int32Binding, data int32) Int32Binding {
	return newCompositeInt32Binding(ptr, data)
}

func newCompositeInt32Binding(ptr Int32Binding, data int32) *int32Binding {
	return &int32Binding{
		observableInt32: newCompositeObservableInt32(ptr, data),
		dependencies:    make([]ObservableInt32, 0, 8),
	}
}

func (ob *int32Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *int32Binding) GetDependencies() []ObservableInt32 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *int32Binding) Bind(observables ...ObservableInt32) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableInt32, oldValue, newValue int32) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *int32Binding) Unbind(observables ...ObservableInt32) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Int64ChangeHandler is called whenever the value of an ObservableInt64 change.
type Int64ChangeHandler *func(observable ObservableInt64, old, new int64)

// ObservableInt64 is an entity that wraps a value and allows to observe value changes.
type ObservableInt64 interface {
	Observable

	Get() int64
	Set(int64)

	AddChangeHandler(...Int64ChangeHandler)
	RemoveChangeHandler(...Int64ChangeHandler)
}

var _ ObservableInt64 = &observableInt64{}

type observableInt64 struct {
	*observable

	data     int64
	handlers map[Int64ChangeHandler]struct{}
}

// NewObservableInt64 returns a new instantiated ObservableInt64.
func NewObservableInt64(data int64) ObservableInt64 {
	oo := newCompositeObservableInt64(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableInt64 returns a new ObservableInt64 object. Unlike NewObservableInt64,
// this function returns an ObservableInt64 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableInt64 will become invalid.
func NewCompositeObservableInt64(data int64) ObservableInt64 {
	oo := newCompositeObservableInt64(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableInt64(ptr ObservableInt64, data int64) *observableInt64 {
	return &observableInt64{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Int64ChangeHandler]struct{}),
	}
}

// Get implements the ObservableInt64 interface.
func (oo *observableInt64) Get() int64 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableInt64 interface.
func (oo *observableInt64) Set(newValue int64) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableInt64 interface.
func (oo *observableInt64) AddChangeHandler(handlers ...Int64ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableInt64 interface.
func (oo *observableInt64) RemoveChangeHandler(handlers ...Int64ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Int64Binding define the minimum required to implement a binding to int64.
type Int64Binding interface {
	ObservableInt64

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableInt64

	// Bind starts observing the dependencies for change.
	Bind(...ObservableInt64)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableInt64)
}

var _ Int64Binding = &int64Binding{}

type int64Binding struct {
	*observableInt64

	dependencies []ObservableInt64
}

func NewInt64Binding(data int64) Int64Binding {
	ob := newCompositeInt64Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeInt64Binding(ptr Int64Binding, data int64) Int64Binding {
	return newCompositeInt64Binding(ptr, data)
}

func newCompositeInt64Binding(ptr Int64Binding, data int64) *int64Binding {
	return &int64Binding{
		observableInt64: newCompositeObservableInt64(ptr, data),
		dependencies:    make([]ObservableInt64, 0, 8),
	}
}

func (ob *int64Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *int64Binding) GetDependencies() []ObservableInt64 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *int64Binding) Bind(observables ...ObservableInt64) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableInt64, oldValue, newValue int64) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *int64Binding) Unbind(observables ...ObservableInt64) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// UintChangeHandler is called whenever the value of an ObservableUint change.
type UintChangeHandler *func(observable ObservableUint, old, new uint)

// ObservableUint is an entity that wraps a value and allows to observe value changes.
type ObservableUint interface {
	Observable

	Get() uint
	Set(uint)

	AddChangeHandler(...UintChangeHandler)
	RemoveChangeHandler(...UintChangeHandler)
}

var _ ObservableUint = &observableUint{}

type observableUint struct {
	*observable

	data     uint
	handlers map[UintChangeHandler]struct{}
}

// NewObservableUint returns a new instantiated ObservableUint.
func NewObservableUint(data uint) ObservableUint {
	oo := newCompositeObservableUint(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableUint returns a new ObservableUint object. Unlike NewObservableUint,
// this function returns an ObservableUint object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableUint will become invalid.
func NewCompositeObservableUint(data uint) ObservableUint {
	oo := newCompositeObservableUint(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableUint(ptr ObservableUint, data uint) *observableUint {
	return &observableUint{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[UintChangeHandler]struct{}),
	}
}

// Get implements the ObservableUint interface.
func (oo *observableUint) Get() uint {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableUint interface.
func (oo *observableUint) Set(newValue uint) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableUint interface.
func (oo *observableUint) AddChangeHandler(handlers ...UintChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableUint interface.
func (oo *observableUint) RemoveChangeHandler(handlers ...UintChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// UintBinding define the minimum required to implement a binding to uint.
type UintBinding interface {
	ObservableUint

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableUint

	// Bind starts observing the dependencies for change.
	Bind(...ObservableUint)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableUint)
}

var _ UintBinding = &uintBinding{}

type uintBinding struct {
	*observableUint

	dependencies []ObservableUint
}

func NewUintBinding(data uint) UintBinding {
	ob := newCompositeUintBinding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeUintBinding(ptr UintBinding, data uint) UintBinding {
	return newCompositeUintBinding(ptr, data)
}

func newCompositeUintBinding(ptr UintBinding, data uint) *uintBinding {
	return &uintBinding{
		observableUint: newCompositeObservableUint(ptr, data),
		dependencies:   make([]ObservableUint, 0, 8),
	}
}

func (ob *uintBinding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *uintBinding) GetDependencies() []ObservableUint {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *uintBinding) Bind(observables ...ObservableUint) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableUint, oldValue, newValue uint) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *uintBinding) Unbind(observables ...ObservableUint) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Uint8ChangeHandler is called whenever the value of an ObservableUint8 change.
type Uint8ChangeHandler *func(observable ObservableUint8, old, new uint8)

// ObservableUint8 is an entity that wraps a value and allows to observe value changes.
type ObservableUint8 interface {
	Observable

	Get() uint8
	Set(uint8)

	AddChangeHandler(...Uint8ChangeHandler)
	RemoveChangeHandler(...Uint8ChangeHandler)
}

var _ ObservableUint8 = &observableUint8{}

type observableUint8 struct {
	*observable

	data     uint8
	handlers map[Uint8ChangeHandler]struct{}
}

// NewObservableUint8 returns a new instantiated ObservableUint8.
func NewObservableUint8(data uint8) ObservableUint8 {
	oo := newCompositeObservableUint8(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableUint8 returns a new ObservableUint8 object. Unlike NewObservableUint8,
// this function returns an ObservableUint8 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableUint8 will become invalid.
func NewCompositeObservableUint8(data uint8) ObservableUint8 {
	oo := newCompositeObservableUint8(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableUint8(ptr ObservableUint8, data uint8) *observableUint8 {
	return &observableUint8{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Uint8ChangeHandler]struct{}),
	}
}

// Get implements the ObservableUint8 interface.
func (oo *observableUint8) Get() uint8 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableUint8 interface.
func (oo *observableUint8) Set(newValue uint8) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableUint8 interface.
func (oo *observableUint8) AddChangeHandler(handlers ...Uint8ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableUint8 interface.
func (oo *observableUint8) RemoveChangeHandler(handlers ...Uint8ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Uint8Binding define the minimum required to implement a binding to uint8.
type Uint8Binding interface {
	ObservableUint8

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableUint8

	// Bind starts observing the dependencies for change.
	Bind(...ObservableUint8)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableUint8)
}

var _ Uint8Binding = &uint8Binding{}

type uint8Binding struct {
	*observableUint8

	dependencies []ObservableUint8
}

func NewUint8Binding(data uint8) Uint8Binding {
	ob := newCompositeUint8Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeUint8Binding(ptr Uint8Binding, data uint8) Uint8Binding {
	return newCompositeUint8Binding(ptr, data)
}

func newCompositeUint8Binding(ptr Uint8Binding, data uint8) *uint8Binding {
	return &uint8Binding{
		observableUint8: newCompositeObservableUint8(ptr, data),
		dependencies:    make([]ObservableUint8, 0, 8),
	}
}

func (ob *uint8Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *uint8Binding) GetDependencies() []ObservableUint8 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *uint8Binding) Bind(observables ...ObservableUint8) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableUint8, oldValue, newValue uint8) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *uint8Binding) Unbind(observables ...ObservableUint8) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Uint16ChangeHandler is called whenever the value of an ObservableUint16 change.
type Uint16ChangeHandler *func(observable ObservableUint16, old, new uint16)

// ObservableUint16 is an entity that wraps a value and allows to observe value changes.
type ObservableUint16 interface {
	Observable

	Get() uint16
	Set(uint16)

	AddChangeHandler(...Uint16ChangeHandler)
	RemoveChangeHandler(...Uint16ChangeHandler)
}

var _ ObservableUint16 = &observableUint16{}

type observableUint16 struct {
	*observable

	data     uint16
	handlers map[Uint16ChangeHandler]struct{}
}

// NewObservableUint16 returns a new instantiated ObservableUint16.
func NewObservableUint16(data uint16) ObservableUint16 {
	oo := newCompositeObservableUint16(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableUint16 returns a new ObservableUint16 object. Unlike NewObservableUint16,
// this function returns an ObservableUint16 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableUint16 will become invalid.
func NewCompositeObservableUint16(data uint16) ObservableUint16 {
	oo := newCompositeObservableUint16(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableUint16(ptr ObservableUint16, data uint16) *observableUint16 {
	return &observableUint16{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Uint16ChangeHandler]struct{}),
	}
}

// Get implements the ObservableUint16 interface.
func (oo *observableUint16) Get() uint16 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableUint16 interface.
func (oo *observableUint16) Set(newValue uint16) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableUint16 interface.
func (oo *observableUint16) AddChangeHandler(handlers ...Uint16ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableUint16 interface.
func (oo *observableUint16) RemoveChangeHandler(handlers ...Uint16ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Uint16Binding define the minimum required to implement a binding to uint16.
type Uint16Binding interface {
	ObservableUint16

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableUint16

	// Bind starts observing the dependencies for change.
	Bind(...ObservableUint16)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableUint16)
}

var _ Uint16Binding = &uint16Binding{}

type uint16Binding struct {
	*observableUint16

	dependencies []ObservableUint16
}

func NewUint16Binding(data uint16) Uint16Binding {
	ob := newCompositeUint16Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeUint16Binding(ptr Uint16Binding, data uint16) Uint16Binding {
	return newCompositeUint16Binding(ptr, data)
}

func newCompositeUint16Binding(ptr Uint16Binding, data uint16) *uint16Binding {
	return &uint16Binding{
		observableUint16: newCompositeObservableUint16(ptr, data),
		dependencies:     make([]ObservableUint16, 0, 8),
	}
}

func (ob *uint16Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *uint16Binding) GetDependencies() []ObservableUint16 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *uint16Binding) Bind(observables ...ObservableUint16) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableUint16, oldValue, newValue uint16) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *uint16Binding) Unbind(observables ...ObservableUint16) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Uint32ChangeHandler is called whenever the value of an ObservableUint32 change.
type Uint32ChangeHandler *func(observable ObservableUint32, old, new uint32)

// ObservableUint32 is an entity that wraps a value and allows to observe value changes.
type ObservableUint32 interface {
	Observable

	Get() uint32
	Set(uint32)

	AddChangeHandler(...Uint32ChangeHandler)
	RemoveChangeHandler(...Uint32ChangeHandler)
}

var _ ObservableUint32 = &observableUint32{}

type observableUint32 struct {
	*observable

	data     uint32
	handlers map[Uint32ChangeHandler]struct{}
}

// NewObservableUint32 returns a new instantiated ObservableUint32.
func NewObservableUint32(data uint32) ObservableUint32 {
	oo := newCompositeObservableUint32(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableUint32 returns a new ObservableUint32 object. Unlike NewObservableUint32,
// this function returns an ObservableUint32 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableUint32 will become invalid.
func NewCompositeObservableUint32(data uint32) ObservableUint32 {
	oo := newCompositeObservableUint32(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableUint32(ptr ObservableUint32, data uint32) *observableUint32 {
	return &observableUint32{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Uint32ChangeHandler]struct{}),
	}
}

// Get implements the ObservableUint32 interface.
func (oo *observableUint32) Get() uint32 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableUint32 interface.
func (oo *observableUint32) Set(newValue uint32) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableUint32 interface.
func (oo *observableUint32) AddChangeHandler(handlers ...Uint32ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableUint32 interface.
func (oo *observableUint32) RemoveChangeHandler(handlers ...Uint32ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Uint32Binding define the minimum required to implement a binding to uint32.
type Uint32Binding interface {
	ObservableUint32

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableUint32

	// Bind starts observing the dependencies for change.
	Bind(...ObservableUint32)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableUint32)
}

var _ Uint32Binding = &uint32Binding{}

type uint32Binding struct {
	*observableUint32

	dependencies []ObservableUint32
}

func NewUint32Binding(data uint32) Uint32Binding {
	ob := newCompositeUint32Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeUint32Binding(ptr Uint32Binding, data uint32) Uint32Binding {
	return newCompositeUint32Binding(ptr, data)
}

func newCompositeUint32Binding(ptr Uint32Binding, data uint32) *uint32Binding {
	return &uint32Binding{
		observableUint32: newCompositeObservableUint32(ptr, data),
		dependencies:     make([]ObservableUint32, 0, 8),
	}
}

func (ob *uint32Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *uint32Binding) GetDependencies() []ObservableUint32 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *uint32Binding) Bind(observables ...ObservableUint32) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableUint32, oldValue, newValue uint32) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *uint32Binding) Unbind(observables ...ObservableUint32) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Uint64ChangeHandler is called whenever the value of an ObservableUint64 change.
type Uint64ChangeHandler *func(observable ObservableUint64, old, new uint64)

// ObservableUint64 is an entity that wraps a value and allows to observe value changes.
type ObservableUint64 interface {
	Observable

	Get() uint64
	Set(uint64)

	AddChangeHandler(...Uint64ChangeHandler)
	RemoveChangeHandler(...Uint64ChangeHandler)
}

var _ ObservableUint64 = &observableUint64{}

type observableUint64 struct {
	*observable

	data     uint64
	handlers map[Uint64ChangeHandler]struct{}
}

// NewObservableUint64 returns a new instantiated ObservableUint64.
func NewObservableUint64(data uint64) ObservableUint64 {
	oo := newCompositeObservableUint64(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableUint64 returns a new ObservableUint64 object. Unlike NewObservableUint64,
// this function returns an ObservableUint64 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableUint64 will become invalid.
func NewCompositeObservableUint64(data uint64) ObservableUint64 {
	oo := newCompositeObservableUint64(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableUint64(ptr ObservableUint64, data uint64) *observableUint64 {
	return &observableUint64{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Uint64ChangeHandler]struct{}),
	}
}

// Get implements the ObservableUint64 interface.
func (oo *observableUint64) Get() uint64 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableUint64 interface.
func (oo *observableUint64) Set(newValue uint64) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableUint64 interface.
func (oo *observableUint64) AddChangeHandler(handlers ...Uint64ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableUint64 interface.
func (oo *observableUint64) RemoveChangeHandler(handlers ...Uint64ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Uint64Binding define the minimum required to implement a binding to uint64.
type Uint64Binding interface {
	ObservableUint64

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableUint64

	// Bind starts observing the dependencies for change.
	Bind(...ObservableUint64)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableUint64)
}

var _ Uint64Binding = &uint64Binding{}

type uint64Binding struct {
	*observableUint64

	dependencies []ObservableUint64
}

func NewUint64Binding(data uint64) Uint64Binding {
	ob := newCompositeUint64Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeUint64Binding(ptr Uint64Binding, data uint64) Uint64Binding {
	return newCompositeUint64Binding(ptr, data)
}

func newCompositeUint64Binding(ptr Uint64Binding, data uint64) *uint64Binding {
	return &uint64Binding{
		observableUint64: newCompositeObservableUint64(ptr, data),
		dependencies:     make([]ObservableUint64, 0, 8),
	}
}

func (ob *uint64Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *uint64Binding) GetDependencies() []ObservableUint64 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *uint64Binding) Bind(observables ...ObservableUint64) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableUint64, oldValue, newValue uint64) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *uint64Binding) Unbind(observables ...ObservableUint64) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Float32ChangeHandler is called whenever the value of an ObservableFloat32 change.
type Float32ChangeHandler *func(observable ObservableFloat32, old, new float32)

// ObservableFloat32 is an entity that wraps a value and allows to observe value changes.
type ObservableFloat32 interface {
	Observable

	Get() float32
	Set(float32)

	AddChangeHandler(...Float32ChangeHandler)
	RemoveChangeHandler(...Float32ChangeHandler)
}

var _ ObservableFloat32 = &observableFloat32{}

type observableFloat32 struct {
	*observable

	data     float32
	handlers map[Float32ChangeHandler]struct{}
}

// NewObservableFloat32 returns a new instantiated ObservableFloat32.
func NewObservableFloat32(data float32) ObservableFloat32 {
	oo := newCompositeObservableFloat32(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableFloat32 returns a new ObservableFloat32 object. Unlike NewObservableFloat32,
// this function returns an ObservableFloat32 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableFloat32 will become invalid.
func NewCompositeObservableFloat32(data float32) ObservableFloat32 {
	oo := newCompositeObservableFloat32(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableFloat32(ptr ObservableFloat32, data float32) *observableFloat32 {
	return &observableFloat32{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Float32ChangeHandler]struct{}),
	}
}

// Get implements the ObservableFloat32 interface.
func (oo *observableFloat32) Get() float32 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableFloat32 interface.
func (oo *observableFloat32) Set(newValue float32) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableFloat32 interface.
func (oo *observableFloat32) AddChangeHandler(handlers ...Float32ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableFloat32 interface.
func (oo *observableFloat32) RemoveChangeHandler(handlers ...Float32ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Float32Binding define the minimum required to implement a binding to float32.
type Float32Binding interface {
	ObservableFloat32

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableFloat32

	// Bind starts observing the dependencies for change.
	Bind(...ObservableFloat32)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableFloat32)
}

var _ Float32Binding = &float32Binding{}

type float32Binding struct {
	*observableFloat32

	dependencies []ObservableFloat32
}

func NewFloat32Binding(data float32) Float32Binding {
	ob := newCompositeFloat32Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeFloat32Binding(ptr Float32Binding, data float32) Float32Binding {
	return newCompositeFloat32Binding(ptr, data)
}

func newCompositeFloat32Binding(ptr Float32Binding, data float32) *float32Binding {
	return &float32Binding{
		observableFloat32: newCompositeObservableFloat32(ptr, data),
		dependencies:      make([]ObservableFloat32, 0, 8),
	}
}

func (ob *float32Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *float32Binding) GetDependencies() []ObservableFloat32 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *float32Binding) Bind(observables ...ObservableFloat32) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableFloat32, oldValue, newValue float32) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *float32Binding) Unbind(observables ...ObservableFloat32) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Float64ChangeHandler is called whenever the value of an ObservableFloat64 change.
type Float64ChangeHandler *func(observable ObservableFloat64, old, new float64)

// ObservableFloat64 is an entity that wraps a value and allows to observe value changes.
type ObservableFloat64 interface {
	Observable

	Get() float64
	Set(float64)

	AddChangeHandler(...Float64ChangeHandler)
	RemoveChangeHandler(...Float64ChangeHandler)
}

var _ ObservableFloat64 = &observableFloat64{}

type observableFloat64 struct {
	*observable

	data     float64
	handlers map[Float64ChangeHandler]struct{}
}

// NewObservableFloat64 returns a new instantiated ObservableFloat64.
func NewObservableFloat64(data float64) ObservableFloat64 {
	oo := newCompositeObservableFloat64(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableFloat64 returns a new ObservableFloat64 object. Unlike NewObservableFloat64,
// this function returns an ObservableFloat64 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableFloat64 will become invalid.
func NewCompositeObservableFloat64(data float64) ObservableFloat64 {
	oo := newCompositeObservableFloat64(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableFloat64(ptr ObservableFloat64, data float64) *observableFloat64 {
	return &observableFloat64{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Float64ChangeHandler]struct{}),
	}
}

// Get implements the ObservableFloat64 interface.
func (oo *observableFloat64) Get() float64 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableFloat64 interface.
func (oo *observableFloat64) Set(newValue float64) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableFloat64 interface.
func (oo *observableFloat64) AddChangeHandler(handlers ...Float64ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableFloat64 interface.
func (oo *observableFloat64) RemoveChangeHandler(handlers ...Float64ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Float64Binding define the minimum required to implement a binding to float64.
type Float64Binding interface {
	ObservableFloat64

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableFloat64

	// Bind starts observing the dependencies for change.
	Bind(...ObservableFloat64)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableFloat64)
}

var _ Float64Binding = &float64Binding{}

type float64Binding struct {
	*observableFloat64

	dependencies []ObservableFloat64
}

func NewFloat64Binding(data float64) Float64Binding {
	ob := newCompositeFloat64Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeFloat64Binding(ptr Float64Binding, data float64) Float64Binding {
	return newCompositeFloat64Binding(ptr, data)
}

func newCompositeFloat64Binding(ptr Float64Binding, data float64) *float64Binding {
	return &float64Binding{
		observableFloat64: newCompositeObservableFloat64(ptr, data),
		dependencies:      make([]ObservableFloat64, 0, 8),
	}
}

func (ob *float64Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *float64Binding) GetDependencies() []ObservableFloat64 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *float64Binding) Bind(observables ...ObservableFloat64) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableFloat64, oldValue, newValue float64) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *float64Binding) Unbind(observables ...ObservableFloat64) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Complex64ChangeHandler is called whenever the value of an ObservableComplex64 change.
type Complex64ChangeHandler *func(observable ObservableComplex64, old, new complex64)

// ObservableComplex64 is an entity that wraps a value and allows to observe value changes.
type ObservableComplex64 interface {
	Observable

	Get() complex64
	Set(complex64)

	AddChangeHandler(...Complex64ChangeHandler)
	RemoveChangeHandler(...Complex64ChangeHandler)
}

var _ ObservableComplex64 = &observableComplex64{}

type observableComplex64 struct {
	*observable

	data     complex64
	handlers map[Complex64ChangeHandler]struct{}
}

// NewObservableComplex64 returns a new instantiated ObservableComplex64.
func NewObservableComplex64(data complex64) ObservableComplex64 {
	oo := newCompositeObservableComplex64(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableComplex64 returns a new ObservableComplex64 object. Unlike NewObservableComplex64,
// this function returns an ObservableComplex64 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableComplex64 will become invalid.
func NewCompositeObservableComplex64(data complex64) ObservableComplex64 {
	oo := newCompositeObservableComplex64(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableComplex64(ptr ObservableComplex64, data complex64) *observableComplex64 {
	return &observableComplex64{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Complex64ChangeHandler]struct{}),
	}
}

// Get implements the ObservableComplex64 interface.
func (oo *observableComplex64) Get() complex64 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableComplex64 interface.
func (oo *observableComplex64) Set(newValue complex64) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableComplex64 interface.
func (oo *observableComplex64) AddChangeHandler(handlers ...Complex64ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableComplex64 interface.
func (oo *observableComplex64) RemoveChangeHandler(handlers ...Complex64ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Complex64Binding define the minimum required to implement a binding to complex64.
type Complex64Binding interface {
	ObservableComplex64

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableComplex64

	// Bind starts observing the dependencies for change.
	Bind(...ObservableComplex64)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableComplex64)
}

var _ Complex64Binding = &complex64Binding{}

type complex64Binding struct {
	*observableComplex64

	dependencies []ObservableComplex64
}

func NewComplex64Binding(data complex64) Complex64Binding {
	ob := newCompositeComplex64Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeComplex64Binding(ptr Complex64Binding, data complex64) Complex64Binding {
	return newCompositeComplex64Binding(ptr, data)
}

func newCompositeComplex64Binding(ptr Complex64Binding, data complex64) *complex64Binding {
	return &complex64Binding{
		observableComplex64: newCompositeObservableComplex64(ptr, data),
		dependencies:        make([]ObservableComplex64, 0, 8),
	}
}

func (ob *complex64Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *complex64Binding) GetDependencies() []ObservableComplex64 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *complex64Binding) Bind(observables ...ObservableComplex64) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableComplex64, oldValue, newValue complex64) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *complex64Binding) Unbind(observables ...ObservableComplex64) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// Complex128ChangeHandler is called whenever the value of an ObservableComplex128 change.
type Complex128ChangeHandler *func(observable ObservableComplex128, old, new complex128)

// ObservableComplex128 is an entity that wraps a value and allows to observe value changes.
type ObservableComplex128 interface {
	Observable

	Get() complex128
	Set(complex128)

	AddChangeHandler(...Complex128ChangeHandler)
	RemoveChangeHandler(...Complex128ChangeHandler)
}

var _ ObservableComplex128 = &observableComplex128{}

type observableComplex128 struct {
	*observable

	data     complex128
	handlers map[Complex128ChangeHandler]struct{}
}

// NewObservableComplex128 returns a new instantiated ObservableComplex128.
func NewObservableComplex128(data complex128) ObservableComplex128 {
	oo := newCompositeObservableComplex128(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableComplex128 returns a new ObservableComplex128 object. Unlike NewObservableComplex128,
// this function returns an ObservableComplex128 object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableComplex128 will become invalid.
func NewCompositeObservableComplex128(data complex128) ObservableComplex128 {
	oo := newCompositeObservableComplex128(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableComplex128(ptr ObservableComplex128, data complex128) *observableComplex128 {
	return &observableComplex128{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[Complex128ChangeHandler]struct{}),
	}
}

// Get implements the ObservableComplex128 interface.
func (oo *observableComplex128) Get() complex128 {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableComplex128 interface.
func (oo *observableComplex128) Set(newValue complex128) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableComplex128 interface.
func (oo *observableComplex128) AddChangeHandler(handlers ...Complex128ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableComplex128 interface.
func (oo *observableComplex128) RemoveChangeHandler(handlers ...Complex128ChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// Complex128Binding define the minimum required to implement a binding to complex128.
type Complex128Binding interface {
	ObservableComplex128

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableComplex128

	// Bind starts observing the dependencies for change.
	Bind(...ObservableComplex128)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableComplex128)
}

var _ Complex128Binding = &complex128Binding{}

type complex128Binding struct {
	*observableComplex128

	dependencies []ObservableComplex128
}

func NewComplex128Binding(data complex128) Complex128Binding {
	ob := newCompositeComplex128Binding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeComplex128Binding(ptr Complex128Binding, data complex128) Complex128Binding {
	return newCompositeComplex128Binding(ptr, data)
}

func newCompositeComplex128Binding(ptr Complex128Binding, data complex128) *complex128Binding {
	return &complex128Binding{
		observableComplex128: newCompositeObservableComplex128(ptr, data),
		dependencies:         make([]ObservableComplex128, 0, 8),
	}
}

func (ob *complex128Binding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *complex128Binding) GetDependencies() []ObservableComplex128 {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *complex128Binding) Bind(observables ...ObservableComplex128) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableComplex128, oldValue, newValue complex128) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *complex128Binding) Unbind(observables ...ObservableComplex128) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// StringChangeHandler is called whenever the value of an ObservableString change.
type StringChangeHandler *func(observable ObservableString, old, new string)

// ObservableString is an entity that wraps a value and allows to observe value changes.
type ObservableString interface {
	Observable

	Get() string
	Set(string)

	AddChangeHandler(...StringChangeHandler)
	RemoveChangeHandler(...StringChangeHandler)
}

var _ ObservableString = &observableString{}

type observableString struct {
	*observable

	data     string
	handlers map[StringChangeHandler]struct{}
}

// NewObservableString returns a new instantiated ObservableString.
func NewObservableString(data string) ObservableString {
	oo := newCompositeObservableString(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableString returns a new ObservableString object. Unlike NewObservableString,
// this function returns an ObservableString object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableString will become invalid.
func NewCompositeObservableString(data string) ObservableString {
	oo := newCompositeObservableString(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableString(ptr ObservableString, data string) *observableString {
	return &observableString{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[StringChangeHandler]struct{}),
	}
}

// Get implements the ObservableString interface.
func (oo *observableString) Get() string {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableString interface.
func (oo *observableString) Set(newValue string) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableString interface.
func (oo *observableString) AddChangeHandler(handlers ...StringChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableString interface.
func (oo *observableString) RemoveChangeHandler(handlers ...StringChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// StringBinding define the minimum required to implement a binding to string.
type StringBinding interface {
	ObservableString

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableString

	// Bind starts observing the dependencies for change.
	Bind(...ObservableString)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableString)
}

var _ StringBinding = &stringBinding{}

type stringBinding struct {
	*observableString

	dependencies []ObservableString
}

func NewStringBinding(data string) StringBinding {
	ob := newCompositeStringBinding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeStringBinding(ptr StringBinding, data string) StringBinding {
	return newCompositeStringBinding(ptr, data)
}

func newCompositeStringBinding(ptr StringBinding, data string) *stringBinding {
	return &stringBinding{
		observableString: newCompositeObservableString(ptr, data),
		dependencies:     make([]ObservableString, 0, 8),
	}
}

func (ob *stringBinding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *stringBinding) GetDependencies() []ObservableString {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *stringBinding) Bind(observables ...ObservableString) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableString, oldValue, newValue string) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *stringBinding) Unbind(observables ...ObservableString) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// ByteChangeHandler is called whenever the value of an ObservableByte change.
type ByteChangeHandler *func(observable ObservableByte, old, new byte)

// ObservableByte is an entity that wraps a value and allows to observe value changes.
type ObservableByte interface {
	Observable

	Get() byte
	Set(byte)

	AddChangeHandler(...ByteChangeHandler)
	RemoveChangeHandler(...ByteChangeHandler)
}

var _ ObservableByte = &observableByte{}

type observableByte struct {
	*observable

	data     byte
	handlers map[ByteChangeHandler]struct{}
}

// NewObservableByte returns a new instantiated ObservableByte.
func NewObservableByte(data byte) ObservableByte {
	oo := newCompositeObservableByte(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableByte returns a new ObservableByte object. Unlike NewObservableByte,
// this function returns an ObservableByte object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableByte will become invalid.
func NewCompositeObservableByte(data byte) ObservableByte {
	oo := newCompositeObservableByte(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableByte(ptr ObservableByte, data byte) *observableByte {
	return &observableByte{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[ByteChangeHandler]struct{}),
	}
}

// Get implements the ObservableByte interface.
func (oo *observableByte) Get() byte {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableByte interface.
func (oo *observableByte) Set(newValue byte) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableByte interface.
func (oo *observableByte) AddChangeHandler(handlers ...ByteChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableByte interface.
func (oo *observableByte) RemoveChangeHandler(handlers ...ByteChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// ByteBinding define the minimum required to implement a binding to byte.
type ByteBinding interface {
	ObservableByte

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableByte

	// Bind starts observing the dependencies for change.
	Bind(...ObservableByte)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableByte)
}

var _ ByteBinding = &byteBinding{}

type byteBinding struct {
	*observableByte

	dependencies []ObservableByte
}

func NewByteBinding(data byte) ByteBinding {
	ob := newCompositeByteBinding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeByteBinding(ptr ByteBinding, data byte) ByteBinding {
	return newCompositeByteBinding(ptr, data)
}

func newCompositeByteBinding(ptr ByteBinding, data byte) *byteBinding {
	return &byteBinding{
		observableByte: newCompositeObservableByte(ptr, data),
		dependencies:   make([]ObservableByte, 0, 8),
	}
}

func (ob *byteBinding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *byteBinding) GetDependencies() []ObservableByte {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *byteBinding) Bind(observables ...ObservableByte) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableByte, oldValue, newValue byte) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *byteBinding) Unbind(observables ...ObservableByte) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// RuneChangeHandler is called whenever the value of an ObservableRune change.
type RuneChangeHandler *func(observable ObservableRune, old, new rune)

// ObservableRune is an entity that wraps a value and allows to observe value changes.
type ObservableRune interface {
	Observable

	Get() rune
	Set(rune)

	AddChangeHandler(...RuneChangeHandler)
	RemoveChangeHandler(...RuneChangeHandler)
}

var _ ObservableRune = &observableRune{}

type observableRune struct {
	*observable

	data     rune
	handlers map[RuneChangeHandler]struct{}
}

// NewObservableRune returns a new instantiated ObservableRune.
func NewObservableRune(data rune) ObservableRune {
	oo := newCompositeObservableRune(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableRune returns a new ObservableRune object. Unlike NewObservableRune,
// this function returns an ObservableRune object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableRune will become invalid.
func NewCompositeObservableRune(data rune) ObservableRune {
	oo := newCompositeObservableRune(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableRune(ptr ObservableRune, data rune) *observableRune {
	return &observableRune{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[RuneChangeHandler]struct{}),
	}
}

// Get implements the ObservableRune interface.
func (oo *observableRune) Get() rune {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableRune interface.
func (oo *observableRune) Set(newValue rune) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableRune interface.
func (oo *observableRune) AddChangeHandler(handlers ...RuneChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableRune interface.
func (oo *observableRune) RemoveChangeHandler(handlers ...RuneChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// RuneBinding define the minimum required to implement a binding to rune.
type RuneBinding interface {
	ObservableRune

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableRune

	// Bind starts observing the dependencies for change.
	Bind(...ObservableRune)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableRune)
}

var _ RuneBinding = &runeBinding{}

type runeBinding struct {
	*observableRune

	dependencies []ObservableRune
}

func NewRuneBinding(data rune) RuneBinding {
	ob := newCompositeRuneBinding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeRuneBinding(ptr RuneBinding, data rune) RuneBinding {
	return newCompositeRuneBinding(ptr, data)
}

func newCompositeRuneBinding(ptr RuneBinding, data rune) *runeBinding {
	return &runeBinding{
		observableRune: newCompositeObservableRune(ptr, data),
		dependencies:   make([]ObservableRune, 0, 8),
	}
}

func (ob *runeBinding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *runeBinding) GetDependencies() []ObservableRune {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *runeBinding) Bind(observables ...ObservableRune) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableRune, oldValue, newValue rune) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *runeBinding) Unbind(observables ...ObservableRune) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// BoolChangeHandler is called whenever the value of an ObservableBool change.
type BoolChangeHandler *func(observable ObservableBool, old, new bool)

// ObservableBool is an entity that wraps a value and allows to observe value changes.
type ObservableBool interface {
	Observable

	Get() bool
	Set(bool)

	AddChangeHandler(...BoolChangeHandler)
	RemoveChangeHandler(...BoolChangeHandler)
}

var _ ObservableBool = &observableBool{}

type observableBool struct {
	*observable

	data     bool
	handlers map[BoolChangeHandler]struct{}
}

// NewObservableBool returns a new instantiated ObservableBool.
func NewObservableBool(data bool) ObservableBool {
	oo := newCompositeObservableBool(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableBool returns a new ObservableBool object. Unlike NewObservableBool,
// this function returns an ObservableBool object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableBool will become invalid.
func NewCompositeObservableBool(data bool) ObservableBool {
	oo := newCompositeObservableBool(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableBool(ptr ObservableBool, data bool) *observableBool {
	return &observableBool{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[BoolChangeHandler]struct{}),
	}
}

// Get implements the ObservableBool interface.
func (oo *observableBool) Get() bool {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableBool interface.
func (oo *observableBool) Set(newValue bool) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableBool interface.
func (oo *observableBool) AddChangeHandler(handlers ...BoolChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableBool interface.
func (oo *observableBool) RemoveChangeHandler(handlers ...BoolChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// BoolBinding define the minimum required to implement a binding to bool.
type BoolBinding interface {
	ObservableBool

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableBool

	// Bind starts observing the dependencies for change.
	Bind(...ObservableBool)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableBool)
}

var _ BoolBinding = &boolBinding{}

type boolBinding struct {
	*observableBool

	dependencies []ObservableBool
}

func NewBoolBinding(data bool) BoolBinding {
	ob := newCompositeBoolBinding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeBoolBinding(ptr BoolBinding, data bool) BoolBinding {
	return newCompositeBoolBinding(ptr, data)
}

func newCompositeBoolBinding(ptr BoolBinding, data bool) *boolBinding {
	return &boolBinding{
		observableBool: newCompositeObservableBool(ptr, data),
		dependencies:   make([]ObservableBool, 0, 8),
	}
}

func (ob *boolBinding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *boolBinding) GetDependencies() []ObservableBool {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *boolBinding) Bind(observables ...ObservableBool) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableBool, oldValue, newValue bool) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *boolBinding) Unbind(observables ...ObservableBool) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}

// ObjectChangeHandler is called whenever the value of an ObservableObject change.
type ObjectChangeHandler *func(observable ObservableObject, old, new unsafe.Pointer)

// ObservableObject is an entity that wraps a value and allows to observe value changes.
type ObservableObject interface {
	Observable

	Get() unsafe.Pointer
	Set(unsafe.Pointer)

	AddChangeHandler(...ObjectChangeHandler)
	RemoveChangeHandler(...ObjectChangeHandler)
}

var _ ObservableObject = &observableObject{}

type observableObject struct {
	*observable

	data     unsafe.Pointer
	handlers map[ObjectChangeHandler]struct{}
}

// NewObservableObject returns a new instantiated ObservableObject.
func NewObservableObject(data unsafe.Pointer) ObservableObject {
	oo := newCompositeObservableObject(nil, data)
	oo.ptr = oo

	return oo
}

// NewCompositeObservableObject returns a new ObservableObject object. Unlike NewObservableObject,
// this function returns an ObservableObject object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, invalid handler will receive a pointer to the composite
// struct when this ObservableObject will become invalid.
func NewCompositeObservableObject(data unsafe.Pointer) ObservableObject {
	oo := newCompositeObservableObject(nil, data)
	oo.ptr = oo

	return oo
}

func newCompositeObservableObject(ptr ObservableObject, data unsafe.Pointer) *observableObject {
	return &observableObject{
		observable: newCompositeObservable(ptr),
		data:       data,
		handlers:   make(map[ObjectChangeHandler]struct{}),
	}
}

// Get implements the ObservableObject interface.
func (oo *observableObject) Get() unsafe.Pointer {
	assert.True(oo.isValid())

	return oo.data
}

// Set implements the ObservableObject interface.
func (oo *observableObject) Set(newValue unsafe.Pointer) {
	assert.True(oo.isValid())

	old := oo.data
	oo.data = newValue
	for handler := range oo.handlers {
		(*handler)(oo, old, newValue)
	}
}

// AddChangeHandler implements the ObservableObject interface.
func (oo *observableObject) AddChangeHandler(handlers ...ObjectChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		oo.handlers[handler] = struct{}{}
	}
}

// RemoveChangeHandler implements the ObservableObject interface.
func (oo *observableObject) RemoveChangeHandler(handlers ...ObjectChangeHandler) {
	assert.True(oo.isValid())

	for _, handler := range handlers {
		_, ok := oo.handlers[handler]
		if ok {
			delete(oo.handlers, handler)
		}
	}
}

// ObjectBinding define the minimum required to implement a binding to unsafe.Pointer.
type ObjectBinding interface {
	ObservableObject

	// Dispose signals to the bindings that it will not be used anymore
	// and any references can be removed.
	Dispose()

	// GetDependencies returns the all the dependencies of this binding.
	GetDependencies() []ObservableObject

	// Bind starts observing the dependencies for change.
	Bind(...ObservableObject)

	// Unbind stops observing the dependencies for change.
	Unbind(...ObservableObject)
}

var _ ObjectBinding = &objectBinding{}

type objectBinding struct {
	*observableObject

	dependencies []ObservableObject
}

func NewObjectBinding(data unsafe.Pointer) ObjectBinding {
	ob := newCompositeObjectBinding(nil, data)
	ob.ptr = ob

	return ob
}

func NewCompositeObjectBinding(ptr ObjectBinding, data unsafe.Pointer) ObjectBinding {
	return newCompositeObjectBinding(ptr, data)
}

func newCompositeObjectBinding(ptr ObjectBinding, data unsafe.Pointer) *objectBinding {
	return &objectBinding{
		observableObject: newCompositeObservableObject(ptr, data),
		dependencies:     make([]ObservableObject, 0, 8),
	}
}

func (ob *objectBinding) Dispose() {
	assert.True(ob.isValid())

	ob.invalid()
	ob.dependencies = nil
}

func (ob *objectBinding) GetDependencies() []ObservableObject {
	assert.True(ob.isValid())

	return ob.dependencies
}

func (ob *objectBinding) Bind(observables ...ObservableObject) {
	assert.True(ob.isValid())

	for _, observable := range observables {
		ob.dependencies = append(ob.dependencies, observable)

		invalidationHandler := func(_ Observable) {
			ob.Unbind(observable)
		}

		observable.AddInvalidationHandler(&invalidationHandler)

		changeHandler := func(_ ObservableObject, oldValue, newValue unsafe.Pointer) {
			if newValue == ob.Get() {
				return
			}

			ob.Set(newValue)
		}

		observable.AddChangeHandler(&changeHandler)
	}
}

func (ob *objectBinding) Unbind(observables ...ObservableObject) {
	assert.True(ob.isValid())

	for _, toUnbind := range observables {
		for i, observable := range ob.dependencies {
			if toUnbind == observable {
				ob.dependencies[i] = ob.dependencies[len(ob.dependencies)-1]
				ob.dependencies = ob.dependencies[:len(ob.dependencies)-1]
			}
		}
	}
}
