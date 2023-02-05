package alteration

import (
	"reflect"

	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// swapFieldType returns a reflect.Type for a new struct type that swaps the field type for fieldName with
// the given newFieldType.
//
// The returned reflect.Type will be empty if an error is returned.
func swapFieldType(origStructType reflect.Type, fieldName string, newFieldType reflect.Type) (reflect.Type, error) {
	if origStructType == nil {
		return nil, wraperr.Errorf(ErrorInvalidType, "alteration: origStructType is nil")
	}

	if origStructType.Kind() != reflect.Struct {
		return nil, wraperr.Errorf(ErrorInvalidType, "alteration: origStruct is not a struct")
	}

	if fieldName == "" {
		return nil, wraperr.Errorf(ErrorInvalidField, "alteration: fieldName is empty")
	}

	if newFieldType == nil {
		return nil, wraperr.Errorf(ErrorInvalidType, "alteration: newFieldType is nil")
	}

	newStructFields := make([]reflect.StructField, origStructType.NumField())
	foundField := false
	for fieldNum := range newStructFields {
		fieldT := origStructType.Field(fieldNum)
		if fieldT.Name == fieldName {
			fieldT.Type = newFieldType
			foundField = true
		}

		newStructFields[fieldNum] = fieldT
	}

	if !foundField {
		return nil, wraperr.Errorf(ErrorInvalidField, "alteration: field %q not found", fieldName)
	}

	return reflect.StructOf(newStructFields), nil
}
