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

import "testing"

func Test_skipToNextCommand(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		wantCount int
		wantError bool
	}{
		{
			name:      "nil",
			wantError: true,
		},
		{
			name:      "empty",
			input:     []byte{},
			wantError: true,
		},
		{
			name:      "non-eom only",
			input:     []byte{0x01},
			wantError: true,
		},
		{
			name:      "eom only",
			input:     []byte{eom},
			wantCount: 1,
		},
		{
			name:      "arbitrary bytes before eom",
			input:     []byte{0x01, 0x02, 0x03, 0x04, eom},
			wantCount: 5,
		},
		{
			name: "non-preamble bytes after eom",
			// eom in this input is "just another byte" because it isn't
			// followed by preamble,preamble.
			input:     []byte{0x01, 0x02, 0x03, 0x04, eom, 0x01},
			wantCount: 5,
			wantError: true,
		},
		{
			name: "single preamble byte after eom",
			// eom in this input is "just another byte" because it isn't
			// followed by preamble,preamble.
			input:     []byte{0x01, 0x02, 0x03, 0x04, eom, preamble},
			wantCount: 5,
			wantError: true,
		},
		{
			name:      "preamble bytes after eom",
			input:     []byte{0x01, 0x02, 0x03, 0x04, eom, preamble, preamble},
			wantCount: 5,
		},
		{
			name:      "eom in middle of message",
			input:     []byte{0x01, 0x02, 0x03, 0x04, eom, 0x06, 0x07, 0x08, eom, preamble, preamble},
			wantCount: 9,
		},
	}

	for _, test := range tests {
		gotCount, err := skipToNextCommand()(test.input)
		gotError := err != nil

		if gotError != test.wantError {
			t.Errorf("%s: skipToNextCommand returned error? %v (%s)", test.name, gotError, err)
		}

		if gotCount != test.wantCount {
			t.Errorf("%s: skipToNextCommand got %d, want %d", test.name, gotCount, test.wantCount)
		}
	}
}
