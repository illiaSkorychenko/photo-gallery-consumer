package aws

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"reflect"
)

func checkDefaultConfig() {
	if reflect.DeepEqual(defaultConfig, aws.Config{}) {
		panic(errors.New("aws default config is empty"))
	}
}
