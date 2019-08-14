package conf

import (
	"github.com/spf13/viper"
)

type Config struct {
	Key    string
	Secret string
}

// Load the config.
func (c *Config) Load() {
	viper.AutomaticEnv()
	c.Key = viper.GetString("GODADDY_KEY")
	c.Secret = viper.GetString("GODADDY_SECRET")
}
