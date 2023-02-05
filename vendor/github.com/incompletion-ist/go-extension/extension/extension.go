// Package extension provides generic types for implementing named
// extensions.
package extension

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// NewFunc is a function that returns a new instance of type T.
type NewFunc[T any] func() T

// Extension is a named extension for type T.
type Extension[T any] struct {
	Name    string
	NewFunc NewFunc[T]
}

// validate returns an error if Extension is invalid.
func (e Extension[T]) validate() error {
	if e.Name == "" {
		return wraperr.Errorf(ErrorInvalid, "extension: Extension name is invalid: %q", e.Name)
	}

	if e.NewFunc == nil {
		return wraperr.Errorf(ErrorInvalid, "extension: Extension has nil NewFunc")
	}

	return nil
}
