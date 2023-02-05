package alteration

import (
	"reflect"
)

// SwapFieldForValue returns a new struct value by swapping out the original struct's
// field named fieldName with the type and value of newValue.
//
// The new struct will be populated with the values from origStruct for all other exported
// fields.
//
// The returned value will be nil if an error is returned.
func SwapFieldForValue(origStruct any, fieldName string, newValue any) (any, error) {
	origStructV := reflect.ValueOf(origStruct)
	newValueV := reflect.ValueOf(newValue)

	newStructV, newStructErr := swapFieldValue(origStructV, fieldName, newValueV)
	if newStructErr != nil {
		return nil, newStructErr
	}

	return newStructV.Interface(), nil
}
