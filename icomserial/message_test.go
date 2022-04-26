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

func Test_parseSerialData(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		want      []messageBytes
		wantError bool
	}{
		{
			name: "nil",
		},
		{
			name:  "empty",
			input: []byte{},
		},
		{
			name:      "bad preamble",
			input:     []byte{0x01},
			wantError: true,
		},
		{
			name:      "preamble only",
			input:     []byte{preamble, preamble},
			wantError: true,
		},
		{
			name:  "preamble and eom only",
			input: []byte{preamble, preamble, eom},
			want: []messageBytes{
				// explicitly empty
				{},
			},
		},
		{
			name:  "valid single message",
			input: []byte{preamble, preamble, 0x01, eom},
			want: []messageBytes{
				{0x01},
			},
		},
		{
			name: "valid multiple messages",
			input: []byte{
				preamble, preamble, 0x01, eom,
				preamble, preamble, 0x02, eom,
			},
			want: []messageBytes{
				{0x01},
				{0x02},
			},
		},
		{
			name: "valid multiple messages, eom in data",
			input: []byte{
				preamble, preamble, 0x01, eom, eom,
				preamble, preamble, 0x02, eom,
			},
			want: []messageBytes{
				{0x01, eom},
				{0x02},
			},
		},
		{
			name: "invalid multiple messages",
			input: []byte{
				preamble, preamble, 0x01, eom,
				preamble, preamble, 0x02,
			},
			wantError: true,
		},
	}

	for _, test := range tests {
		got, err := parseSerialData(test.input)
		gotError := err != nil

		if gotError != test.wantError {
			t.Errorf("%s: parseSerialData() returned error? %v (%s)", test.name, gotError, err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s: parseSerialData() got\n%#v, want\n%#v", test.name, got, test.want)
		}
	}
}
