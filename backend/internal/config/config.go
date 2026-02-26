package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Addr string
	}
	DB struct {
		DSN string
	}
	JWT struct {
		Secret       string
		AccessTTL    time.Duration
		RefreshTTL   time.Duration
		Issuer       string
		CookieDomain string
	}
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")

	// defaults
	v.SetDefault("server.addr", ":18601")
	v.SetDefault("jwt.secret", "dev-secret-change-me")
	v.SetDefault("jwt.accessTTL", "30m")
	v.SetDefault("jwt.refreshTTL", "168h")
	v.SetDefault("jwt.issuer", "go-user-center")

	v.SetEnvPrefix("UC")
	v.BindEnv("db.dsn", "DB_DSN")
	v.BindEnv("server.addr", "SERVER_ADDR")
	v.BindEnv("jwt.secret", "JWT_SECRET")
	v.BindEnv("jwt.accessTTL", "JWT_ACCESSTTL")
	v.BindEnv("jwt.refreshTTL", "JWT_REFRESHTTL")
	v.BindEnv("jwt.issuer", "JWT_ISSUER")
	v.AutomaticEnv()

	_ = v.ReadInConfig() // optional

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
