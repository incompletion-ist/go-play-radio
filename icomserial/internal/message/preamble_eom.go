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

package message

import "go.incompletion.ist/play-radio/errorwrap"

const (
	preamble byte = 0xFE
	eom      byte = 0xFD
)

func expectSinglePreamble() bytesHandler {
	return func(bytes []byte) (int, error) {
		if len(bytes) < 1 || bytes[0] != preamble {
			return 0, errorwrap.WrapError(ErrorParsing, nil, "message: missing preamble")
		}

		return 1, nil
	}
}

func expectFullPreamble() bytesHandler {
	return func(bytes []byte) (int, error) {
		return handleBytes(
			bytes,
			expectSinglePreamble(),
			expectSinglePreamble(),
		)
	}
}

func expectEom() bytesHandler {
	return func(bytes []byte) (int, error) {
		if len(bytes) < 1 || bytes[0] != eom {
			return 0, errorwrap.WrapError(ErrorParsing, nil, "message: missing eom")
		}

		return 1, nil
	}
}
