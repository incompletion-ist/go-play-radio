package logger

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// LevelChecker is the interface for types that implement CheckLevel.
type LevelChecker[LevelT any] interface {
	CheckLevel(LevelT) (bool, error)
}

// CheckLevel returns true if level is satisfied by checker. It returns an error if checker is nil.
func CheckLevel[LevelT any](checker LevelChecker[LevelT], level LevelT) (bool, error) {
	if checker == nil {
		return false, wraperr.Errorf(ErrorCheckLevel, "CheckLevel called with nil checker")
	}

	checkOk, checkLevelErr := checker.CheckLevel(level)
	wrappedCheckLevelErr := wraperr.WrapfIfErr(ErrorCheckLevel, checkLevelErr, "CheckLevel returned error: %s", checkLevelErr)

	return checkOk, wrappedCheckLevelErr
}
