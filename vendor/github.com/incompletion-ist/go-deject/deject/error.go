package deject

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// ErrorCode indicates which type of error occurred.
type ErrorCode int

const (
	// ErrorUndefined is an undefined error.
	ErrorUndefined int = iota

	// ErrorNilValue indicates a nil value was provided when disallowed.
	ErrorNilValue

	// ErrorInvalid indicates an invalid value was provided.
	ErrorInvalid

	// ErrorDuplicate indicates a duplicate dependency type was attempted
	// to be injected.
	ErrorDuplicate
)

// Error is the type for errors returned by this package.
type Error = wraperr.Error[ErrorCode]
