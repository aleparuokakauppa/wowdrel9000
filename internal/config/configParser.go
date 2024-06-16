package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
    Server ServerConfig
    Words WordFile
}

type ServerConfig struct {
    Port                int
    WordRotateInterval  int
}

type WordFile struct {
    Path                string
}

var (
    cfg *Config
    cfgErr error
    cfgOnce sync.Once
)

func GetConfig() (*Config, error) {
    cfgOnce.Do(func() {
        viper.SetConfigName("config")
        viper.SetConfigType("yaml")
        viper.AddConfigPath("./cfg/")
        err := viper.ReadInConfig()
        if err != nil {
            cfgErr = fmt.Errorf("Error in reading the config file: %v", err)
        }
        err = viper.Unmarshal(&cfg)
        if err != nil {
            cfgErr = fmt.Errorf("Failed to unmarshal config: %v", err)
            return
        }
    })
    return cfg, cfgErr
}
