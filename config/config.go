package config

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/romaxa83/hra/pkg/logger"
	"github.com/spf13/viper"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "HRA service config path")
}

type Config struct {
	ServiceName string         `mapstructure:"serviceName"`
	Logger      *logger.Config `mapstructure:"logger"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv("CONFIG_PATH")
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, nil
}
