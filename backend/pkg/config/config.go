package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var ConfigYaml []byte

// Config :
type Config struct {
	Prd  AppConfig `mapstructure:"prd"`
	Stg  AppConfig `mapstructure:"stg"`
	Dev  AppConfig `mapstructure:"dev"`
	Test AppConfig `mapstructure:"test"`
}

// HTTP :
type HTTP struct {
	Cors []string `mapstructure:"cors" validate:"required"`
	Port int      `mapstructure:"port" validate:"required"`
}

// Logger :
type Logger struct {
	Debug bool `mapstructure:"debug"`
}

// Postgres :
type Postgres struct {
	DBName  string `mapstructure:"dbname" validate:"required"`
	Host    string `mapstructure:"host" validate:"required"`
	Pass    string `mapstructure:"pass" validate:"required"`
	Port    string `mapstructure:"port" validate:"required"`
	Sslmode string `mapstructure:"sslmode" validate:"required"`
	User    string `mapstructure:"user" validate:"required"`
	Pseudo  bool
}

// AppConfig :
type AppConfig struct {
	HTTP     HTTP     `mapstructure:"http"`
	Logger   Logger   `mapstructure:"logger"`
	Postgres Postgres `mapstructure:"postgres"`
}

// Prepare :
func Prepare() AppConfig {
	viper.SetConfigName("config")
	viper.SetEnvPrefix("yondeco")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(bytes.NewReader(ConfigYaml)); err != nil {
		slog.Error("Failed to read config", slog.Any("err", err))
		os.Exit(1)
	}
	viper.AutomaticEnv()

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		slog.Error("Failed to unmarshal config", slog.Any("err", err))
		os.Exit(1)
	}

	var appConfig AppConfig
	env := viper.GetString("env.name")
	switch env {
	case "prd":
		appConfig = c.Prd
	case "stg":
		appConfig = c.Stg
	case "dev":
		appConfig = c.Dev
	case "test":
		// no validate when test
		appConfig = c.Test
	default:
		panic(fmt.Sprintf("Unknown env: %s", env))
	}
	return appConfig
}
