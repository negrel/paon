package tree

// Style define all the existing style
type Style struct {
	// Computed parent stack style to one style
	parentStyle *Style

	// Style of the node
	ownStyle []*Style
}
