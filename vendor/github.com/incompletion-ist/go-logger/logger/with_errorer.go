package logger

// WithErrorer is the interface for types that implement WithError.
type WithErrorer[MessageT any] interface {
	// WithError returns a new MessageT which includes error.
	WithError(error) MessageT
}
