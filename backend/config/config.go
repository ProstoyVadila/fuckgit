package config

import "github.com/spf13/viper"

// Config is the configuration for the backend.
type Config struct {
	GinMode string `mapstructure:"GIN_MODE"`
	Addr    string `mapstructure:"ADDRESS"`
}

func New(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
