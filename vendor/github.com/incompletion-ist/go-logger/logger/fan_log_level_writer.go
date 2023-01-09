package logger

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

type fanLevelLogWriter[LevelT any, MessageT any] struct {
	writers []LevelLogWriter[LevelT, MessageT]
}

func (fanWriter fanLevelLogWriter[LevelT, MessageT]) WriteLevelLog(level LevelT, message MessageT) error {
	for _, writer := range fanWriter.writers {
		if writer == nil {
			return wraperr.Errorf(ErrorWriteLevelLog, "FanLevelLogWriter attempted WriteLevelLog on nil LevelLogWriter")
		}

		if err := writer.WriteLevelLog(level, message); err != nil {
			return wraperr.Wrapf(ErrorWriteLevelLog, err, "FanLevelLogWriter encountered WriteLevelLog error: %s", err)
		}
	}

	return nil
}

// NewFanLevelLogWriter returns a LevelLogWriter that writes messages
// to all writers. It returns an error immediately if any of the writers
// return an error.
func NewFanLevelLogWriter[LevelT any, MessageT any](writers ...LevelLogWriter[LevelT, MessageT]) LevelLogWriter[LevelT, MessageT] {
	return fanLevelLogWriter[LevelT, MessageT]{
		writers: writers,
	}
}
