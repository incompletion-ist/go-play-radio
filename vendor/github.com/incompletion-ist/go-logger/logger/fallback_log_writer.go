package logger

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

type fallbackLogWriter[MessageT any] struct {
	writers []LogWriter[MessageT]
}

func (fallbackWriter fallbackLogWriter[MessageT]) WriteLog(message MessageT) error {
	var lastError error
	for _, writer := range fallbackWriter.writers {
		lastError = WriteLog(writer, message)

		if lastError == nil {
			return nil
		}

		if withErrorer, ok := any(message).(WithErrorer[MessageT]); ok {
			message = withErrorer.WithError(lastError)
		}
	}

	return wraperr.Wrapf(ErrorWriteLog, lastError, "FallbackLogWriter last LogWriter returned error: %s", lastError)
}

// NewFallbackLogWriter returns a LogWriter that writes a message to writers until
// no error is returned. If returns an error if the final writer returns an error.
func NewFallbackLogWriter[MessageT any](writers ...LogWriter[MessageT]) LogWriter[MessageT] {
	return fallbackLogWriter[MessageT]{
		writers: writers,
	}
}
