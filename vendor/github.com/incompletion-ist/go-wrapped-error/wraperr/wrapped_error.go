package wraperr

// wrappedError represents a wrapped error for a generic CodeT type.
type wrappedError[CodeT any] struct {
	code    CodeT
	wrapped error
	message string
}

// Code returns the stored code.
func (err wrappedError[CodeT]) Code() CodeT {
	return err.code
}

// WrappedError returns the wrapped error.
func (err wrappedError[CodeT]) WrappedError() error {
	return err.wrapped
}

// Error returns the Error's message.
func (err wrappedError[CodeT]) Error() string {
	return err.message
}

// sealed satisfies the ErrorCoder and ErrorWrapper interfaces which require the
// unexported method to limit interface implementations to this package only.
func (err wrappedError[CodeT]) sealed() {}
