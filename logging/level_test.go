package logging

import (
	"testing"

	"github.com/incompletion-ist/go-logger/logger"
)

// TestLevel_LevelChecker tests that Level is a valid LevelChecker[Level]
func TestLevel_LevelChecker(t *testing.T) {
	_, _ = logger.CheckLevel[Level](Level(0), Level(1))
}

func TestLevel_CheckLevel(t *testing.T) {
	tests := []struct {
		name       string
		inputlevel Level
		checkLevel Level
		wantError  bool
		wantOk     bool
	}{
		{
			name:   "empty",
			wantOk: true,
		},
		{
			name:       "equal, non-zero",
			inputlevel: Info,
			checkLevel: Info,
			wantOk:     true,
		},
		{
			name:       "non-equal, satisfied",
			inputlevel: Info,
			checkLevel: Warn,
			wantOk:     true,
		},
		{
			name:       "non-equal, not satisfied",
			inputlevel: Info,
			checkLevel: Debug,
		},
		{
			name:       "unknown configured level",
			inputlevel: 1000,
			wantError:  true,
		},
		{
			name:       "unknown checkLevel",
			checkLevel: levelTooHigh,
			wantError:  true,
		},
	}

	for _, test := range tests {
		gotOk, err := test.inputlevel.CheckLevel(test.checkLevel)
		gotError := err != nil

		if gotError != test.wantError {
			t.Errorf("%s: CheckLevel returned error? %v (%s)", test.name, gotError, err)
		}

		if gotOk != test.wantOk {
			t.Errorf("%s: CheckLevel got %v", test.name, gotOk)
		}
	}
}
