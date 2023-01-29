package deject

import "reflect"

// typeOf returns the reflection Type for a generic type.
func typeOf[T any]() reflect.Type {
	valueT := reflect.TypeOf((*T)(nil)).Elem()

	return valueT
}
