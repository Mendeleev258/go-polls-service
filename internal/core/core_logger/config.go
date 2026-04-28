package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel  string `envconfig:"LOGGER_LEVEL" required:"true"`
	LogFolder string `envconfig:"LOGGER_FOLDER" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	conf, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get core_logger config: %w", err)
		panic(err)
	}

	return conf
}
