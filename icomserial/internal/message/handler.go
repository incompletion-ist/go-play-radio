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

import "go.incompletion.ist/play-radio/transceiver"

// bytesHandler is a function that performs an action on a slice of bytes,
// returning the number of bytes that were handled.
type bytesHandler func([]byte) (int, error)

// handleBytes runs a series of byteHandlers against data, returning the
// total number of bytes handled.
func handleBytes(data []byte, handlers ...bytesHandler) (int, error) {
	totalCount := 0

	for _, handler := range handlers {
		handlerCount, err := handler(data[totalCount:])
		totalCount += handlerCount
		if err != nil {
			return totalCount, err
		}
	}

	return totalCount, nil
}

func handleNextMessage(conf *transceiver.Configuration) bytesHandler {
	return func(data []byte) (int, error) {
		// checking source/target addresses not implemented
		var throwawayAddress byte

		return handleBytes(data,
			expectFullPreamble(),
			handleAddress(&throwawayAddress),
			handleAddress(&throwawayAddress),
			handleCommand(conf),
			expectEom(),
		)
	}
}

func HandleData(data []byte, conf *transceiver.Configuration) error {
	totalHandledCount := 0

	for len(data) > totalHandledCount {
		singleHandledCount, err := handleNextMessage(conf)(data[totalHandledCount:])
		if err != nil {
			return err
		}

		totalHandledCount += singleHandledCount
	}

	return nil
}
