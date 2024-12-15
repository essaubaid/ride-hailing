package logging

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

// GetLogger initializes and returns the singleton logger instance.
func GetLogger() *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()

		// Set the output to stdout
		logger.Out = os.Stdout

		// Set the log format to JSON
		logger.SetFormatter(&logrus.JSONFormatter{})

		// Set the log level (default to Info)
		logLevel := os.Getenv("LOG_LEVEL")
		level, err := logrus.ParseLevel(logLevel)
		if err != nil {
			level = logrus.InfoLevel // Fallback to Info level
		}
		logger.SetLevel(level)
	})
	return logger
}
