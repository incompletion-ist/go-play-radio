package driver

import (
	"plugin"

	"github.com/incompletion-ist/go-extension/extension"
	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// Driver defines the standard for providing and consuming driver extensions.
type Driver = extension.Extension[OpenCloser]

// Load opens file as a Go Plugin and returns the Driver returned by its NewDriver function.
//
// The plugin's NewDriver function must have this signature:
//
//	NewDriver() func() driver.Driver
func Load(file string) (Driver, error) {
	var zeroDriver Driver

	p, pErr := plugin.Open(file)
	if pErr != nil {
		return zeroDriver, wraperr.Wrapf(
			ErrorPlugin,
			pErr,
			"driver: unable to open Go plugin: %s",
			pErr,
		)
	}

	newFuncSymbol, newFuncSymbolErr := p.Lookup("NewDriver")
	if newFuncSymbolErr != nil {
		return zeroDriver, wraperr.Wrapf(
			ErrorPlugin,
			pErr,
			"driver: unable to look up NewDriver: %s",
			newFuncSymbolErr,
		)
	}

	newFunc, newFuncOk := newFuncSymbol.(func() Driver)
	if !newFuncOk {
		return zeroDriver, wraperr.Errorf(
			ErrorPlugin,
			`driver: NewDriver is not a "func() Driver"`,
		)
	}

	return newFunc(), nil
}
