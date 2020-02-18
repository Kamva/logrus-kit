// Inspired by logrus_mate.
package logrushook

import (
	"github.com/Kamva/logrus-kit"
	"github.com/johntdyer/slackrus"
	"github.com/sirupsen/logrus"
)

const (
	slackUrlKey      = "log.slack.url"
	slackLevelsKey   = "log.slack.levels"
	slackChannelKey  = "log.slack.channel"
	slackEmojiKey    = "log.slack.emoji"
	slackUsernameKey = "log.slack.username"
)

type SlackHookConfig struct {
	URL      string
	Levels   []string
	Channel  string
	Emoji    string
	Username string
}

func init() {
	logruskit.RegisterHook("slack", NewSlackHook)
}

func NewSlackHook(config logruskit.Config) (hook logrus.Hook, err error) {
	conf := SlackHookConfig{}

	if config != nil {
		conf.URL = config.GetString(slackUrlKey)
		conf.Levels = config.GetList(slackLevelsKey)
		conf.Channel = config.GetString(slackChannelKey)
		conf.Emoji = config.GetString(slackEmojiKey)
		conf.Username = config.GetString(slackUsernameKey)
	}

	levels := []logrus.Level{}

	if conf.Levels != nil {
		for _, level := range conf.Levels {
			if lv, e := logrus.ParseLevel(level); e != nil {
				err = e
				return
			} else {
				levels = append(levels, lv)
			}
		}
	}

	if len(levels) == 0 && conf.Levels != nil {
		levels = append(levels, logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel)
	}

	hook = &slackrus.SlackrusHook{
		HookURL:        conf.URL,
		AcceptedLevels: levels,
		Channel:        conf.Channel,
		IconEmoji:      conf.Emoji,
		Username:       conf.Username,
	}

	return
}
