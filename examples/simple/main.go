package main

import (
	"github.com/kamva/logrus-kit"
	_ "github.com/kamva/logrus-kit/logrusbase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.SetDefault("log.level", "info")
	v.SetDefault("log.formatter", "text")
	v.SetDefault("log.output", "stderr")
	v.SetDefault("log.hooks", "")

	logger := logrus.New()

	err := logruskit.Tune(logger, logruskit.NewViperAdapter(v))

	if err != nil {
		panic(err)
	}

	logger.WithFields(logrus.Fields{"test": "is ok"}).Info("salam")
}
