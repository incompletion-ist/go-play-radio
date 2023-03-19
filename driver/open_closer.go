package driver

import (
	"context"

	"github.com/incompletion-ist/go-deject/deject"
)

// OpenCloser is the interface for types that can open and close deject.Registries.
type OpenCloser interface {
	// Open prepares and returns the resulting deject.Registry.
	//
	// providedRegistry is a Registry of dependencies available to the RegistryOpenCloser. It may be
	// used to provide loggers, etc.
	Open(context.Context, ...Option) (*deject.Registry, error)

	// Close triggers a graceful close with a nil CloseError. Returns an error if encountered
	// while attempting to perform the Close.
	Close() error

	// Closed returns a channel that will be closed when the registry is closed.
	Closed() <-chan struct{}

	// CloseError returns the error that triggered the close. Returns nil if not closed, or if it
	// was a graceful (non-error triggered) close.
	CloseError() error
}
