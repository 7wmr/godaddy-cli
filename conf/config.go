package conf

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	Key    string
	Secret string
}

// Load the config.
func (c *Config) Load() error {
	viper.AutomaticEnv()
	c.Key = viper.GetString("GODADDY_KEY")
	c.Secret = viper.GetString("GODADDY_SECRET")

	if c.Key == "" {
		return errors.New("Variable value not set GODADDY_KEY")
	}

	if c.Secret == "" {
		return errors.New("Variable value not set GODADDY_SECRET")
	}

	return nil
}
