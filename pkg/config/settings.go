package config

import (
	"github.com/spf13/viper"
	"log"
)

//Settings represent the environment configuration of this application.
type Settings struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

//New instance of Settings type.
func New() *Settings {
	s := &Settings{}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No env file found.", err)
	}

	//try to assign read variables into golang struct
	err = viper.Unmarshal(&s)
	if err != nil {
		log.Fatal("Error trying to unmarshal configuration.", err)
	}

	return s
}
