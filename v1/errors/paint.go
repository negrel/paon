package errors

// PREMADE PAINT ERRORS

// SurfaceError return a surface error
func SurfaceError(message string) error {
	return NewError("SurfaceError", message)
}
