package logruskit

import "github.com/spf13/viper"

type ViperAdapter struct {
	viper *viper.Viper
}

func (v ViperAdapter) Get(key string) interface{} {
	return v.viper.Get(key)
}

func (v ViperAdapter) GetString(key string) string {
	return v.viper.GetString(key)
}

func (v ViperAdapter) GetInt64(key string) int64 {
	return v.viper.GetInt64(key)
}

func (v ViperAdapter) GetFloat64(key string) float64 {
	return v.viper.GetFloat64(key)
}

func (v ViperAdapter) GetBool(key string) bool {
	return v.viper.GetBool(key)
}

func (v ViperAdapter) GetList(key string) []string {
	return v.viper.GetStringSlice(key)
}

// NewViperAdapter return new instance of viper adapter to use as logruskit Config.
func NewViperAdapter(viper *viper.Viper) Config {
	return &ViperAdapter{viper: viper}
}

// Assert that viper adapter implements Config.
var _ Config = &ViperAdapter{}
