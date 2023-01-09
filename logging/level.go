package logging

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// Level is an ordered logging level.
type Level int

const (
	// Fatal indicates a fatal error was encountered. It is the lowest level, and
	// so is always satisfied by any configured level.
	Fatal Level = iota

	// Warn is used for a potentially serious or inefficient situation or configuration.
	Warn

	// Info is used to convey important, but not verbose, messages about the process.
	Info

	// Debug is used for verbose debugging information.
	Debug

	// levelTooHigh is the lowest unknown level. Additional levels must be added above
	// levelTooHigh. As it is unexported, it is subject to change.
	levelTooHigh
)

// Name returns the name of the Level. It returns an error if the Level is not one of the
// known constants.
func (level Level) Name() (string, error) {
	if level >= levelTooHigh {
		err := wraperr.Errorf(ErrorUnknownLogLevel, "unknown LogLevel: %d", level)

		return "", err
	}

	levelNames := map[Level]string{
		Fatal: "FATAL",
		Warn:  "WARN",
		Info:  "INFO",
		Debug: "DEBUG",
	}

	levelName, ok := levelNames[level]

	if !ok {
		err := wraperr.Errorf(ErrorUnknownLogLevel, "unexpectedly unknown LogLevel name: %d", level)

		return "", err
	}

	return levelName, nil
}

// CheckLevel returns true if checkLevel is satisfied by the current Level. It returns an
// error if the configured or check Level is not one of the known constants.
func (level Level) CheckLevel(checkLevel Level) (bool, error) {
	if _, err := level.Name(); err != nil {
		return false, err
	}

	if _, err := checkLevel.Name(); err != nil {
		return false, err
	}

	return level >= checkLevel, nil
}
