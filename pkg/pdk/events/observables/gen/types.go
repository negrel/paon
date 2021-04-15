package main

var types = [...]string{
	// int
	"int",
	"int8",
	"int16",
	"int32",
	"int64",

	// uint
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",

	// float
	"float32",
	"float64",

	// complex
	"complex64",
	"complex128",

	"string",
	"byte",
	"rune",
	"bool",

	"interface{}",
}

var integers []string = types[:5]
var unsignedIntegers []string = types[5:10]
var floats []string = types[10:12]
var complexes []string = types[12:14]
