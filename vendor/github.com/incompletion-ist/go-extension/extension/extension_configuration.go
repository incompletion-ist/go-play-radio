package extension

import (
	"bytes"
	"encoding/json"

	"github.com/incompletion-ist/go-altered-struct/alteration"
	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// ExtensionConfiguration associates an instance of ExtensionT with a named Extension.
//
// It implements custom JSON and YAML unmarshaling into a concrete instance of ExtensionT
// by consulting a Registry. To accomplish this, it must be embedded at some level within
// an ExtenstionConfigurationUnmarshaler.
type ExtensionConfiguration[ExtensionT any] struct {
	ExtensionName string     `json:"extension" yaml:"extension"`
	Extension     ExtensionT `json:"configuration" yaml:"configuration"`
}

// unmarshal performs custom unmarshaling.
func (configuration *ExtensionConfiguration[ExtensionT]) unmarshal(unmarshalFunc alteration.UnmarshalFunc) error {
	if availableRegistry == nil {
		return wraperr.Errorf(ErrorUnmarshal, "extension: attempted unmarshaling ExtensionConfiguration outside of ExtensionConfigurationUnmarshaler")
	}

	registry, registryOk := availableRegistry.(*Registry[ExtensionT])
	if !registryOk {
		return wraperr.Errorf(ErrorUnmarshal, "extension: attempted unmarshaling ExtensionConfiguration with incorrect Registry type, this indicates a bug in the extension package")
	}

	extensionName, extensionNameErr := alteration.AlterFetch[string](
		*configuration,
		alteration.AlterUnmarshal(
			alteration.AlterSwapFieldForZero[any]("Extension"),
			unmarshalFunc,
		),
		"ExtensionName",
	)
	if extensionNameErr != nil {
		return wraperr.Wrapf(ErrorUnmarshal, extensionNameErr, "extension: error unmarshaling for ExtensionName: %s", extensionNameErr)
	}

	foundExtension, foundExtensionOk := registry.Lookup(extensionName)
	if !foundExtensionOk {
		return wraperr.Errorf(ErrorUnmarshal, "extension: extension not found with name %q", extensionName)
	}

	newExtension := foundExtension.NewFunc()
	unmarshaledExtension, unmarshaledExtensionErr := alteration.AlterFetch[ExtensionT](
		*configuration,
		alteration.AlterUnmarshal(
			alteration.AlterSwapFieldForValue("Extension", newExtension),
			unmarshalFunc,
		),
		"Extension",
	)
	if unmarshaledExtensionErr != nil {
		return wraperr.Wrapf(ErrorUnmarshal, unmarshaledExtensionErr, "extension: error unmarshaling for Extension: %s", unmarshaledExtensionErr)
	}

	configuration.ExtensionName = extensionName
	configuration.Extension = unmarshaledExtension

	return nil
}

// UnmarshalJSON implements custom JSON unmarshaling.
func (configuration *ExtensionConfiguration[ExtensionT]) UnmarshalJSON(data []byte) error {
	unmarshalFunc := func(value any) error {
		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.DisallowUnknownFields()

		return decoder.Decode(value)
	}

	return configuration.unmarshal(unmarshalFunc)
}

// UnmarshalYAML implements custom YAML unmarshaling.
func (configuration *ExtensionConfiguration[ExtensionT]) UnmarshalYAML(unmarshalFunc func(any) error) error {
	return configuration.unmarshal(unmarshalFunc)
}
