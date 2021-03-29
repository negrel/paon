package render

type Surface interface {
	// Apply applies the given Patch on this Surface.
	Apply(Patch)

	// Clear clears this Surface.
	Clear()

	// Flush flushes all the Patch of the Surface buffer.
	Flush()
}
