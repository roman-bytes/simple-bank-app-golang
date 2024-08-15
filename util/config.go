package util

import (
    "github.com/spf13/viper"
    "time"
)

type Config struct {
    DBDriver string `mapstructure:"DB_DRIVER"`
    DBSource string `mapstructure:"DB_SOURCE"`
    ServerAddress string `mapstructure:"SERVER_ADDRESS"`
    TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
    AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
    RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
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