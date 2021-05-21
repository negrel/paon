package tree

// LifeCycleStage define the life cycle step of a Widget.
type LifeCycleStage uint8

const (
	BeforeMountLifeCycleStage LifeCycleStage = iota
	MountedLifeCycleStage
	BeforeUnmountLifeCycleStage
	UnmountedLifeCycleStage
	_maxLifeCycle
)
