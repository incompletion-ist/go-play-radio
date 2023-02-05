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
	Open(ctx context.Context, providedRegistry *deject.Registry) (*deject.Registry, error)

	// Close triggers a graceful close.
	Close() error

	// Closed returns a channel that will be closed when the registry is closed.
	Closed() chan<- struct{}
}
