package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type configuration struct {
	Host     string
	Port     int
	LogLevel string
	Database struct {
		Driver string
		Host   string
	}
}

var Configuration configuration

func setDefaults() {
	viper.SetDefault("Host", "0.0.0.0")
	viper.SetDefault("Port", 8080)
	viper.SetDefault("LogLevel", "debug")

	viper.SetDefault("Database.Driver", "postgres")
	viper.SetDefault("Database.Host", "user='postgres' dbname='jaded' host='127.0.0.1' password='postgres' port='5432' sslmode='disable'")

}

func InitConfig() {
	setDefaults()

	viper.BindEnv("Host", "HOST")
	viper.BindEnv("Port", "PORT")
	viper.BindEnv("LogLevel", "LOGLEVEL")

	viper.BindEnv("Database.Driver", "DATABASE_DRIVER")
	viper.BindEnv("Database.Host", "DATABASE_HOST")

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(&Configuration); err != nil {
		log.Error("Error Unmarshaling Configuration: ", err)
	}

	logLevel, _ := log.ParseLevel(Configuration.LogLevel)
	log.SetLevel(logLevel)

	log.Infof("Configuration: %+v Initialized", Configuration)
}
