package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Config is the configuration for the backend.
type Config struct {
	GinMode       string `mapstructure:"GIN_MODE"`
	Addr          string `mapstructure:"ADDRESS"`
	Neo4jURL      string `mapstructure:"NEO4J_URL"`
	Neo4jUser     string `mapstructure:"NEO4J_USER"`
	Neo4jPassword string `mapstructure:"NEO4J_PASSWORD"`
}

type Neo4j struct {
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

func (c *Config) SetGinMode() {
	gin.SetMode(c.GinMode)
}
