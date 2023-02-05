package driver

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// ErrorCode identifies the type of error that was encountered.
type ErrorCode int

const (
	// ErrorUndefined is an undefined error.
	ErrorUndefined ErrorCode = iota

	// ErrorPlugin indicates an error was encountered while attempting to load a Go plugin.
	ErrorPlugin
)

// Error is the type for errors returned by this package.
type Error = wraperr.Error[ErrorCode]
