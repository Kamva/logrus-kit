package logrusoutput

import (
	"github.com/Kamva/logrus-kit"
	"io"
)

func init() {
	logruskit.RegisterOutput("null", NewNullOutput)
}

type NullWriter struct {
}

func (w *NullWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func NewNullOutput(logruskit.Config) (writer io.Writer, err error) {
	writer = new(NullWriter)
	return
}
