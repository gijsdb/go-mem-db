package logging

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	// "gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger     zerolog.Logger
	loggerOnce sync.Once
)

func CreateOrGetMultiOutputLogger() zerolog.Logger {
	loggerOnce.Do(func() {
		// fileLogger := lumberjack.Logger{
		// 	Filename:   "/var/log/camera-logger.json",
		// 	MaxSize:    100,
		// 	MaxAge:     28,
		// 	MaxBackups: 3,
		// 	Compress:   false,
		// }

		multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stderr})
		logger = zerolog.New(multi).With().Timestamp().Logger()
	})

	return logger
}
