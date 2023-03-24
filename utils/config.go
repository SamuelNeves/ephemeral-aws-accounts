package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	ProgramName string `mapstructure:"PROGRAM_NAME"`
}

var lock = &sync.Mutex{}

var singleInstance *Config

func GetConfig() *Config {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = getConfig()
		}
	}
	return singleInstance
}
func getConfig() *Config {
	var cfg Config

	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())

	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())

	}
	return &cfg
}
