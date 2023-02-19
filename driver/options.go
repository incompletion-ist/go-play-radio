package driver

import "github.com/incompletion-ist/go-play-radio/logging"

// Options is a struct of optional dependencies that may be utilized by an OpenCloser.
type Options struct {
	LogWriter logging.LevelWriter
}

// NewOptions returns a new Options by applying each provided Option.
func NewOptions(optionFuncs ...Option) Options {
	var newOptions Options

	for _, optionFunc := range optionFuncs {
		optionFunc(&newOptions)
	}

	return newOptions
}

// Option is a function that operates on an Options value.
type Option func(*Options)

// WithLogWriter is an Option that sets LogWriter.
func WithLogWriter(logWriter logging.LevelWriter) Option {
	return func(o *Options) {
		o.LogWriter = logWriter
	}
}

// WithOptions is an Option that replaces the entire existing Options with the Options provided.
func WithOptions(options Options) Option {
	return func(o *Options) {
		*o = options
	}
}
