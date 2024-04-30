package layout

// Checkout QT doc.
// https://doc.qt.io/qt-6/qsizepolicy.html#Policy-enum
type SizePolicy uint8

const (
	GrowFlag SizePolicy = 1 << iota
	ExpandFlag
	ShrinkFlag
	IgnoreFlag
)

const (
	// SizeHint is the only acceptable size, it cannot be grown or shrunk.
	FixedSizePolicy SizePolicy = 0

	// SizeHint is minimal, and sufficient. It can be expanded, but there is no
	// advantage to it being larger. It cannot be smaller than the size provided.
	MinimumSizePolicy SizePolicy = GrowFlag

	// SizeHint is a maximum. It can be shrunk any amount without detriment if
	// needed. It cannot be larger than the size provided.
	MaximumSizePolicy SizePolicy = ShrinkFlag

	// SizeHint is best, but the widget can be shrunk and still be useful. It can
	// be expanded, but there is no advantage to it being larger.
	PreferredSizePolicy SizePolicy = GrowFlag | ShrinkFlag

	// SizeHint is a sensible size, but it can be shrunk and still be useful. It
	// can make use of extra space, so it should get as much space as possible.
	ExpandingSizePolicy SizePolicy = GrowFlag | ShrinkFlag | ExpandFlag

	// SizeHint is minimal, and sufficient. It can make use of extra space, so it
	// should get as much space as possible
	MinimumExpandingSizePolicy SizePolicy = GrowFlag | ExpandFlag

	// SizeHint is ignored. It will get as much space as possible.
	IgnoredSizePolicy SizePolicy = ShrinkFlag | GrowFlag | IgnoreFlag
)
