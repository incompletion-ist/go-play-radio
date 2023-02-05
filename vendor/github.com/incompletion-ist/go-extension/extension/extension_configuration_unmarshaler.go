package extension

import (
	"bytes"
	"encoding/json"
)

// ExtensionConfigurationUnmarshaler implements custom JSON and YAML unmarshaling directly
// into Content, with ExtensionConfiguration members having access to the configured Registry.
type ExtensionConfigurationUnmarshaler[ExtensionT any, ContentT any] struct {
	Registry *Registry[ExtensionT]
	Content  ContentT
}

// unmarshal performs unmarshaling directly into Content, within a withPlugins context.
func (unmarshaler *ExtensionConfigurationUnmarshaler[ExtensionT, ContentT]) unmarshal(
	unmarshalFunc func(any) error,
) error {
	var newContent ContentT

	var unmarshalErr error
	withAvailableRegistry(unmarshaler.Registry, func() {
		unmarshalErr = unmarshalFunc(&newContent)
	})
	if unmarshalErr != nil {
		return unmarshalErr
	}

	unmarshaler.Content = newContent

	return nil
}

// UnmarshalJSON implements custom JSON unmarshaling.
func (unamrshaler *ExtensionConfigurationUnmarshaler[ExtensionT, ContentT]) UnmarshalJSON(
	data []byte,
) error {
	return unamrshaler.unmarshal(func(value any) error {
		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.DisallowUnknownFields()

		return decoder.Decode(value)
	})
}

// UnmarshalYAML implements custom YAML unmarshaling.
func (unmarshaler *ExtensionConfigurationUnmarshaler[ExtensionT, ContentT]) UnmarshalYAML(
	unmarshalFunc func(any) error,
) error {
	return unmarshaler.unmarshal(unmarshalFunc)
}
