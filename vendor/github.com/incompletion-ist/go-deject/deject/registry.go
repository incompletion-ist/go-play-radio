package deject

import (
	"reflect"
	"sync"

	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// Registry stores and fetches statically typed dependencies.
type Registry struct {
	dependencies []dependency
	mu           sync.Mutex
}

// fetchDependencyTypeUnsafe fetches a stored dependency. It does
// not check for nil Registry or lock before making changes. This
// must be performed by the caller.
func (registry *Registry) fetchDependencyTypeUnsafe(dependencyType reflect.Type) (dependency, bool) {
	for _, dep := range registry.dependencies {
		if dep.reflectType == dependencyType {
			return dep, true
		}
	}

	return dependency{}, false
}

// fetchDependencyType fetches a dependency of the given type.
func (registry *Registry) fetchDependencyType(dependencyType reflect.Type) (dependency, bool) {
	if registry == nil {
		return dependency{}, false
	}

	registry.mu.Lock()
	defer registry.mu.Unlock()

	return registry.fetchDependencyTypeUnsafe(dependencyType)
}

// injectDependency injects a dependency.
func (registry *Registry) injectDependency(dep dependency) error {
	if registry == nil {
		return wraperr.Errorf(ErrorNilValue, "deject: Registry is nil")
	}

	if !dep.isValid() {
		return wraperr.Errorf(ErrorInvalid, "deject: dependency is not valid")
	}

	registry.mu.Lock()
	defer registry.mu.Unlock()

	if _, found := registry.fetchDependencyTypeUnsafe(dep.reflectType); found {
		return wraperr.Errorf(ErrorDuplicate, "deject: dependency type %q already in Registry", dep.reflectType)
	}

	registry.dependencies = append(registry.dependencies, dep)

	return nil
}

// injectRegistry injects each stored dependency in extraRegistry.
func (registry *Registry) injectRegistry(extraRegistry *Registry) error {
	if extraRegistry == nil {
		return wraperr.Errorf(ErrorNilValue, "deject: Registry is nil")
	}

	for _, dep := range extraRegistry.dependencies {
		if injectErr := registry.injectDependency(dep); injectErr != nil {
			return injectErr
		}
	}

	return nil
}
