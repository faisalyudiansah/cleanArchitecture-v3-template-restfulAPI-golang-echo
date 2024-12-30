package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// NewViper returns viper instance
func NewViper(filename, filetype string, path ...string) (*viper.Viper, error) {
	config := viper.New()

	config.SetConfigName(filename)
	config.SetConfigType(filetype)

	for _, p := range path {
		config.AddConfigPath(p)
	}

	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		log.Warnf("failed to load %v file", filename)
		// return nil, errors.Wrap(err, "viper failed to read in config")
	}

	return config, nil
}
