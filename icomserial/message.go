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

package icomserial

import (
	"go.incompletion.ist/play-radio/errorwrap"
)

type messageBytes []byte

const (
	preamble byte = 0xFE
	eom      byte = 0xFD
)

// parseSerialData returns a slice of messageBytes for the given input data.
// An error is returned if all of the data can't be parsed into valid messageBytes.
func parseSerialData(data []byte) ([]messageBytes, error) {
	if len(data) == 0 {
		return nil, nil
	}

	if len(data) < 2 {
		return nil, errorwrap.WrapError(ErrorCommandParsing, nil, "icomserial: incomplete preable %v", data)
	}

	if !(data[0] == preamble && data[1] == preamble) {
		return nil, errorwrap.WrapError(ErrorCommandParsing, nil, "icomserial: invalid preamble (0x%X 0x%X)", data[0], data[1])
	}
	data = data[2:]

	msgBytes := messageBytes{}
	for i, msgByte := range data {
		if msgByte == eom {
			remainingMessagesBytes, err := parseSerialData(data[i+1:])
			if err == nil {
				messagesBytes := []messageBytes{msgBytes}
				messagesBytes = append(messagesBytes, remainingMessagesBytes...)
				return messagesBytes, nil
			}
		}

		msgBytes = append(msgBytes, msgByte)
	}

	return nil, errorwrap.WrapError(ErrorCommandParsing, nil, "icomserial: eom not found")
}
