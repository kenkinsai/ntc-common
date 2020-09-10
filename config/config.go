package config

import (
	"fmt"
	"os"

	"github.com/kenkinsai/ntc-common/adapter"
	Viper "github.com/spf13/viper"
)

// Config struct ..
type Config struct {
	Name    string            `mapstructure:"name"`
	Port    map[string]string `mapstructure:"port"`
	Version string            `mapstructure:"version"`
	Debug   bool              `mapstructure:"debug"`
	MYSQL   adapter.MYSQLs
	API     adapter.APIs
	Other   adapter.Others
}

var config *Config

func init() {
	var folder string

	env := os.Getenv("APPLICATION_ENV")

	switch env {
	case "master", "dev", "uat":
		folder = env
	default:
		folder = "local"
	}

	path := fmt.Sprintf("config/%v", folder)

	//Get base config
	config = new(Config)
	fetchDataToConfig(path, "base", config)

	//Get all sub config
	fetchDataToConfig(path, "other", &(config.Other))
	fetchDataToConfig(path, "api", &(config.API))
	fetchDataToConfig(path, "mysql", &(config.MYSQL))
}

func fetchDataToConfig(configPath, configName string, result interface{}) {
	viper := Viper.New()
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	err := viper.ReadInConfig() // Find and read the config file
	if err == nil {             // Handle errors reading the config file
		err = viper.Unmarshal(result)
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}
	}
}

// GetConfig func
func GetConfig() *Config {
	return config
}
