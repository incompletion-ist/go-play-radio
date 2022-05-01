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

import (
	"math"

	"go.incompletion.ist/play-radio/errorwrap"
)

// getNibbles returns a slice of bytes from dataBytes, one returned byte
// for each nibble in dataBytes, with the least significant nibble of each
// byte returned first.
//
// For example, these dataBytes:
//
// {0x10, 0x32 0x54}
//
// Will be returned as:
//
// {0x00, 0x01, 0x02, 0x03, 0x04, 0x05}
func getNibbles(dataBytes []byte) []byte {
	nibbleBytes := make([]byte, 2*len(dataBytes))

	for i, dataByte := range dataBytes {
		upperByte := dataByte >> 4
		lowerByte := dataByte & 0x0F

		nibbleBytes[i*2+0] = lowerByte
		nibbleBytes[i*2+1] = upperByte
	}

	return nibbleBytes
}

// dataAsInt returns an int by treating each input byte as a pair of
// most signifant and least significant power of 10 digits, and each
// subsequent byte as the next powers of 10.
//
// For example, this input data:
//
// {0x21, 0x43, 0x65, 0x87}
//
// Is returned as this int:
//
// 87654321
func dataAsInt(data []byte) int64 {
	var intValue int64

	for i, dataNibble := range getNibbles(data) {
		intValue = intValue + int64(dataNibble)*int64(math.Pow10(i))
	}

	return intValue
}

func handleInt(dataLen int, value *int64) bytesHandler {
	return func(data []byte) (int, error) {
		if dataLen < 1 {
			return 0, errorwrap.WrapError(ErrorParsing, nil, "message: unable to handle int, dataLen must be greater than zero")
		}

		if len(data) < dataLen {
			return 0, errorwrap.WrapError(ErrorParsing, nil, "message: unable to handle int, need data length %d, have %d", dataLen, len(data))
		}

		*value = dataAsInt(data[0:dataLen])

		return dataLen, nil
	}
}
