package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type ExtraConfig struct {
	Extra interface{} `mapstructure:"extra"`
}


type Config struct {
	Server            ServerConfig            `mapstructure:"server"`
	Database          DatabaseConfig          `mapstructure:"database"`
	Logging           LoggingConfig           `mapstructure:"logging"`
	Cache             CacheConfig             `mapstructure:"cache"`
	ExternalServices  ExternalServicesConfig  `mapstructure:"external_services"`
	Features          FeaturesConfig          `mapstructure:"features"`
	Security          SecurityConfig          `mapstructure:"security"`
	App               AppConfig               `mapstructure:"app"`
}

type ServerConfig struct {
	Host           string        `mapstructure:"host"`
	Port           int           `mapstructure:"port"`
	Environment    string        `mapstructure:"environment"`
	Debug          bool          `mapstructure:"debug"`
	ReadTimeout    time.Duration `mapstructure:"read_timeout"`
	WriteTimeout   time.Duration `mapstructure:"write_timeout"`
	IdleTimeout    time.Duration `mapstructure:"idle_timeout"`
	ExtraConfig
}

type Database struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	ExtraConfig
}

type LoggingConfig struct {
	Level string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	OutputPath string `mapstructure:"output_path"`
	ExtraConfig
}

type CacheConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	MinIdleConnections int `mapstructure:"min_idle_connections"`
	PoolSize int    `mapstructure:"pool_size"`
	ExtraConfig

}

type ExternalService struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Timeout  int    `mapstructure:"timeout"`
	APIKey   string `mapstructure:"api_key"`
	ExtraConfig
}

type ExternalServicesConfig struct {
	Services map[string]ExternalService `mapstructure:"services"`
}

type DatabaseConfig struct {
	Primary Database `mapstructure:"primary"`
	Secondary Database `mapstructure:"replica"`
}

type FeaturesConfig struct {
	Features map[string]bool `mapstructure:",remain"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	TokenDuration time.Duration `mapstructure:"token_duration"`
	Issuer string `mapstructure:"issuer"`
	ExtraConfig
}

type CorsConfig struct {
	AllowedOrigins   []string `mapstructure:"allowed_origins"`
	AllowedMethods   []string `mapstructure:"allowed_methods"`
	AllowedHeaders   []string `mapstructure:"allowed_headers"`
	ExposedHeaders   []string `mapstructure:"exposed_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int      `mapstructure:"max_age"`
	ExtraConfig
}


type SecurityConfig struct {
	JWT  JWTConfig
	Cors CorsConfig
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	BaseURL string `mapstructure:"base_url"`
	LogLevel string `mapstructure:"log_level"`
	Description string `mapstructure:"description"`
	ExtraConfig
}


func  LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")

	viper.SetEnvPrefix(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

