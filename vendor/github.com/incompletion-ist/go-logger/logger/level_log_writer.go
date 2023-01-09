package logger

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// LevelLogWriter is the interface for types that implement WriteLevelLog.
type LevelLogWriter[LevelT any, MessageT any] interface {
	// WriteLevelLog writes the message if level is satisfied.
	WriteLevelLog(LevelT, MessageT) error
}

// WriteLevelLog writes message to writer if level is satisfied.
//
// If message is a WithLeveler, it will have level added to it via WithLevel prior to being written.
//
// Returns an error if writer is nil.
func WriteLevelLog[LevelT any, MessageT any](writer LevelLogWriter[LevelT, MessageT], level LevelT, message MessageT) error {
	if writer == nil {
		return wraperr.Errorf(ErrorWriteLevelLog, "attempted WriteLevelLog on nil LevelLogWriter")
	}

	var messageInterface any = message
	if formatter, formatterOk := messageInterface.(WithLeveler[LevelT, MessageT]); formatterOk {
		var formatterErr error
		message, formatterErr = formatter.WithLevel(level)
		if formatterErr != nil {
			return wraperr.Wrapf(ErrorWithLevel, formatterErr, "WithLevel returned error: %s", formatterErr)
		}
	}

	writeLogErr := writer.WriteLevelLog(level, message)

	return wraperr.WrapfIfErr(ErrorWriteLog, writeLogErr, "WriteLog returned error: %s", writeLogErr)
}

// MustWriteLevelLog calls WriteLevelLog and panics if an error is encountered.
func MustWriteLevelLog[LevelT any, MessageT any](writer LevelLogWriter[LevelT, MessageT], level LevelT, message MessageT) {
	if err := WriteLevelLog(writer, level, message); err != nil {
		panic(err)
	}
}
