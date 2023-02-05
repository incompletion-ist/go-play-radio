package alteration

import (
	"reflect"

	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// swapFieldValue returns a new reflect.Value for a new struct value by swapping out the
// original struct's field named fieldName with the type and value of newValue.
//
// The new struct will be populated with the values from origStruct for all other exported
// fields.
func swapFieldValue(origStructV reflect.Value, fieldName string, newValueV reflect.Value) (reflect.Value, error) {
	if !origStructV.IsValid() {
		return reflect.Value{}, wraperr.Errorf(ErrorInvalidValue, "alteration: origStruct is invalid")
	}

	if fieldName == "" {
		return reflect.Value{}, wraperr.Errorf(ErrorInvalidField, "alteration: fieldName is empty")
	}

	if !newValueV.IsValid() {
		return reflect.Value{}, wraperr.Errorf(ErrorInvalidValue, "alteration: newValue is invalid")
	}

	origStructT := origStructV.Type()
	newValueT := newValueV.Type()

	newStructT, newStructTErr := swapFieldType(origStructT, fieldName, newValueT)
	if newStructTErr != nil {
		return reflect.Value{}, newStructTErr
	}

	newStructPtrV := reflect.New(newStructT)
	newStructV := newStructPtrV.Elem()

	for fieldNum := 0; fieldNum < origStructV.NumField(); fieldNum++ {
		fieldT := origStructT.Field(fieldNum)
		if !fieldT.IsExported() {
			continue
		}

		origFieldV := origStructV.Field(fieldNum)
		newFieldV := newStructV.Field(fieldNum)

		if fieldT.Name == fieldName {
			newFieldV.Set(newValueV)
		} else {
			newFieldV.Set(origFieldV)
		}
	}

	return newStructV, nil
}
