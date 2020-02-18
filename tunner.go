package logruskit

import (
	"fmt"
	"github.com/Kamva/logrus-kit/internal/util"
	"github.com/sirupsen/logrus"
)

// tuneBaseConfig tune base config on the logger.
func tuneBaseConfig(logger *logrus.Logger, config Config) error {
	lvl := config.GetString(levelKey)

	// Set level
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return err
	}

	logger.SetLevel(level)

	return nil
}

func tuneFormatter(logger *logrus.Logger, config Config) error {
	name := util.FallbackString(config.GetString(formatterKey), "text")
	formatterFn := getFormatter(name)

	if formatterFn == nil {
		return fmt.Errorf("formatter %s not found", name)
	}

	formatter, err := formatterFn(config)

	if err != nil {
		return err
	}

	logger.SetFormatter(formatter)
	return nil
}

// tuneHooks tune defined hooks for logger.
func tuneHooks(logger *logrus.Logger, config Config) error {
	for _, name := range config.GetList(hooksKey) {
		hookFn := getHook(name)

		if hookFn == nil {
			return fmt.Errorf("hook %s not found", name)
		}

		hook, err := hookFn(config)

		if err != nil {
			return err
		}

		logger.AddHook(hook)
	}

	return nil
}

func tuneOutput(logger *logrus.Logger, config Config) error {
	outputFn := getOutput(util.FallbackString(config.GetString(outputKey), "stdout"))
	output, err := outputFn(config)

	if err != nil {
		return err
	}

	logger.SetOutput(output)
	return nil
}

// Tune tune your logger with settings.
func Tune(logger *logrus.Logger, config Config) error {
	// Base tuning
	if err := tuneBaseConfig(logger, config); err != nil {
		return err
	}

	// Tune formatter
	if err := tuneFormatter(logger, config); err != nil {
		return err
	}

	// Tune hooks
	if err := tuneHooks(logger, config); err != nil {
		return err
	}

	return tuneOutput(logger, config)
}
