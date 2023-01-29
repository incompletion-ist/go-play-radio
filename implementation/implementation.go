package implementation

import (
	"context"

	"github.com/incompletion-ist/go-deject/deject"
)

// Implementation is the interface for types that can be opened to return
// a deject.Registry of provided implemented functionality.
type Implementation interface {
	// Open returns a Registry of implemented functionality.
	//
	// passedRegistry is a Registry of dependencies available to the Implementation.
	// This may be used to provide loggers, etc.
	//
	// returnedRegistry is a Registry of functionality dependencies provided by the
	// Implementation. This permits an Implementation to effectively opt-in to providing
	// interfaces to the consumer.
	Open(ctx context.Context, passedRegistry *deject.Registry) (returnedRegistry *deject.Registry, err error)

	// Close closes the Implementation.
	Close() error

	// Closed returns a channel that will be closed when the Implementation is closed.
	Closed() chan<- struct{}
}
