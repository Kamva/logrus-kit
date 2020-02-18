package logruskit

type Config interface {
	Get(key string) interface{}
	GetString(key string) string
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetBool(key string) bool
	GetList(key string) []string
}

// Base config keys.
const (
	levelKey     = "log.level"
	outputKey    = "log.output"
	formatterKey = "log.formatter"
	hooksKey     = "log.hooks"
)
