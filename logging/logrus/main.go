package main

import (
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

	// Example 6: Logging a Fatal Error
	logrus.Fatal("This is a fatal error")

}
