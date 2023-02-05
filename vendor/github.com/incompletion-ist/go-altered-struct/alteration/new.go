package alteration

import (
	"reflect"

	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// New returns a value that is a pointer to a copy of origValue. It is useful
// for getting values that may be directly passed to unmarshaling libraries.
func New(origValue any) (any, error) {
	origValueV := reflect.ValueOf(origValue)
	if !origValueV.IsValid() {
		return nil, wraperr.Errorf(ErrorInvalidValue, "alteration: origValue is invalid")
	}

	origValueT := origValueV.Type()

	newValuePtrV := reflect.New(origValueT)
	newValueV := newValuePtrV.Elem()
	newValueV.Set(origValueV)

	return newValuePtrV.Interface(), nil
}
