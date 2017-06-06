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

// Info logs an info level message with standard fields
func Info(msg string) {
	logger.Info(msg)
}

// Warn logs a warn level message with standard fields
func Warn(msg string) {
	logger.Warn(msg)
}

// Error logs an error level message with standard fields
func Error(msg string) {
	logger.Error(msg)
}

// Fatal logs a fatal level message with standard fields
func Fatal(msg string) {
	logger.Fatal(msg)
}
