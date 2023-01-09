package logging

import (
	"errors"
	"testing"

	"github.com/incompletion-ist/go-logger/logger"
)

// TestMessage_WithErroer tests that Message is a WithErroer[Message].
func TestMessage_WithErroer(t *testing.T) {
	var withErrorer logger.WithErrorer[Message]
	var message Message

	withErrorer = message
	_ = withErrorer.WithError(nil)
}

// TestMessage_WithLeveler tests that Message is a WithLeveler[Level, Message].
func TestMessage_WithLeveler(t *testing.T) {
	var withLeveler logger.WithLeveler[Level, Message]
	var message Message

	withLeveler = message
	_, _ = withLeveler.WithLevel(Fatal)
}

func TestMessage_WithLevel(t *testing.T) {
	tests := []struct {
		name         string
		inputMessage Message
		inputLevel   Level
		wantMessage  Message
		wantError    bool
	}{
		{
			name:        "zero value message and level",
			wantMessage: "FATAL: ",
		},
		{
			name:         "non-zero message and level",
			inputMessage: "original message",
			inputLevel:   Warn,
			wantMessage:  "WARN: original message",
		},
		{
			name:         "unknown level",
			inputMessage: "original message",
			inputLevel:   levelTooHigh,
			wantMessage:  "original message",
			wantError:    true,
		},
	}

	for _, test := range tests {
		gotMessage, err := test.inputMessage.WithLevel(test.inputLevel)
		gotError := err != nil

		if gotError != test.wantError {
			t.Errorf("%s: WithLevel returned error? %v (%s)", test.name, gotError, err)
		}

		if gotMessage != test.wantMessage {
			t.Errorf("%s: WithLevel got %s, want %s", test.name, gotMessage, test.wantMessage)
		}
	}
}

func TestMessage_WithError(t *testing.T) {
	tests := []struct {
		name         string
		inputMessage Message
		inputError   error
		wantMessage  Message
	}{
		{
			name: "empty",
		},
		{
			name:         "nil error",
			inputMessage: "original message",
			wantMessage:  "original message",
		},
		{
			name:         "non-nil error",
			inputMessage: "original message",
			inputError:   errors.New("forced error"),
			wantMessage:  "original message (with error: forced error)",
		},
	}

	for _, test := range tests {
		gotMessage := test.inputMessage.WithError(test.inputError)

		if gotMessage != test.wantMessage {
			t.Errorf("%s: WithError got %s, want %s", test.name, gotMessage, test.wantMessage)
		}
	}
}
