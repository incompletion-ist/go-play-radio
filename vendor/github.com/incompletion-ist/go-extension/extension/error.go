package extension

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// ErrorCode identifies the type of error encountered.
type ErrorCode int

const (
	// ErrorUndefined is an undefined error.
	ErrorUndefined ErrorCode = iota

	// ErrorInvalid indicates a value is invalid.
	ErrorInvalid

	// ErrorDuplicate indicates a value is duplicated.
	ErrorDuplicate

	// ErrorUnmarshal indicates an error was encountered during unmarshaling.
	ErrorUnmarshal
)

// Error is the type for errors returned by this package.
type Error = wraperr.Error[ErrorCode]
