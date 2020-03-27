package tree

// NonDocumentTypeChildNode interface contains
// methods that are particular to Node objects
// that can have a parent, but not suitable
// for DocumentType.
// https://developer.mozilla.org/en-US/docs/Web/API/NonDocumentTypeChildNode
// https://dom.spec.whatwg.org/#interface-nondocumenttypechildnode
type NonDocumentTypeChildNode interface {
	PreviousElementSibling() Element
	NextElementSibling() Element
}
