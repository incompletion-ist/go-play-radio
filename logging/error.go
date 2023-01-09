package logging

// ErrorCode indicates which type of error occurred.
type ErrorCode int

const (
	// ErrorUndefined is an undefined error.
	ErrorUndefined ErrorCode = iota

	// ErrorUnknownLogLevel indicates an unknown LogLevel was used.
	ErrorUnknownLogLevel
)
