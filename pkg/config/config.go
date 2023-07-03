package config

import (
	"github.com/spf13/viper"
)

type EnvType string

const (
	Dev  EnvType = "dev"
	Prod EnvType = "prod"
	Test EnvType = "test"
)

type Config struct {
	Env              EnvType `mapstructure:"ENV"`
	LocalAWSEndpoint string  `mapstructure:"LOCAL_AWS_ENDPOINT"`
	AWSRegion        string  `mapstructure:"AWS_REGION"`
	S3Buket          string  `mapstructure:"S3_BUCKET"`
	SQSQueue         string  `mapstructure:"SQS_QUEUE"`
}

var AppConfig Config

func init() {
	err := LoadConfig(".")
	if err != nil {
		panic(err)
	}
}

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&AppConfig)

	return
}
