package gom

import "strings"

// The CharacterData abstract interface represents
// a Node object that contains characters.
// https://developer.mozilla.org/en-US/docs/Web/API/CharacterData
// https://dom.spec.whatwg.org/#interface-characterdata
type CharacterData struct {
	/* INTERFACE */
	NonDocumentTypeChildNode
	/* PROPERTIES */
	*Node
	data string
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// Data return textual data contained in this object
// https://dom.spec.whatwg.org/#dom-characterdata-data
func (cd *CharacterData) Data() string {
	return cd.data
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
	cd.data = data
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
	var b strings.Builder

	b.WriteString(cd.data)
	b.WriteString(data)

	cd.data = b.String()

	return cd.data
}

// DeleteData methods remove the specified amount of
// characters, starting at the specified offset. This
// method returns the shortened string.
// https://dom.spec.whatwg.org/#dom-characterdata-deletedata
func (cd *CharacterData) DeleteData(offset, count uint) string {
	r := []rune(cd.data)
	r = append(r[offset:], r[offset+count:]...)

	cd.data = string(r)

	return cd.data
}

// InsertData methods inserts the specified characters,
// at the specified offset. This method returns the modified
// string.
// https://dom.spec.whatwg.org/#dom-characterdata-insertdata
func (cd *CharacterData) InsertData(offset uint, data string) string {
	r := []rune(cd.data)
	d := []rune(data)

	r = append(r[offset:], d...)
	r = append(r, r[offset+1:]...)

	cd.data = string(r)

	return cd.data
}

// ReplaceData replace the specified amount of characters,
// starting at the specified offset, with the specified
// string and return the modified string.
// https://dom.spec.whatwg.org/#concept-cd-replace
func (cd *CharacterData) ReplaceData(offset, count uint, data string) {
	length := cd.Length()

}

// SubstringData methods return a string with the specified
// length and starting at the specified offset.
// https://dom.spec.whatwg.org/#dom-characterdata-substringdata
func (cd *CharacterData) SubstringData(offset, count uint) string {
	r := []rune(cd.data)

	return string(r[offset : offset+count])
}
