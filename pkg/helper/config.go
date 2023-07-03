package helper

import (
	"errors"
	"photo-gallery-consumer/pkg/config"
)

func CheckAppConfig() {
	if (config.AppConfig == config.Config{}) {
		panic(errors.New("app config is empty"))
	}
}
