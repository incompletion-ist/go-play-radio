package wraperr

import "fmt"

// ErrorCoder is the interface for types that implement ErrorCode.
type ErrorCoder[CodeT any] interface {
	// Code returns the error code.
	Code() CodeT

	// sealed limits valid ErrorWrapper implementations to this package only.
	sealed()
}

// ErrorWrapper is the interface for types that implement WrappedError.
type ErrorWrapper interface {
	// WrappedError returns the wrapped error.
	WrappedError() error

	// sealed limits valid ErrorWrapper implementations to this package only.
	sealed()
}

// Error is a coded, wrapped error.
type Error[CodeT any] interface {
	ErrorCoder[CodeT]
	ErrorWrapper
	error
}

// Wrapf returns an Error that wraps err.
func Wrapf[CodeT any](code CodeT, err error, format string, a ...any) error {
	return wrappedError[CodeT]{
		code:    code,
		wrapped: err,
		message: fmt.Sprintf(format, a...),
	}
}

// WrapfIfErr returns an Error that wraps err.
//
// Returns nil if err is nil.
func WrapfIfErr[CodeT any](code CodeT, err error, format string, a ...any) error {
	if err == nil {
		return nil
	}

	return Wrapf(code, err, format, a...)
}

// Errorf returns an Error.
func Errorf[CodeT any](code CodeT, format string, a ...any) error {
	return Wrapf(code, nil, format, a...)
}
