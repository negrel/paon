package events

// LifeCycleStage define the life cycle step of a Widget.
type LifeCycleStage uint8

const (
	LCSIntial LifeCycleStage = iota
	LCSBeforeMount
	LCSMounted
	LCSBeforeUnmount
	LCSUnmounted
	_maxLifeCycle
)
