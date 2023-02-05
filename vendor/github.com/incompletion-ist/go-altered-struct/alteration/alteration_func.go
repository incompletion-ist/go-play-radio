package alteration

import (
	"github.com/incompletion-ist/go-wrapped-error/wraperr"
)

// AlterationFunc is a function that returns an altered value.
type AlterationFunc func(any) (any, error)

// AlterSwapFieldForZero returns an AlterationFunc that swaps fieldName
// with the zero value for FieldT.
func AlterSwapFieldForZero[FieldT any](fieldName string) AlterationFunc {
	return func(value any) (any, error) {
		return SwapFieldForZero[FieldT](value, fieldName)
	}
}

// AlterInPlace returns an AlterationFunc that calls inPlaceFunc and returns
// the value it was passed as a parameter.
func AlterInPlace(inPlaceFunc func(any) error) AlterationFunc {
	return func(value any) (any, error) {
		fnErr := inPlaceFunc(value)
		if fnErr != nil {
			return nil, wraperr.Wrapf(ErrorFunctionalParam, fnErr, "alteration: inPlaceFunc returned error: %s", fnErr)
		}

		return value, nil
	}
}

// AlterSwapFieldForValue returns an AlterationFunc that swaps fieldName
// with the type and value of newValue.
func AlterSwapFieldForValue(fieldName string, newValue any) AlterationFunc {
	return func(value any) (any, error) {
		return SwapFieldForValue(value, fieldName, newValue)
	}
}

// ComposeAlterationFunc returns a new AlterationFunc that is composed of multiple
// other AlterationFuncs.
func ComposeAlterationFunc(alterationFuncs ...AlterationFunc) AlterationFunc {
	return func(value any) (any, error) {
		for _, fn := range alterationFuncs {
			alteredValue, alteredValueErr := fn(value)
			if alteredValueErr != nil {
				return nil, alteredValueErr
			}

			value = alteredValue
		}

		return value, nil
	}
}

// AlterUnmarshal returns an AlterationFunc that performs commonly needed steps to unmarshal
// an altered struct:
//
// * calls alterationFunc (preparation before unmarshaling occurrs here)
//
// * calls unmarshalFunc (on pointer to the altered struct)
//
// * returns the unmarshaled modified struct
func AlterUnmarshal(alterationFunc AlterationFunc, unmarshalFunc UnmarshalFunc) AlterationFunc {
	return ComposeAlterationFunc(
		alterationFunc,
		New,
		AlterInPlace(unmarshalFunc),
		Elem,
	)
}

// AlterFetch returns the value for fieldName after performing alterationFunc on value.
func AlterFetch[FieldT any](value any, alterationFunc AlterationFunc, fieldName string) (FieldT, error) {
	var fieldZero FieldT

	altered, alteredErr := alterationFunc(value)
	if alteredErr != nil {
		return fieldZero, alteredErr
	}

	return FetchField[FieldT](altered, fieldName)
}
