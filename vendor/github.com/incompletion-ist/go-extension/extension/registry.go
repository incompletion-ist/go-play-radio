package extension

import (
	"sync"

	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// Registry is a registry for Extensions of type T.
type Registry[T any] struct {
	extensions map[string]Extension[T]
	mu         sync.Mutex
}

// NewRegistry returns a new Registry for type T.
func NewRegistry[T any]() *Registry[T] {
	return &Registry[T]{}
}

// Register adds Extensions to the Registry.
func (r *Registry[T]) Register(extensions ...Extension[T]) error {
	if r == nil {
		return wraperr.Errorf(ErrorInvalid, "extension: Registry is nil")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, p := range extensions {
		if validateErr := p.validate(); validateErr != nil {
			return validateErr
		}

		if _, ok := r.extensions[p.Name]; ok {
			return wraperr.Errorf(ErrorDuplicate, "extension: Extension with name %q already exists", p.Name)
		}

		if r.extensions == nil {
			r.extensions = make(map[string]Extension[T])
		}

		r.extensions[p.Name] = p
	}

	return nil
}

// Lookup returns the Extension for extensionName.
func (r *Registry[T]) Lookup(extensionName string) (Extension[T], bool) {
	if r == nil {
		return Extension[T]{}, false
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	foundExtension, ok := r.extensions[extensionName]
	if !ok {
		return Extension[T]{}, false
	}

	return foundExtension, true
}
