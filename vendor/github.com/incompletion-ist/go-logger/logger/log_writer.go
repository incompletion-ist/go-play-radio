package logger

import "github.com/incompletion-ist/go-wrapped-error/wraperr"

// LogWriter is the interface for types that write messages of type MessageT.
type LogWriter[MessageT any] interface {
	WriteLog(MessageT) error
}

// WriteLog writes message to writer. It returns an error if writer is nil.
func WriteLog[MessageT any](writer LogWriter[MessageT], message MessageT) error {
	if writer == nil {
		return wraperr.Errorf(ErrorWriteLog, "attempted MustWriteLog on nil LogWriter")
	}

	writeLogErr := writer.WriteLog(message)

	return wraperr.WrapfIfErr(ErrorWriteLog, writeLogErr, "WriteLog returned error: %s", writeLogErr)
}

// MustWriteLog writes message to writer. It panics if an error is encountered.
func MustWriteLog[MessageT any](writer LogWriter[MessageT], message MessageT) {
	if err := WriteLog(writer, message); err != nil {
		panic(err)
	}
}
