package errors

import "fmt"

type _error struct {
	name    string
	message string
}

// New return a new error
func New(name, message string) error {
	return &_error{
		name:    name,
		message: message,
	}
}

// Error ...
func (e *_error) Error() string {
	return fmt.Sprintf("%v - %v", e.name, e.message)
}

// PREMADE ERRORS

// NotFoundError return a not found error
func NotFoundError(message string) error {
	return New("NotFoundError", message)
}

// HierarchyRequestError return a hierarchy request error
func HierarchyRequestError(message string) error {
	return New("HierarchyRequestError", message)
}

// TypeError return a type error
func TypeError(message string) error {
	return New("TypeError", message)
}
