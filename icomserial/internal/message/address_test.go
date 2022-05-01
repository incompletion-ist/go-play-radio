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

func Test_handleAddress(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		wantCount   int
		wantError   bool
		wantAddress byte
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
			name:        "single byte for address",
			input:       []byte{0x01},
			wantCount:   1,
			wantAddress: 0x01,
		},
		{
			name:        "address and more",
			input:       []byte{0x01, 0x02, 0x03, 0x04},
			wantCount:   1,
			wantAddress: 0x01,
		},
	}

	for _, test := range tests {
		var gotAddress byte
		gotCount, err := handleAddress(&gotAddress)(test.input)
		gotError := err != nil

		if gotError != test.wantError {
			t.Errorf("%s: handleAddress returned error? %v (%s)", test.name, gotError, err)
		}

		if gotCount != test.wantCount {
			t.Errorf("%s: handleAddress got %d, want %d", test.name, gotCount, test.wantCount)
		}

		if gotAddress != test.wantAddress {
			t.Errorf("%s: handleAddress set address 0x%X, want 0x%X", test.name, gotAddress, test.wantAddress)
		}
	}
}
