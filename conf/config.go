package conf

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Key    string
	Secret string
	API    string
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

// GetAuth return the header string for API call.
func (c *Config) GetAuth() (string, string) {
	return "Authorization", fmt.Sprintf("sso-key %s:%s", c.Key, c.Secret)
}

// GetAPI return the API url string.
func (c *Config) GetAPI() string {
	return "https://api.godaddy.com"
}
