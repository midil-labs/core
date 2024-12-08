package config

import (
	"log"
	"strings"
	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"port"`
	Environment  string `mapstructure:"environment"`
	Debug        bool   `mapstructure:"debug"`
	MongoURI     string `mapstructure:"mongo_uri"`
	RedisAddr    string `mapstructure:"redis_addr"`
	RedisDB      int    `mapstructure:"redis_db"`
	RedisPass    string `mapstructure:"redis_pass"`
	TracingURL   string `mapstructure:"tracing_url"`
	PrometheusOn bool   `mapstructure:"prometheus_on"`
}

func Load() (*Config, error) {
	v := viper.New()
	// Set defaults
	v.SetDefault("port", "8080")
	v.SetDefault("environment", "development")
	v.SetDefault("debug", false)
	v.SetDefault("prometheus_on", true)

	// Env variables are primary
	v.SetEnvPrefix("myapp") // e.g. MYAPP_PORT, MYAPP_ENVIRONMENT
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Optionally read from config file if exists (yaml, json, toml)
	v.SetConfigName("config") 
	v.SetConfigType("yaml")
	v.AddConfigPath(".") 
	v.AddConfigPath("/etc/myapp/") 
	_ = v.ReadInConfig() // no error check: we can handle missing file scenario

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	log.Printf("Configuration loaded: %+v", cfg)
	return &cfg, nil
}
