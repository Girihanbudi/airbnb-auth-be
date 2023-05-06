package env

import (
	"airbnb-auth-be/internal/pkg/env/config"
	"airbnb-auth-be/internal/pkg/log"
	"fmt"

	"github.com/spf13/viper"
)

const Instance string = "Env"

// global env declaration
var CONFIG config.Config

type Options struct {
	Path     string
	FileName string
	Ext      string
}

func NewDefaultOptions() Options {
	return Options{
		Path:     "./env",
		FileName: "config",
		Ext:      "yaml",
	}
}

func InitEnv(options Options) {
	log.Event(Instance, "reading config...")

	viper.AddConfigPath(options.Path)
	viper.SetConfigName(options.FileName)
	viper.SetConfigType(options.Ext)

	env := config.Config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(Instance, "failed to read config", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal(Instance, "failed to unmarshal config", err)
	}

	log.Event(Instance, fmt.Sprintf("using %s stage mode", env.Stage))
	CONFIG = env
}

func ProvideEnv() config.Config {
	return CONFIG
}
