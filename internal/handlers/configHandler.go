package handlers

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
    Server ServerConfig
    Database DatabaseConfig
}

type ServerConfig struct {
    URI                 string
    Port                int
    WordRotateInterval  int
    readTimeout         int
    writeTimeout        int
}

type DatabaseConfig struct {
    Host            string
    Port            int
    User            string
    Password        string
    DBName          string
    DBTimeout       int
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
        viper.AddConfigPath("./configs/")
        err := viper.ReadInConfig()
        if err != nil {
            cfgErr = fmt.Errorf("Error in reading the config file: %v", err)
        }
        err = viper.Unmarshal(cfg)
        if err != nil {
            cfgErr = fmt.Errorf("Failed to unmarshal config: %v", err)
            return
        }
    })
    return cfg, cfgErr
}
