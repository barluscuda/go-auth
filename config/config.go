package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerMode string `mapstructure:"SERVER_MODE"`
	ServerPort int    `mapstructure:"SERVER_PORT"`

	DBUrl string `mapstructure:"DB_URL"`

	JWTAccessSecret  string        `mapstructure:"JWT_ACCESS_SECRET"`
	JWTRefreshSecret string        `mapstructure:"JWT_REFRESH_SECRET"`
	JWTAccessExpire  time.Duration `mapstructure:"JWT_ACCESS_EXPIRE"`
	JWTRefreshExpire time.Duration `mapstructure:"JWT_REFRESH_EXPIRE"`
}

func Load() (*Config, error) {
	v := viper.New()

	// read .env
	v.SetConfigFile(".env")
	v.SetConfigType("env")

	// fallback system env
	v.AutomaticEnv()

	// ensure env naming compatibility
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// default values
	v.SetDefault("SERVER_MODE", "dev")
	v.SetDefault("SERVER_PORT", 3200)

	v.SetDefault("JWT_ACCESS_EXPIRE", "24h")
	v.SetDefault("JWT_REFRESH_EXPIRE", "720h") // 30 days

	// load .env if exists
	_ = v.ReadInConfig()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}