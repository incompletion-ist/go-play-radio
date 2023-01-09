package logging

import "github.com/incompletion-ist/go-logger/logger"

// LevelWriter is the interface for writing leveled log messages.
type LevelWriter = logger.LevelLogWriter[Level, Message]
