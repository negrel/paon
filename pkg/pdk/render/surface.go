package render

type Surface interface {
	// Apply applies the given Patch on this Surface.
	Apply(Patch)

	// Clear clears this Surface.
	Clear()

	// Flush flushes all the Patch in this Surface buffer.
	Flush()
}
