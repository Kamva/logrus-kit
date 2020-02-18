### Logrus-kit is a kit for logrus.
**Logrus-kit inspired by [Logrus mate](https://github.com/gogap/logrus_mate)** 

#### Install
```
go get github.com/Kamva/logrus-kit
```

#### Formatter list:
- json
- text

#### Hook list:
- slack

#### Output list
- stdout
- stderr
- null

#### Example
**Note:** To register default logrus-kit formatters,hooks and 
outputs , import logrusbase : e.g   
`import _ "github.com/Kamva/logrus-kit/logrusbase"`

Example: 
```go
package mypackage

import (
	"github.com/Kamva/logrus-kit"
	_ "github.com/Kamva/logrus-kit/logrusbase"
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
```

#### Todo: 
- [ ] Add more formatter and logger.
- [ ] Write tests.
- [ ] CI.
- [ ] Write example.
- [ ] PR to logrus to add this to other logrus tools.