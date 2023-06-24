package main

import (
	"money-keeper/cmd"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	cmd.Run()
}
