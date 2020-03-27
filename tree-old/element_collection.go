package tree

// ElementCollection is a live collection of element.
type ElementCollection struct {
	col []*Element
}

func newElementCollection() *ElementCollection {
	return &ElementCollection{}
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// Length return the element collection length
func (el *ElementCollection) Length() int {
	return len(el.col)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Item return the specific element at the zero-based
// index.
func (el *ElementCollection) Item(index int) *Element {
	if index >= 0 && index < el.Length() {
		return el.col[index]
	}

	return nil
}
