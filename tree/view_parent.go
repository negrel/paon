package tree

// ViewParent represent a View container.
type ViewParent interface {
	// FirstChild return the first children.
	FirstChild() View

	// LastChild return the last children.
	LastChild() View

	// IndexOf return the index of the given view.
	IndexOf(View) (int, error)

	// Item return the view at the given child index.
	// If there is no child at this index, return nil
	Item(int) View

	// Contains return wether or the ViewParent
	// contain the given view.
	Contains(View) bool

	// AppendChild append the given view at the end
	// of the child
	AppendChild(View) error

	// InsertBefore insert the given child before
	// the reference child. If the reference child is
	// nil, the child is appended.
	InsertBefore(View, View) error

	// RemoveChild remove the child view.
	RemoveChild(View)
}

var _ ViewParent = &_viewParent{}

type _viewParent struct {
	*_view

	childViews []View
}

/*****************************************************
 ***************** GETTERS & SETTERS *****************
 *****************************************************/
// ANCHOR Getters & setter

func (vp *_viewParent) FirstChild() View {
	if len := len(vp.childViews); len > 0 {
		return vp.childViews[0]
	}

	return nil
}

func (vp *_viewParent) Item(index int) View {
	if 0 <= index && index < len(vp.childViews) {
		return vp.childViews[index]
	}

	return nil
}

func (vp *_viewParent) LastChild() View {
	if len := len(vp.childViews); len > 0 {
		return vp.childViews[len-1]
	}

	return nil
}

/*****************************************************
 ********************** METHODS **********************
 *****************************************************/
// ANCHOR Methods

func (vp *_viewParent) AppendChild(view View) error {
	return vp.InsertBefore(view, nil)
}

func (vp *_viewParent) Contains(view View) bool {
	// Checking first class view child
	for _, child := range vp.childViews {
		if child == view {
			return true
		}
	}

	for _, child := range vp.childViews {
		if cvp, isViewParent := child.(ViewParent); isViewParent {
			if cvp.Contains(view) {
				return true
			}
		}
	}

	return false
}

func (vp *_viewParent) IndexOf(view View) (int, error) {
	for i, child := range vp.childViews {
		// Comparing pointer
		if child == view {
			return i, nil
		}
	}

	return 0, NotFoundError("The given view is not a direct child of this ViewParent.")
}

// InsertBefore insert the given view before the given
// reference child.
func (vp *_viewParent) InsertBefore(view, referenceChild View) error {
	if vp, isViewParent := view.(ViewParent); isViewParent {
		if vp.Contains(view) {
			return HierarchyRequestError("The new child contains the parent.")
		}
	}

	// Checking that view to insert is not the layout itself
	if vp == view {
		return HierarchyRequestError("The new child is the parent itself.")
	}

	// Inserting
	index, err := vp.IndexOf(referenceChild)

	if err != nil {
		index = len(vp.childViews)
	}

	var before []View = append(vp.childViews[:index], view)
	var after []View = vp.childViews[index:]

	vp.childViews = append(before, after...)

	// Adopt the child view
	view.AdoptedBy(vp)

	return nil
}

// RemoveChild remove the given child view
func (vp *_viewParent) RemoveChild(view View) {
	view.Abandoned()

	index, err := vp.IndexOf(view)

	// If found as child view
	if err == nil {
		vp.childViews = append(vp.childViews[:index], vp.childViews[index:]...)
	}
}
