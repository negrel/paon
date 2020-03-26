package exception

// NOTE Some DOMException will never be used. (OFFICIAL DOM)

// DOMException is the possible error returned
// by DOM function.
// https://heycam.github.io/webidl/#idl-DOMException
const (
	_ = iota
	// Deprecated. Use RangeError instead.
	IndexSizeError

	// Deprecated. Use RangeError instead.
	GOMStringSizeError

	// The operation would yield an incorrect node tree.
	HierarchyRequestError

	// The object is in the wrong document.
	WrongDocumentError

	// The string contains invalid characters.
	InvalidCharacterError

	// Deprecated.
	NoDataAllowedError

	// The object can not be modified.
	NoModificationAllowedError

	// The object can not be found here.
	NotFoundError

	// The operation is not supported.
	NotSupportedError

	// The attribute is in use.
	InUseAttributeError

	// The object is in an invalid state.
	InvalidStateError

	// The string did not match the expected pattern.
	SyntaxError

	// The object can not be modified in this way.
	InvalidModificationError

	// The operation is not allowed by Namespaces in XML.
	NamespaceError

	// Deprecated. Use TypeError for invalid arguments,
	// NotSupportedError DOMException for unsupported
	// operations, and NotAllowedError DOMException
	// for denied requests instead.
	InvalidAccessError

	// Deprecated.
	ValidationError

	// Deprecated. Use TypeError instead.
	TypeMismatchError

	// The operation is insecure.
	SecurityError

	// A network error occurred.
	NetworkError

	// The operation was aborted.
	AbortError

	// The given URL does not match another URL.
	URLMismatchError

	// The quota has been exceeded.
	QuotaExceededError

	// The operation timed out.
	TimeoutError

	// The supplied node is incorrect or has an incorrect
	// ancestor for this operation.
	InvalidNodeTypeError

	// The object can not be cloned.
	DataCloneError

	// The encoding operation (either encoded or decoding)
	// failed.
	EncodingError

	// The I/O read operation failed.
	NotReadableError

	// The operation failed for an unknown transient reason.
	UnknownError

	// A mutation operation in a transaction failed because
	// a constraint was not satisfied.
	ConstraintError

	// Provided data is inadequate.
	DataError

	// A request was placed against a transaction which is
	// currently not active, or which is finished.
	TransactionInactiveError

	// The mutating operation was attempted in a readonly
	// transaction.
	ReadOnlyError

	// An attempt was made to open a database using a lower
	// version than the existing version.
	VersionError

	// The operation failed for an operation-specific reason.
	OperationError

	// The request is not allowed by the user agent or the
	// platform in the current context, possibly because
	// the user denied permission.
	NotAllowedError
)

// Map exception code to their name.
var Map = map[int]string{
	IndexSizeError:             "IndexSizeError",
	GOMStringSizeError:         "GOMStringSizeError",
	HierarchyRequestError:      "HierarchyRequestError",
	WrongDocumentError:         "WrongDocumentError",
	InvalidCharacterError:      "InvalidCharacterError",
	NoDataAllowedError:         "NoDataAllowedError",
	NoModificationAllowedError: "NoModificationAllowedError",
	NotFoundError:              "NotFoundError",
	NotSupportedError:          "NotSupportedError",
	InUseAttributeError:        "InUseAttributeError",
	InvalidStateError:          "InvalidStateError",
	SyntaxError:                "SyntaxError",
	InvalidModificationError:   "InvalidModificationError",
	NamespaceError:             "NamespaceError",
	InvalidAccessError:         "InvalidAccessError",
	ValidationError:            "ValidationError",
	TypeMismatchError:          "TypeMismatchError",
	SecurityError:              "SecurityError",
	NetworkError:               "NetworkError",
	AbortError:                 "AbortError",
	URLMismatchError:           "URLMismatchError",
	QuotaExceededError:         "QuotaExceededError",
	TimeoutError:               "TimeoutError",
	InvalidNodeTypeError:       "InvalidNodeTypeError",
	DataCloneError:             "DataCloneError",
	EncodingError:              "EncodingError",
	NotReadableError:           "NotReadableError",
	UnknownError:               "UnknownError",
	ConstraintError:            "ConstraintError",
	DataError:                  "DataError",
	TransactionInactiveError:   "TransactionInactiveError",
	ReadOnlyError:              "ReadOnlyError",
	VersionError:               "VersionError",
	OperationError:             "OperationError",
	NotAllowedError:            "NotAllowedError",
}
