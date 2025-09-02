package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var AppConfig Config

func SetConfig() error {
	cfg := getCfgDir()

	configDir := filepath.Join(cfg, "toney")
	configFile := filepath.Join(configDir, "config.toml")

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Check if config file exists, if not create it with defaults
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		AppConfig = DefaultConfig()
		viper.SetConfigFile(configFile)
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to write default config: %w", err)
		}
		return nil
	}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	AppConfig = DefaultConfig()
	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	return nil
}

func getCfgDir() string {
	switch runtime.GOOS {
	case "linux", "darwin":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".config", "toney")
	case "windows":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".toney")
	default:
		cfg, _ := os.UserConfigDir()
		return cfg
	}
}
