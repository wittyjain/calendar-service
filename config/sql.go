package config

import (
	"github.com/spf13/viper"
)

// SQLConfig holds the SQL configuration parameters
type SQLConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	URL      string `mapstructure:"url"`
	Database string `mapstructure:"database"`
}

// LoadSQLConfig loads SQL configuration from a YAML file
func LoadSQLConfig() (*SQLConfig, error) {
	var config SQLConfig

	viper.SetConfigName("config") 
	viper.AddConfigPath(".")      // Look for the config file in the current directory
	viper.SetConfigType("yaml")   // Config file is of YAML type

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
