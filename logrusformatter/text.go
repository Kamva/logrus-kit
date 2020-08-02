package logrusformatter

import (
	"github.com/kamva/logrus-kit"
	"github.com/sirupsen/logrus"
)

func init() {
	logruskit.RegisterFormatter("text", logrusTextFormatter)
}

// logrusTextFormatter return new logrus text formatter.
func logrusTextFormatter(logruskit.Config) (logrus.Formatter, error) {
	return &logrus.TextFormatter{}, nil
}
