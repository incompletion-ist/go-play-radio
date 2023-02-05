package extension

import "sync"

var availableRegistry any
var availableRegistryMu sync.Mutex

func setAvailableRegistry[ExtensionT any](registry *Registry[ExtensionT]) {
	availableRegistry = registry
}

func clearAvailableRegistry() {
	availableRegistry = nil
}

// withAvailableRegistry calls fn after setting availableRegistry to the given Registry.
// Upon fn's completion, availableRegistry is cleared.
//
// availableRegistryMu is locked during during fn's execution, so care must be taken
// not to call withAvailableRegistry in a nested manner.
func withAvailableRegistry[ExtensionT any](registry *Registry[ExtensionT], fn func()) {
	availableRegistryMu.Lock()
	defer availableRegistryMu.Unlock()
	defer clearAvailableRegistry()

	setAvailableRegistry(registry)

	fn()
}
