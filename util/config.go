package util

import (
    "github.com/spf13/viper"
)

type Config struct {
    DBDriver string `mapstructure:"DB_DRIVER"`
    DBSource string `mapstructure:"DB_SOURCE"`
    ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
    // Load configuration from file
    viper.AddConfigPath(path)
    viper.SetConfigName("app")
    viper.SetConfigType("env")

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    // Unmarshal the configuration into the Config struct
    err = viper.Unmarshal(&config)
    return
}