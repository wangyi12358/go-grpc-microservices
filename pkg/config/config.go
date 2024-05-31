package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Conf struct {
}

var Config Conf

func Setup() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	fmt.Printf("env: %s\n", env)
	flag.Parse()
	viper.SetEnvPrefix(env)
	viper.AutomaticEnv()
	viper.SetConfigName("config." + env)
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if viper.Unmarshal(&Config) != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
