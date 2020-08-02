package logrusformatter

import (
	"github.com/kamva/logrus-kit"
	"github.com/sirupsen/logrus"
)

func init() {
	logruskit.RegisterFormatter("json", logrusJsonFormatter)
}

// logrusJsonFormatter define return new json formatter.
func logrusJsonFormatter(logruskit.Config) (logrus.Formatter, error) {
	return &logrus.JSONFormatter{}, nil
}
