package tree

import (
	e "github.com/negrel/gom/exception"
)

// The CharacterData abstract interface represents
// a Node object that contains characters.
// https://developer.mozilla.org/en-US/docs/Web/API/CharacterData
// https://dom.spec.whatwg.org/#interface-characterdata
type CharacterData struct {
	/* INTERFACE */
	NonDocumentTypeChildNode
	/* PROPERTIES */
	*Node
	data []rune
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// Data return textual data contained in this object
// https://dom.spec.whatwg.org/#dom-characterdata-data
func (cd *CharacterData) Data() string {
	return string(cd.data)
}

// Length return the size of the string contained in
// CharacterData.data.
// https://dom.spec.whatwg.org/#dom-characterdata-length
func (cd *CharacterData) Length() int {
	return len(cd.data)
}

// SetData set the textual data conatined in this object.
// https://dom.spec.whatwg.org/#dom-characterdata-data
func (cd *CharacterData) SetData(data string) {
	cd.data = []rune(data)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// AppendData method append the given string to the
// CharacterData data. This method returns data contains
// the concatenated string.
// https://dom.spec.whatwg.org/#dom-characterdata-appenddata
func (cd *CharacterData) AppendData(data string) string {
	cd.ReplaceData(cd.Length(), 0, data)
	return cd.Data()
}

// DeleteData methods remove the specified amount of
// characters, starting at the specified offset. This
// method returns the shortened string.
// https://dom.spec.whatwg.org/#dom-characterdata-deletedata
func (cd *CharacterData) DeleteData(offset, count int) string {
	cd.ReplaceData(offset, count, "")
	return cd.Data()
}

// InsertData methods inserts the specified characters,
// at the specified offset. This method returns the modified
// string.
// https://dom.spec.whatwg.org/#dom-characterdata-insertdata
func (cd *CharacterData) InsertData(offset int, data string) string {
	cd.ReplaceData(offset, 0, data)
	return cd.Data()
}

// ReplaceData replace the specified amount of characters,
// starting at the specified offset, with the specified
// string and return the modified string.
// https://dom.spec.whatwg.org/#concept-cd-replace
func (cd *CharacterData) ReplaceData(offset, count int, data string) e.Exception {
	length := cd.Length()

	if offset > length {
		return e.New(e.IndexSizeError, "", "")
	}

	if (offset + count) > length {
		count = length - offset
	}

	d := append(cd.data[offset:], []rune(data)...)
	cd.data = append(d, cd.data[offset+count:]...)

	return nil
}

// SubstringData methods return a string with the specified
// length and starting at the specified offset.
// https://dom.spec.whatwg.org/#dom-characterdata-substringdata
func (cd *CharacterData) SubstringData(offset, count int) (string, e.Exception) {
	length := cd.Length()

	if offset > length {
		return "", e.RangeError("", "")
	}

	if (offset + count) > length {
		return string(cd.data[offset:]), nil
	}

	return string(cd.data[offset : offset+count]), nil
}
