package aws

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"photo-gallery-consumer/pkg/config"
)

func configS3(cfg aws.Config) *S3Storage {
	return newS3(cfg, config.AppConfig.S3Buket)
}

type S3Storage struct {
	buket  string
	client *s3.Client
}

func newS3(cfg aws.Config, buket string) *S3Storage {
	return &S3Storage{
		buket:  buket,
		client: s3.NewFromConfig(cfg),
	}
}

func (s S3Storage) UploadFile(buf []byte, key string) error {
	reader := bytes.NewReader(buf)
	input := &s3.PutObjectInput{
		Bucket: &s.buket,
		Key:    &key,
		Body:   reader,
	}
	_, err := s.client.PutObject(context.TODO(), input)

	return err
}

func (s S3Storage) GetFile(key string) (io.Reader, error) {
	input := &s3.GetObjectInput{
		Bucket: &s.buket,
		Key:    &key,
	}
	output, err := s.client.GetObject(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return output.Body, err
}

func (s S3Storage) DeleteFile(key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: &s.buket,
		Key:    &key,
	}
	_, err := s.client.DeleteObject(context.TODO(), input)

	return err
}
