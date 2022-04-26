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

package errorwrap_test

import (
	"fmt"

	"go.incompletion.ist/play-radio/errorwrap"
)

type ErrorCode int

type Error errorwrap.WrappedError[ErrorCode]

type ExternalErrorCode int

type ExternalError errorwrap.WrappedError[ExternalErrorCode]

const (
	ErrorUndefined        ErrorCode = iota // = 0
	ErrorUnhandledCommand                  // = 1
)

const (
	ErrorExternalUndefined ExternalErrorCode = iota // = 0
	ErrorExternalFileError                          // = 1
)

func Example_wrappedError() {
	err := errorwrap.WrapError(ErrorUnhandledCommand, nil, "unhandled command: %X", 0xFF)

	if codedError, ok := err.(errorwrap.Coder); ok {
		if codedError.IsCode(ErrorUnhandledCommand) {
			fmt.Printf("encountered unhandled command (%s)\n", err)
		}

		// ErrorUnhandledCommand and ErrorExternalFileError both equal 1, but are of differnet types,
		// so IsCode(ErrorExternalFileError) returns false
		if codedError.IsCode(ErrorExternalFileError) {
			fmt.Printf("external file error: %s\n", err)
		}
	}
	// Output: encountered unhandled command (unhandled command: FF)
}
