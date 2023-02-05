package alteration

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// ErrorCode indicates the type of a returned error.
type ErrorCode int

const (
	// ErrorUndefined represents an undefined error.
	ErrorUndefined ErrorCode = iota

	// ErrorInvalidType indicates an invalid reflect.Type was used.
	ErrorInvalidType

	// ErrorInvalidField indicates the referenced field name isn't valid.
	ErrorInvalidField

	// ErrorInvalidValue indicates a provided value was invalid.
	ErrorInvalidValue

	// ErrorFunctionalParam indicates an error was returned by a functional
	// parameter when called.
	ErrorFunctionalParam
)

// Error is the type for errors returned by this package.
type Error = wraperr.Error[ErrorCode]
