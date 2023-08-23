package logs

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	logger zerolog.Logger
	once   sync.Once
)

func InitializeLogger() {
	once.Do(func() {
		logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	})
}

func GetLoggerInstance() zerolog.Logger {
	InitializeLogger()
	return logger
}
