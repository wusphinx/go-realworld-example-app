package otel

import "github.com/sirupsen/logrus"

type formatter struct {
	Formatter logrus.Formatter
}

func (l formatter) Format(entry *logrus.Entry) ([]byte, error) {
	if entry.Context != nil {
		traceID, _ := entry.Context.Value(defaultTraceIDHeader).(string)
		entry.Data[defaultTraceIDHeader] = traceID
	}

	return l.Formatter.Format(entry)
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&formatter{Formatter: &logrus.JSONFormatter{}})
	return logger
}
