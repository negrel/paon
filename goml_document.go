package gom

// GOMLDocument is a light HTMLDocument
// used for non-web document.
type GOMLDocument struct {
	document
	title      string
	readyState bool
}
