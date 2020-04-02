package errors

import "fmt"

type _error struct {
	name    string
	message string
}

// NewError return a new error
func NewError(name, message string) error {
	return &_error{
		name:    name,
		message: message,
	}
}

// Error ...
func (e *_error) Error() string {
	return fmt.Sprintf("%v - %v", e.name, e.message)
}
