package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var AppConfig Config

func SetConfig() error {
	cfg, _ := os.UserConfigDir()
	viper.SetConfigFile(filepath.Join(cfg, "toney", "config.toml"))

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	AppConfig = DefaultConfig()

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}
