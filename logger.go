package logging

import (
	"log"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
)

var (
	standardFields logrus.Fields
	logger         *logrus.Logger

	// Info logs a info level message
	Info = logger.Info

	// Warn logs a warn level message
	Warn = logger.Warn

	// Error logs a error level message
	Error = logger.Error

	// Fatal logs a fatal level message
	Fatal = logger.Fatal
)

func init() {
	logger = logrus.New()
	logger.Formatter = customFormatter{&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	}}
}

// SetStandardFields sets up the service name, version, hostname and pid fields
func SetStandardFields(service, version string) {
	hostname, _ := os.Hostname()
	standardFields = logrus.Fields{
		"service":  service,
		"version":  version,
		"hostname": hostname,
		"pid":      os.Getpid(),
	}
}

// UsePrettyPrint tells the logger to print in human readable format
func UsePrettyPrint() {
	logger.Formatter = customFormatter{&logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  time.RFC3339Nano,
		QuoteEmptyFields: true,
	}}
}

// ErrorLogger creates a logger than can plug in to an HTTP server
func ErrorLogger() (basicLogger *log.Logger, dispose func()) {
	w := logger.WriterLevel(logrus.ErrorLevel)
	basicLogger = log.New(w, "", 0)
	dispose = func() {
		w.Close()
	}

	return
}
