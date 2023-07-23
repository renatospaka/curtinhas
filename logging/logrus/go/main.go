package main

import (
	"log/syslog"
	"os"

	"github.com/sirupsen/logrus"
	slhooks "github.com/sirupsen/logrus/hooks/syslog"
)

func main() {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	// log.Formatter.(*logrus.JSONFormatter).DisableColors = false
	log.Formatter.(*logrus.JSONFormatter).DisableTimestamp = false

	sl, err := slhooks.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")
	if err != nil {
		log.Fatalf("error occurred: %v\n", err)
	}
	log.Hooks.Add(sl)
	log.Out = os.Stdout

	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logrus.Entry)
			log.WithFields(logrus.Fields{
				"omg":         true,
				"err_animal":  entry.Data["animal"],
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
				"number":      100,
			}).Error("The ice breaks!")
			log.Fatal("This is the END") // or use Fatal() to force the process to exit with a nonzero code
		}
	}()

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 0,
	}).Trace("Went to the beach")

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"temperature": -4,
	}).Debug("Temperature changes")

	log.WithFields(logrus.Fields{
		"animal": "orca",
		"size":   9009,
	}).Panic("It's over 9000!")
}
