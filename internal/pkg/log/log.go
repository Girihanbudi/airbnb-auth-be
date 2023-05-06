package log

import (
	"airbnb-auth-be/internal/pkg/log/prefixed"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = defaultFormatter()
	log.Level = logrus.DebugLevel
}

func defaultFormatter() *prefixed.TextFormatter {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true

	// Set specific colors for prefix and timestamp
	formatter.SetColorScheme(&prefixed.ColorScheme{
		PrefixStyle:    "blue+b",
		TimestampStyle: "white+h",
	})

	return formatter
}

func Event(instance, msg string) {
	if instance != "" {
		log.WithFields(logrus.Fields{
			"prefix": instance,
		}).Info(msg)
	} else {
		log.Info(msg)
	}
}

func Fatal(instance, msg string, err error) {
	log.WithFields(logrus.Fields{
		"prefix": instance,
	}).Fatal(msg, ": ", err.Error())
}

func Error(instance, msg string, err error) {
	log.WithFields(logrus.Fields{
		"prefix": instance,
	}).Errorln(msg, ": ", err.Error())
}
