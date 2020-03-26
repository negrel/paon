package exception

import (
	"fmt"
	"io"
)

// Exception is a type of object that represents an
// error and which can be thrown or treated as a
// first class value by implementations.
// https://heycam.github.io/webidl/#idl-exceptions
type Exception interface {
	/* GETTERS & SETTERS (props) */
	Message() string
	Name() string
	/* METHODS */
	String() string
	Fprint(w io.Writer)
	Print()
}

var _ Exception = &exception{}

type exception struct {
	message string
	name    string
}

// New return a new GOMException.
func New(code int, format string, msg ...interface{}) Exception {
	var msgg string = fmt.Sprintf(format, msg...)

	return &exception{
		name:    Map[code],
		message: msgg,
	}
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// Message return the error message
func (e *exception) Message() string {
	return e.message
}

// Name return the error name
func (e *exception) Name() string {
	return e.name
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Fprint method print the error to the given
// writer
func (e *exception) Fprint(w io.Writer) {
	fmt.Fprint(w, e.String())
}

// Print the GOM Error
func (e *exception) Print() {
	fmt.Print(e.String())
}

// String method return the formatted string.
func (e *exception) String() string {
	return fmt.Sprintf("[%v] - %v", e.Name(), e.Message())
}
