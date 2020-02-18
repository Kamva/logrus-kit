package logruskit

import (
	"github.com/sirupsen/logrus"
	"io"
)

// FormatterFunc should use to define new formatter.
type FormatterFunc func(config Config) (logrus.Formatter, error)

// HookFunc should use to define new hook.
type HookFunc func(config Config) (logrus.Hook, error)

// OutputFunc is function to use for defining new outputs.
type OutputFunc func(config Config) (io.Writer, error)

var (
	formatters = make(map[string]FormatterFunc)
	hooks      = make(map[string]HookFunc)
	outputs    = make(map[string]OutputFunc)
)

// RegisterOutput register writer.
func RegisterOutput(name string, fn OutputFunc) {
	outputs[name] = fn
}

// RegisterFormatter function register new formatter.
func RegisterFormatter(name string, fn FormatterFunc) {
	formatters[name] = fn
}

// RegisterHook function register new hook.
func RegisterHook(name string, fn HookFunc) {
	hooks[name] = fn
}

func getFormatter(name string) FormatterFunc {
	return formatters[name]
}

func getHook(name string) HookFunc {
	return hooks[name]
}

func getOutput(name string) OutputFunc {
	return outputs[name]
}
