package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig(path string) {
	viper.SetConfigFile(path + ".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error config file: %w \n", err))
	}
}
