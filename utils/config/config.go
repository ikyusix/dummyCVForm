package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppName     string `mapstructure:"APP_NAME"`
	Environment string `mapstructure:"ENVIRONMENT"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Timeout     int    `mapstructure:"TIMEOUT"`
	Db          string `mapstructure:"DB"`
	Host        string `mapstructure:"HOST"`
}

var (
	MyConfig *Config
	Log      = logrus.New()
)

func init() {

	// force all writes to regular log to logger
	log.SetOutput(Log.Writer())
	log.SetFlags(0)

	// configure logging for environment
	Log.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		ForceQuote:    true,
		FullTimestamp: true,
	}

	Log.Println("Start Load Config")

	var err error
	MyConfig, err = LoadConfig(".")
	if err != nil {
		Log.Errorf("Load Config Failed : " + err.Error())
	}
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("local")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return config, nil
}
