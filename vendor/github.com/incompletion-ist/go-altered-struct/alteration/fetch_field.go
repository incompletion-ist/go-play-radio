package alteration

import (
	"reflect"

	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// FetchField returns structValue's field value as FieldTypeT. It will return an error if
// the field is not of type FieldTypeT.
//
// The returned value will be the zero value of FieldTypeT if an error is returned.
func FetchField[FieldTypeT any](structValue any, fieldName string) (FieldTypeT, error) {
	var zeroFieldValue FieldTypeT

	structValueV := reflect.ValueOf(structValue)
	if !structValueV.IsValid() {
		return zeroFieldValue, wraperr.Errorf(ErrorInvalidValue, "alteration: structValue is invalid")
	}

	if structValueV.Kind() != reflect.Struct {
		return zeroFieldValue, wraperr.Errorf(ErrorInvalidType, "alteration: structValueV is not a struct")
	}

	if fieldName == "" {
		return zeroFieldValue, wraperr.Errorf(ErrorInvalidField, "alteration: fieldName is empty")
	}

	fieldValueV := structValueV.FieldByName(fieldName)
	fieldValueI := fieldValueV.Interface()

	fieldValueTyped, fieldValueTypedOk := fieldValueI.(FieldTypeT)
	if !fieldValueTypedOk {
		return zeroFieldValue, wraperr.Errorf(ErrorInvalidType, "alteration: field %q is not of type %T", fieldName, zeroFieldValue)
	}

	return fieldValueTyped, nil
}
