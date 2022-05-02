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

func Test_handleAllMessages(t *testing.T) {
	tests := []struct {
		name          string
		input         []byte
		wantFrequency operating.Frequency
		wantError     bool
	}{
		{
			name: "nil",
		},
		{
			name:  "empty",
			input: []byte{},
		},
		{
			name:          "set frequency once",
			input:         []byte{preamble, preamble, defaultTransceiverAddress, defaultControllerAddress, sendFrequency, 0x10, 0x32, 0x54, 0x76, 0x98, eom},
			wantFrequency: 9876543210,
		},
		{
			name: "set frequency twice",
			input: []byte{
				preamble, preamble, defaultTransceiverAddress, defaultControllerAddress, sendFrequency, 0x10, 0x32, 0x54, 0x76, 0x98, eom,
				preamble, preamble, defaultTransceiverAddress, defaultControllerAddress, sendFrequency, 0x89, 0x67, 0x45, 0x23, 0x01, eom,
			},
			wantFrequency: 123456789,
		},
		{
			name: "set frequency before bad command",
			input: []byte{
				preamble, preamble, defaultTransceiverAddress, defaultControllerAddress, sendFrequency, 0x10, 0x32, 0x54, 0x76, 0x98, eom,
				eom,
			},
			wantFrequency: 9876543210,
			wantError:     true,
		},
		{
			name: "set frequency after bad command",
			input: []byte{
				eom,
				preamble, preamble, defaultTransceiverAddress, defaultControllerAddress, sendFrequency, 0x10, 0x32, 0x54, 0x76, 0x98, eom,
			},
			wantError: true,
		},
	}

	for _, test := range tests {
		var conf transceiver.Configuration
		err := handleAllMessages(test.input, &conf)
		gotError := err != nil
		gotFrequency := conf.Frequency.Get()

		if gotError != test.wantError {
			t.Errorf("%s: handleAllMessages returned error? %v (%s)", test.name, gotError, err)
		}

		if gotFrequency != test.wantFrequency {
			t.Errorf("%s: handleAllMessages set frequency: %d, want %d", test.name, gotFrequency, test.wantFrequency)
		}
	}
}
