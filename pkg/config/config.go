package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Service struct {
	Port string
	Name string
}

type Services struct {
	User Service
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Type     string
}

type Conf struct {
	Database Database
	Services Services
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
	fmt.Print("Config loaded\n")
}
