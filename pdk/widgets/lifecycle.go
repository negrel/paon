package widgets

// LifeCycleStage define a stage in the life of a Widget.
type LifeCycleStage uint8

const (
	// LCSInitial stage correspond to the stage of a new Widget that haven't
	// been mounted in the widget tree.
	LCSInitial LifeCycleStage = iota
	// LCSBeforeMount stage is set just before inserting the widget in a Layout.
	// The widget can already be mounted in another part of the widget tree when it enter this stage.
	LCSBeforeMount
	// LCSMounted stage is set just after a widget has been mounted.
	LCSMounted
	// LCSBeforeUnmount stage is set just before unmounting the widget of a Layout.
	LCSBeforeUnmount
	// LCSUnmounted stage is set just after a widget has been removed from his parent.
	LCSUnmounted
	_maxLifeCycle
)

// String implements the fmt.Stringer interface.
func (lcs LifeCycleStage) String() string {
	switch lcs {
	case LCSInitial:
		return "initial"
	case LCSBeforeMount:
		return "before mount"
	case LCSMounted:
		return "mounted"
	case LCSBeforeUnmount:
		return "before unmount"
	case LCSUnmounted:
		return "unmounted"
	default:
		panic("invalid life cycle stage")
	}
}
