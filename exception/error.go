package exception

import "fmt"

// Error objects are thrown when runtime errors
// occur. The Error object can also be used as a
// base object for user-defined exceptions.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Error
type Error = Exception

// EvalError object indicates an error regarding
// the global eval() function. This exception is not
// thrown by JavaScript anymore, however the EvalError
// object remains for compatibility.
// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Objets_globaux/EvalError
func EvalError(format string, msg ...interface{}) Error {
	msgg := fmt.Sprintf(format, msg...)

	return &exception{
		name:    "EvalError",
		message: msgg,
	}
}

// RangeError object indicates an error when
// a value is not in the set or range of allowed
// values.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RangeError
func RangeError(format string, msg ...interface{}) Error {
	msgg := fmt.Sprintf(format, msg...)

	return &exception{
		name:    "RangeError",
		message: msgg,
	}
}

// ReferenceError object represents an error when
// a non-existent variable is referenced.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/ReferenceError
func ReferenceError(format string, msg ...interface{}) Error {
	msgg := fmt.Sprintf(format, msg...)

	return &exception{
		name:    "ReferenceError",
		message: msgg,
	}
}

// TypeError object represents an error when an
// operation could not be performed, typically
// (but not exclusively) when a value is not of
// the expected type.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/TypeError
func TypeError(format string, msg ...interface{}) Error {
	msgg := fmt.Sprintf(format, msg...)

	return &exception{
		name:    "TypeError",
		message: msgg,
	}
}

// URIError object represents an error when a global
// URI handling function was used in a wrong way.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/URIError
func URIError(format string, msg ...interface{}) Error {
	msgg := fmt.Sprintf(format, msg...)

	return &exception{
		name:    "URIError",
		message: msgg,
	}
}
