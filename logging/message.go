package logging

import "fmt"

// Message represents a log message.
type Message string

// WithLevel returns a new Message with level's name prepended. It
// returns the original Message and an error if level is unknown.
func (message Message) WithLevel(level Level) (Message, error) {
	levelName, err := level.Name()
	if err != nil {
		return message, err
	}

	withLevel := fmt.Sprintf("%s: %s", levelName, message)

	return Message(withLevel), nil
}

// WithError returns a new Message with err appended. It returns the original
// Message if err is nil.
func (message Message) WithError(err error) Message {
	if err == nil {
		return message
	}

	withError := fmt.Sprintf("%s (with error: %s)", message, err)

	return Message(withError)
}
