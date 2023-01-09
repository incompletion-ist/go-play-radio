package logger

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// LevelLogger is a LevelLogWriter composed of a LevelChecker and LogWriter.
type LevelLogger[LevelT any, MessageT any] struct {
	LevelChecker LevelChecker[LevelT]
	LogWriter    LogWriter[MessageT]
}

// WriteLevelLog writes message to the logger's LogWriter if level is satisfied by its LevelChecker.
func (logger LevelLogger[LevelT, MessageT]) WriteLevelLog(level LevelT, message MessageT) error {
	checkLevelOk, checkLevelErr := CheckLevel(logger.LevelChecker, level)
	if checkLevelErr != nil {
		return wraperr.Wrapf(ErrorCheckLevel, checkLevelErr, "LevelLogger encountered CheckLevel error: %s", checkLevelErr)
	}

	if !checkLevelOk {
		return nil
	}

	writeLogErr := WriteLog(logger.LogWriter, message)

	return wraperr.WrapfIfErr(ErrorWriteLog, writeLogErr, "LevelLogger encountered WriteLog error: %s", writeLogErr)
}
