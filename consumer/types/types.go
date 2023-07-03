package types

import (
	"io"
	"photo-gallery-consumer/pkg/service/compressor/types"
)

type ExternalStorage interface {
	UploadFile(buf []byte, key string) error
	GetFile(key string) (io.Reader, error)
	DeleteFile(key string) error
}

type MessageBroker interface {
	ReceiveMessage(timeout int32) (map[string]string, *string, error)
	DeleteMessage(handler *string) error
}

type ImageAttributes struct {
	Id         string                 `mapstructure:"id"`
	Resolution types.CompressionLevel `mapstructure:"resolution"`
	Format     string                 `mapstructure:"format"`
}
