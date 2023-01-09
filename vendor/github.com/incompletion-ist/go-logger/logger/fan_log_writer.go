package logger

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

type fanLogWriter[MessageT any] struct {
	writers []LogWriter[MessageT]
}

func (fanWriter fanLogWriter[MessageT]) WriteLog(message MessageT) error {
	for _, writer := range fanWriter.writers {
		if err := WriteLog(writer, message); err != nil {
			return wraperr.Wrapf(ErrorWriteLog, err, "FanLogWriter encountered WritLog error: %s", err)
		}
	}

	return nil
}

// NewFanLogWriter returns a LogWriter that writes messages to all
// writers. It returns an error immediately if any of the writers
// return an error.
func NewFanLogWriter[MessageT any](writers ...LogWriter[MessageT]) LogWriter[MessageT] {
	return fanLogWriter[MessageT]{
		writers: writers,
	}
}
