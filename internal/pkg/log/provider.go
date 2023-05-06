package log

import "github.com/sirupsen/logrus"

func NewLogger(instance string, trimNewLine bool) *logrus.Entry {
	newLogger := logrus.New()
	formatter := defaultFormatter()
	formatter.TrimNewLine = trimNewLine
	newLogger.Level = logrus.DebugLevel

	entry := newLogger.WithFields(logrus.Fields{
		"prefix": instance,
	})

	entry.Logger.Formatter = formatter

	return entry
}
