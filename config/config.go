package config

import (
	"fmt"
	"log"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Web      WebConfig
	Logger   Log
	Datadog  Datadog
	UserList SheetConfig `yaml:"sheet"`
}

type WebConfig struct {
	Host    string
	Port    string
	LogFile string
	Datadog Datadog `yaml:"datadog"`
}

type SheetConfig struct {
	Port                 string
	LogFile              string
	UsersSheetID         string
	Datadog              Datadog `yaml:"datadog"`
	GrpcClientEndPoint   string
	GoogleAppCredential  map[string]interface{} `yaml:"googleAppCredential"`
	GoogleUserCredential GoogleAuthToken        `yaml:"googleUserCredential"`
}

type GoogleAuthToken struct {
	AccessToken  string `yaml:"access_token"`
	TokenType    string `yaml:"token_type,omitempty"`
	RefreshToken string `yaml:"refresh_token,omitempty"`
	Expiry       string `yaml:"expiry,omitempty"`
}

type Log struct {
	LogLevel string
	LogPath  string
}

type Datadog struct {
	Env     string `yaml:"env"`
	Service string `yaml:"service"`
	Version string `yaml:"version"`
}

func LoadConfig(filename, path string) *Config {
	var configuration Config
	var configName string
	configName = "config"
	//	var path string
	var configPath = "."
	if filename != "" {
		configName = filename
	}

	if path != "" {
		configPath = path
	}

	//For Local Config
	if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
		viper.SetConfigName(configName)
		// Set the path to look for the configurations file
		viper.AddConfigPath(configPath)
		// Enable VIPER to read Environment Variables
		viper.AutomaticEnv()
		viper.SetConfigType("yml")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("error reading config file")
			log.Fatalf("Error reading config file, %s", err.Error())
		}
		err := viper.UnmarshalExact(&configuration)
		if err != nil {
			fmt.Println("error decoding config file")
			//	log.Errorf("Unable to decode into struct, %v", err)
			log.Fatalf("Unable to decode into struct, %v", err)
		}
		return &configuration
	}
	//For deployment config
	viper.SetConfigName(configName)
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	viper.SetConfigFile("config.yml")
	err := viper.MergeInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = viper.UnmarshalExact(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return &configuration
}
