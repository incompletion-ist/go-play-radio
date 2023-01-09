package logging

import "github.com/incompletion-ist/go-logger/logger"

// Writer is the interface for writing log messages.
type Writer = logger.LogWriter[Message]
