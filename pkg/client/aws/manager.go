package aws

import (
	"photo-gallery-consumer/pkg/config"
	"photo-gallery-consumer/pkg/helper"
)

func GetExternalStorage() *S3Storage {
	checkDefaultConfig()

	return configS3(defaultConfig)
}

func GetMessageBroker() *SQSBroker {
	helper.CheckAppConfig()
	checkDefaultConfig()

	messageBroker, err := configSQS(defaultConfig, config.AppConfig.SQSQueue)
	if err != nil {
		panic(err)
	}

	return messageBroker
}
