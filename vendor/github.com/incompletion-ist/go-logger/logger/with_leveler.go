package logger

// WithLeveler is the interface for types that implement WithLevel.
type WithLeveler[LevelT any, MessageT any] interface {
	WithLevel(LevelT) (MessageT, error)
}
