package gom

// DocumentType represent a node containing a doctype.
// https://developer.mozilla.org/en-US/docs/Web/API/documentType
// https://dom.spec.whatwg.org/#documentType
type DocumentType interface {
	/* Private */
	setName(string)
	setPublicId(string)
	setSystemId(string)
	/* EMBEDDED INTERFACE */
	Node
	/* GETTERS & SETTERS (props) */
	Name() string
	PublicId() string
	SystemId() string
}

var _ DocumentType = &documentType{}

type documentType struct {
	*Node
	name     string
	publicId string
	systemId string
}

func newDocumentType(name string) DocumentType {
	return &documentType{
		node:     &Node{},
		name:     name,
		publicId: "",
		systemId: "",
	}
}

func (dt *documentType) setName(name string) {
	dt.name = name
}

func (dt *documentType) setPublicId(pId string) {
	dt.publicId = pId
}

func (dt *documentType) setSystemId(sId string) {
	dt.systemId = sId
}

/*****************************************************
 **************** Embedded interface *****************
 *****************************************************/
// ANCHOR Embedded interface

/* Node */
/* - Props */

// NodeName return the GOML-uppercased name
func (dt *documentType) NodeName() string {
	return dt.name
}

// NodeType return the "DocumentTypeNode" type.
func (dt *documentType) NodeType() NodeType {
	return DocumentTypeNode
}

/* - Methods */

// CloneNode return a clone of the DocumentType.
func (dt *documentType) CloneNode(_ bool) Node {
	clone := newDocumentType(dt.Name())

	clone.setPublicId(dt.publicId)
	clone.setSystemId(dt.publicId)

	return clone
}

// IsEqualNode method return whether two DocumentType are equal.
func (dt *documentType) isEqualNode(other Node) bool {
	if other == nil {
		goto notEqual
	}

	if dt.NodeType() != other.NodeType() {
		goto notEqual
	}

	// Type switch
	switch otherDt := other.(type) {
	case DocumentType:
		// Document Type specific test
		if dt.name != otherDt.Name() {
			goto notEqual
		}

		if dt.publicId != otherDt.PublicId() {
			goto notEqual
		}

		if dt.systemId != otherDt.SystemId() {
			goto notEqual
		}

	default:
		goto notEqual
	}

	return true

notEqual:
	return false
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// Name return the name of the document type
// eg "goml" for <!DOCTYPE GOML>
func (dt *documentType) Name() string {
	return dt.name
}

// PublicId return an empty string
func (dt *documentType) PublicId() string {
	return dt.publicId
}

// SystemId return an empty string
func (dt *documentType) SystemId() string {
	return dt.systemId
}
