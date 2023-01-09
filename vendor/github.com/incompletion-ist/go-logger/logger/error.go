package logger

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// ErrorCode identifies the type of error encountered.
type ErrorCode int

const (
	// ErrorUndefined is an undefined ErrorCode.
	ErrorUnknown ErrorCode = iota

	// ErrorCheckLevel indicates CheckLevel returned an error.
	ErrorCheckLevel

	// ErrorWithLevel indicates WithLevel returned an error.
	ErrorWithLevel

	// ErrorWriteLog indicates WriteLog returned an error.
	ErrorWriteLog

	// ErrorWriteLevelLog indicates WriterLevelLog returned an error.
	ErrorWriteLevelLog
)

// Error is an error returned by the logger package.
type Error = wraperr.Error[ErrorCode]
