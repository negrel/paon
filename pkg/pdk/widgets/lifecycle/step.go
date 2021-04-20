package lifecycle

type Step uint8

const (
	BeforeCreate Step = iota
	Created
	BeforeMount
	Mounted
	BeforeUpdate
	Updated
	BeforeUnmount
	Unmounted
	_maxLifeCycle
)

type Hooks [_maxLifeCycle]func()

func fn() {}

func MakeHooks() Hooks {
	return Hooks{
		BeforeCreate:  fn,
		Created:       fn,
		BeforeMount:   fn,
		Mounted:       fn,
		BeforeUpdate:  fn,
		Updated:       fn,
		BeforeUnmount: fn,
		Unmounted:     fn,
	}
}
