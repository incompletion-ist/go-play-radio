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
	"testing"

	"go.incompletion.ist/play-radio/operating"
	"go.incompletion.ist/play-radio/transceiver"
)

func Test_handleFrequency(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		wantCount int
		wantError bool
		wantValue operating.Frequency
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
			name:      "too short",
			input:     []byte{0x10, 0x32, 0x54, 0x76},
			wantError: true,
		},
		{
			name:      "just right",
			input:     []byte{0x10, 0x32, 0x54, 0x76, 0x98},
			wantCount: 5,
			wantValue: 9876543210,
		},
		{
			name:      "more than enough",
			input:     []byte{0x10, 0x32, 0x54, 0x76, 0x98, 0x10},
			wantCount: 5,
			wantValue: 9876543210,
		},
	}

	for _, test := range tests {
		var conf transceiver.Configuration
		gotCount, err := handleFrequency(&conf)(test.input)
		gotError := err != nil
		gotValue := conf.Frequency.Get()

		if gotError != test.wantError {
			t.Errorf("%s: handleFrequency returned error? %v (%s)", test.name, gotError, err)
		}

		if gotCount != test.wantCount {
			t.Errorf("%s: handleFrequency returned %d, want %d", test.name, gotCount, test.wantCount)
		}

		if gotValue != test.wantValue {
			t.Errorf("%s: handleFrequency set %d, want %d", test.name, gotValue, test.wantValue)
		}
	}
}
