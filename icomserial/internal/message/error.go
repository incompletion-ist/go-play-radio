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

// ErrorCode describes the type of error encountered.
type ErrorCode int

const (
	// ErrorUndefined indicates ErrorCode is unset.
	ErrorUndefined ErrorCode = iota

	// ErrorParsing indicates an error parsing data.
	ErrorParsing

	// ErrorCommandUnknown indicates a command is unknown. This may simply
	// mean the command is not yet implemented.
	ErrorCommandUnknown
)
