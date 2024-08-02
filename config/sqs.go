package config

import (
	"sync"

	"github.com/spf13/viper"
)

// SQSConfig holds the SQS configuration parameters
type SQSConfig struct {
	Region     string `mapstructure:"region"`
	Endpoint   string `mapstructure:"endpoint"`
	Credentials struct {
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"credentials"`
	Queues map[string]string `mapstructure:"queues"` // Maps queue names to URLs
}

// singleton instance and mutex for SQSConfig
var (
	instance *SQSConfig
	once     sync.Once
)

// LoadSQSConfig loads SQS configuration from a YAML file
func LoadSQSConfig() (*SQSConfig, error) {
	var config SQSConfig
	var err error

	once.Do(func() {
		viper.SetConfigName("config") // Name of the config file (without extension)
		viper.AddConfigPath(".")      // Look for the config file in the current directory
		viper.SetConfigType("yaml")   // Config file is of YAML type

		if err = viper.ReadInConfig(); err != nil {
			return
		}

		if err = viper.Unmarshal(&config); err != nil {
			return
		}

		instance = &config
	})

	return instance, err
}
