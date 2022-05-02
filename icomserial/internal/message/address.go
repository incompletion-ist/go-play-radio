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
	defaultTransceiverAddress byte = 0xA2
	defaultControllerAddress  byte = 0xE0
)

func handleAddress(address *byte) bytesHandler {
	return func(data []byte) (int, error) {
		if len(data) < 1 {
			return 0, errorwrap.WrapError(ErrorParsing, nil, "message: expected address byte")
		}

		*address = data[0]

		return 1, nil
	}
}
