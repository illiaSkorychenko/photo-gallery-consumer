package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"photo-gallery-consumer/pkg/config"
	"photo-gallery-consumer/pkg/helper"
)

var defaultConfig aws.Config

func init() {
	if err := setConfig(); err != nil {
		panic(err)
	}
}

func makeConfig() (defaultConfig aws.Config, err error) {
	var opt []func(*awsConfig.LoadOptions) error

	if config.AppConfig.Env == config.Dev {
		opt = []func(*awsConfig.LoadOptions) error{
			awsConfig.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:               config.AppConfig.LocalAWSEndpoint,
						HostnameImmutable: true,
					}, nil
				})),
			awsConfig.WithDefaultRegion(config.AppConfig.AWSRegion),
		}
	} else {
		opt = []func(*awsConfig.LoadOptions) error{
			awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("AKID", "SECRET_KEY", "TOKEN")),
			awsConfig.WithDefaultRegion(config.AppConfig.AWSRegion),
		}
	}

	defaultConfig, err = awsConfig.LoadDefaultConfig(
		context.TODO(),
		opt...,
	)
	if err != nil {
		return
	}

	return
}

func setConfig() (err error) {
	helper.CheckAppConfig()
	defaultConfig, err = makeConfig()

	return
}
