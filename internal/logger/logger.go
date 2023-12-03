package logger

import (
	"context"
	"io"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func init() {
	// unix timestamp is faster to parse
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

// New returns a new Logger instance
// w is the writer to write the log to, usually set as StdOut or file stream
func New(w io.Writer) *Logger {
	syncer := zerolog.SyncWriter(w)
	l := zerolog.New(syncer).With().Timestamp().Logger()

	return &Logger{
		Logger: &l,
	}
}

// WithContext attaches a context to the logger
func (l *Logger) WithContext(ctx context.Context) context.Context {
	return l.Logger.WithContext(ctx)
}

// FromContext returns a logger instance from the context
func FromContext(ctx context.Context) *Logger {
	return &Logger{
		Logger: zerolog.Ctx(ctx),
	}
}
