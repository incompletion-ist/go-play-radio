// Copyright 2022 Micah Kemp
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errorwrap

import (
	"fmt"
	"reflect"
)

// Code defines the underlying types that can be used as for wrapped error codes.
type Code interface {
	~int
}

// Coder is a convenience interface to simplify checking if a returned error is
// errorwrap.WrappedError.
type Coder interface {
	IsCode(interface{}) bool
}

// WrappedError represents a wrapped error with a related error code of the given type.
type WrappedError[T Code] struct {
	Original error

	code    T
	message string
}

// Error returns the error message.
func (wrapped WrappedError[T]) Error() string {
	return wrapped.message
}

// IsCode returns true if the given code is deeply equal to the wrapped error's code.
// It is only deeply equal if it of the same named type and value, permitting differention
// between different types of wrapped errors, each with their own custom code type.
func (wrapped WrappedError[T]) IsCode(code interface{}) bool {
	return reflect.DeepEqual(code, wrapped.code)
}

// WrapError returns a new WrappedError.
func WrapError[T Code](code T, err error, messageF string, messageArgs ...interface{}) error {
	return WrappedError[T]{
		code:     code,
		Original: err,
		message:  fmt.Sprintf(messageF, messageArgs...),
	}
}
