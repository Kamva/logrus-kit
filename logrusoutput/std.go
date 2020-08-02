package logrusoutput

import (
	"github.com/kamva/logrus-kit"
	"io"
	"os"
)

func init() {
	logruskit.RegisterOutput("stdout", NewStdoutOutput)
	logruskit.RegisterOutput("stderr", NewStderrWriter)
}

func NewStdoutOutput(logruskit.Config) (writer io.Writer, err error) {
	writer = os.Stdout
	return
}

func NewStderrWriter(logruskit.Config) (writer io.Writer, err error) {
	writer = os.Stderr
	return
}
