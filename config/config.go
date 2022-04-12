package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf *viper.Viper

func init() {
	Conf = viper.New()
	Conf.SetConfigName("conf")
	Conf.AddConfigPath(".")
	Conf.SetConfigType("yaml")
	if err := Conf.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
