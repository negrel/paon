package styles

// Styled define object with styling properties.
type Styled interface {
	Style() Style
}

// Style is a generic interface for widget styling.
type Style interface {
	// Compute styling properties and returns it. This enable Style object
	// to compute dynamically their properties and, therefore, support style inheritance,
	// relative units (percentage, viewport), etc.
	Compute() ComputedStyle
}
