package alteration

import (
	"reflect"

	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// Elem returns the value pointed to by valuePtr. It is useful to obtain the non-pointer
// value after unmarshaling into the value from New().
func Elem(valuePtr any) (any, error) {
	valuePtrV := reflect.ValueOf(valuePtr)
	if !valuePtrV.IsValid() {
		return nil, wraperr.Errorf(ErrorInvalidValue, "alteration: ptrValue is invalid")
	}

	if valuePtrV.Kind() != reflect.Pointer {
		return nil, wraperr.Errorf(ErrorInvalidValue, "alteration: ptrValue is not a pointer")
	}

	valueV := valuePtrV.Elem()

	return valueV.Interface(), nil
}
