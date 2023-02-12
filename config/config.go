package config

import (
	"fmt"

	"github.com/alireza-mf/go-ecommerce/util"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	MongoURI          string `mapstructure:"MONGODB_URI"`
	MongoDatabaseName string `mapstructure:"MONGODB_DATABASE_NAME"`
}

var config Config

func InitConfig() {
	viper.AddConfigPath("../")
	viper.SetConfigFile(".env")
	viper.WatchConfig()

	loadConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Env file has changed.")
		loadConfig()
	})
}

func loadConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config = Config{
		Port:              getVariable("PORT", "8000"),
		MongoURI:          getVariable("MONGODB_URI", "mongodb://localhost:27017/"),
		MongoDatabaseName: getVariable("MONGODB_DATABASE_NAME", "ecommerce"),
	}
}

func getVariable(envConfig string, defaultConfig string) string {
	envValue := viper.GetString(envConfig)
	return util.If(envValue == "", envValue, defaultConfig)
}

func GetConfig() Config {
	return config
}
