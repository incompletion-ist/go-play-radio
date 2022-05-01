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

func Test_expectSinglePreamble(t *testing.T) {
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
			name:      "not preamble",
			input:     []byte{0x01},
			wantError: true,
		},
		{
			name:      "preamble",
			input:     []byte{preamble},
			wantCount: 1,
		},
	}

	for _, test := range tests {
		gotCount, err := expectSinglePreamble()(test.input)
		gotError := err != nil

		if gotError != test.wantError {
			t.Errorf("%s: expectSinglePreamble returned error? %v (%s)", test.name, gotError, err)
		}

		if gotCount != test.wantCount {
			t.Errorf("%s: expectSinglePreamble got %d, want %d", test.name, gotCount, test.wantCount)
		}
	}
}

func Test_expectFullPreamble(t *testing.T) {
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
			name:      "not preamble 1",
			input:     []byte{0x01, preamble},
			wantError: true,
		},
		{
			name:      "not preamble 2",
			input:     []byte{preamble, 0x01},
			wantCount: 1,
			wantError: true,
		},
		{
			name:      "incomplete preamble",
			input:     []byte{preamble},
			wantCount: 1,
			wantError: true,
		},
		{
			name:      "full preamble",
			input:     []byte{preamble, preamble},
			wantCount: 2,
		},
	}

	for _, test := range tests {
		gotCount, err := expectFullPreamble()(test.input)
		gotError := err != nil

		if gotError != test.wantError {
			t.Errorf("%s: expectFullPreamble returned error? %v (%s)", test.name, gotError, err)
		}

		if gotCount != test.wantCount {
			t.Errorf("%s: expectFullPreamble got %d, want %d", test.name, gotCount, test.wantCount)
		}
	}
}

func Test_expectEom(t *testing.T) {
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
			name:      "not eom",
			input:     []byte{0x01},
			wantError: true,
		},
		{
			name:      "eom",
			input:     []byte{eom},
			wantCount: 1,
		},
	}

	for _, test := range tests {
		gotCount, err := expectEom()(test.input)
		gotError := err != nil

		if gotError != test.wantError {
			t.Errorf("%s: expectEom returned error? %v (%s)", test.name, gotError, err)
		}

		if gotCount != test.wantCount {
			t.Errorf("%s: expectEom got %d, want %d", test.name, gotCount, test.wantCount)
		}
	}
}
