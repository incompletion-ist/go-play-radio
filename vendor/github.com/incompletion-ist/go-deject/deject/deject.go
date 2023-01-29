// Package deject enables dependency injection.
package deject

// InjectDependency injects a dependency of type DependencyT with the given
// value into Registry.
func InjectDependency[DependencyT any](registry *Registry, value DependencyT) error {
	newDependency, newDependencyErr := dependencyOf(value)
	if newDependencyErr != nil {
		return newDependencyErr
	}

	return registry.injectDependency(newDependency)
}

// FetchDependency fetches the dependency of type DependencyT from the Registry.
// Returns false if the given type is not present in the Registry.
func FetchDependency[DependencyT any](registry *Registry) (DependencyT, bool) {
	dependencyType := typeOf[DependencyT]()

	dep, found := registry.fetchDependencyType(dependencyType)
	if found {
		return dep.reflectValue.Interface().(DependencyT), true
	}

	var dependencyZeroValue DependencyT
	return dependencyZeroValue, false
}

// InjectRegistries injects each dependency that exists in each of the given
// Registries.
func (registry *Registry) InjectRegistries(registries ...*Registry) error {
	for _, extraRegistry := range registries {
		if injectErr := registry.injectRegistry(extraRegistry); injectErr != nil {
			return injectErr
		}
	}

	return nil
}

// MergeRegistries returns a new Registry consisting of all dependencies
// present in each of the given Registries.
func MergeRegistries(registries ...*Registry) (*Registry, error) {
	newRegistry := &Registry{}

	if mergeErr := newRegistry.InjectRegistries(registries...); mergeErr != nil {
		return nil, mergeErr
	}

	return newRegistry, nil
}
