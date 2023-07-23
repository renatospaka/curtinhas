package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type BugTrackerHook struct{}
func (h *BugTrackerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
		logrus.InfoLevel,
	}
}

func (hook *BugTrackerHook) Fire(entry *logrus.Entry) error {
	// Log the entry to your bug tracker here
	return nil
}

type CustomFormatter struct{}
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s: %s\n", entry.Level.String(), entry.Message)), nil
}

func main() {
	// Example 1: Basic Logging with Logrus
	logrus.Error("Something went wrong!")

	// Example 2: Formatting Logs with Logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("This is an informational message.")

	// Example 3: Implementing Hooks with Logrus
	logrus.AddHook(&BugTrackerHook{})
	logrus.Warn("This is a warning message.")

	// Example 4: Setting the Log Level
	logrus.WithFields(logrus.Fields{
		"metric": "credit-card",
		"quantity": 13,
	}).Info("New card added")

	// Example 8: Using The Trace Level
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("This is a trace log")

	// Example 9: Logging to an io.Writer
	logrus.SetOutput(os.Stdout)
	logrus.Info("This log is written to stdout")

	// Example 10: Logging in Different Timezones
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339})
	logrus.Info("Log with RFC3339 timestamp format.")

	// Example 11: Log Entry Order
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg: "@message",
		},
	})
	logrus.Info("Log with consistent field order.")

	// // Example 12: Log Redaction
	// logrus.SetFormatter(&logrus.TextFormatter{DisableKeys: true})
	// logrus.WithField("password", "secret").Info("Sensitive data is redacted.")

	// Example 13: Custom Log Formatter
	logrus.SetFormatter(new(CustomFormatter))
	logrus.Info("Log with a custom formatter.")

	// Example 6: Logging a Fatal Error
	logrus.Fatal("This is a fatal error.")

	// // Example 7: Logging With a Panic
	// logrus.Panic("This is a panic error")
}
