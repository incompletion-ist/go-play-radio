package alteration

import "reflect"

// SwapFieldForZero returns a new struct value by swapping out the original struct's
// field named fieldName with the type and zero value for FieldTypeT.
//
// The new struct will be populated with the values from origStruct for all other exported
// fields.
//
// The returned value will be nil if an error is returned.
func SwapFieldForZero[FieldTypeT any](origStruct any, fieldName string) (any, error) {
	origStructV := reflect.ValueOf(origStruct)

	zeroFieldValueT := reflect.TypeOf((*FieldTypeT)(nil)).Elem()
	zeroFieldValueV := reflect.Zero(zeroFieldValueT)

	newStructV, newStructErr := swapFieldValue(origStructV, fieldName, zeroFieldValueV)
	if newStructErr != nil {
		return nil, newStructErr
	}

	return newStructV.Interface(), nil
}
