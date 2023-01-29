package deject

import (
	"reflect"

	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// dependency is an injectable dependency.
type dependency struct {
	reflectType  reflect.Type
	reflectValue reflect.Value
}

// dependencyOf returns a Dependency for the given value, which must be non-nil.
func dependencyOf[DependencyT any](value DependencyT) (dependency, error) {
	dependencyType := typeOf[DependencyT]()

	if dependencyType.Kind() != reflect.Interface {
		return dependency{}, wraperr.Errorf(ErrorInvalid, "deject: invalid dependency value, must be an interface")
	}

	dependencyValue := reflect.ValueOf(value)
	if !dependencyValue.IsValid() {
		return dependency{}, wraperr.Errorf(ErrorNilValue, "deject: invalid dependency value, must be non-nil")
	}

	return dependency{
		reflectType:  dependencyType,
		reflectValue: dependencyValue,
	}, nil
}

// mustDependencyOf returns a Dependency for the given value. It panics
// if the requirements of DependencyOf are not met.
func mustDependencyOf[T any](value T) dependency {
	d, err := dependencyOf(value)
	if err != nil {
		panic(err)
	}

	return d
}

// isValid returns false if the Dependency has not been properly instantiated.
func (d dependency) isValid() bool {
	return d.reflectValue.IsValid()
}
