package config

import (
	"context"
	"log"
	"strings"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

// Config represents the configuration for the application.
var Config Cfg

// Cfg represents the struct for the configuration.
type Cfg struct {
	Domain      string `mapstructure:"domain"       validate:"required"`
	KeyPath     string `mapstructure:"key_path"     validate:"required"`
	ServerPort  string `mapstructure:"server_port"  validate:"required"`
	FrontendURL string `mapstructure:"frontend_url" validate:"required"`
}

func init() {
	viper.AllowEmptyEnv(false)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	_ = viper.BindEnv("domain", "DOMAIN")
	_ = viper.BindEnv("key_path", "KEY_PATH")
	_ = viper.BindEnv("server_port", "SERVER_PORT")
	_ = viper.BindEnv("frontend_url", "FRONTEND_URL")

	// Unmarshal into struct
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal("Failed to unmarshal config:", err)
	}

	validate := validator.New()
	err = validate.StructCtx(context.Background(), Config)
	if err != nil {
		log.Fatal("Validation failed:", err)
	}
}
