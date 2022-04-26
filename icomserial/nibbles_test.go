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
	"reflect"
	"testing"
)

func Test_getNibbles(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name: "nil",
			want: []byte{},
		},
		{
			name:  "empty",
			input: []byte{},
			want:  []byte{},
		},
		{
			name:  "single byte",
			input: []byte{0x12},
			want:  []byte{0x02, 0x01},
		},
		{
			name:  "multiple bytes",
			input: []byte{0x12, 0x34, 0x56, 0x78, 0x90},
			want:  []byte{0x02, 0x01, 0x04, 0x03, 0x06, 0x05, 0x08, 0x07, 0x00, 0x09},
		},
	}

	for _, test := range tests {
		got := getNibbles(test.input)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s: getNibbles() got\n%#v, want\n%#v", test.name, got, test.want)
		}
	}
}

func Test_dataAsInt(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int64
	}{
		{
			name: "nil",
		},
		{
			name: "empty",
		},
		{
			name:  "single byte",
			input: []byte{0x12},
			want:  12,
		},
		{
			name:  "multiple bytes",
			input: []byte{0x21, 0x43, 0x65, 0x87, 0x09},
			want:  987654321,
		},
	}

	for _, test := range tests {
		got := dataAsInt(test.input)

		if got != test.want {
			t.Errorf("%s: dataAsInt() got %d, want %d", test.name, got, test.want)
		}
	}
}
