package main

import (
	"fmt"
	"github.com/midil-labs/core/pkg/config"
)


func main() {
	cfg, err := config.LoadConfig("./pkg/config")
	if err != nil {
		fmt.Printf("error loading config: %v\n", err)
	}
	fmt.Printf("config: % #v\n", cfg)
}



package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// Environment types
const (
	Development = "development"
	Staging     = "staging"
	Production  = "production"
)

// Configuration interface allows for custom configuration extensions
type Configuration interface {
	Validate() error
	SetDefaults()
}

// GlobalConfig is a thread-safe configuration manager
type GlobalConfig struct {
	mu       sync.RWMutex
	config   *viper.Viper
	env      string
	services map[string]Configuration
}

// NewConfigManager creates a new configuration manager
func NewConfigManager(environment string) *GlobalConfig {
	return &GlobalConfig{
		config:   viper.New(),
		env:      environment,
		services: make(map[string]Configuration),
	}
}

// LoadConfig loads configuration from multiple sources
func (g *GlobalConfig) LoadConfig(configPaths []string, configName string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	// Set configuration name and type
	g.config.SetConfigName(configName)
	g.config.SetConfigType("yaml")

	// Add config paths
	for _, path := range configPaths {
		g.config.AddConfigPath(path)
	}

	// Enable environment variable overrides
	g.config.SetEnvPrefix("APP")
	g.config.AutomaticEnv()
	g.config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read configuration
	if err := g.config.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config: %w", err)
	}

	return nil
}

// RegisterService allows registering service-specific configurations
func (g *GlobalConfig) RegisterService(name string, cfg Configuration) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	// Set defaults for the service configuration
	cfg.SetDefaults()

	// Unmarshal configuration for the specific service
	serviceKey := strings.ToLower(name)
	if err := g.config.UnmarshalKey(serviceKey, cfg); err != nil {
		return fmt.Errorf("error unmarshaling %s config: %w", name, err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("invalid %s configuration: %w", name, err)
	}

	g.services[serviceKey] = cfg
	return nil
}

// GetService retrieves a registered service configuration
func (g *GlobalConfig) GetService(name string) (Configuration, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	service, exists := g.services[strings.ToLower(name)]
	return service, exists
}

// Environment returns the current environment
func (g *GlobalConfig) Environment() string {
	return g.env
}

// Helper function to find config files
func FindConfigFiles(searchPaths []string, configName string) (string, error) {
	for _, path := range searchPaths {
		configPath := filepath.Join(path, configName+".yaml")
		if _, err := os.Stat(configPath); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("config file not found")
}