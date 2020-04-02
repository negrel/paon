package errors

// PREMADE WIDGET ERRORS

// NotFoundError return a not found error
func NotFoundError(message string) error {
	return NewError("NotFoundError", message)
}

// HierarchyRequestError return a hierarchy request error
func HierarchyRequestError(message string) error {
	return NewError("HierarchyRequestError", message)
}
