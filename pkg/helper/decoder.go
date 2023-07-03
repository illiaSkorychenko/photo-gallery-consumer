package helper

import (
	"github.com/mitchellh/mapstructure"
)

func DecodeStringMap(data map[string]string, output any) (err error) {
	if err = mapstructure.Decode(data, output); err != nil {
		return
	}

	return
}
