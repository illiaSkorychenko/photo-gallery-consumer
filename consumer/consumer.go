package consumer

import (
	"fmt"
	"log"
	"photo-gallery-consumer/consumer/types"
	"photo-gallery-consumer/pkg/helper"
	"photo-gallery-consumer/pkg/service/compressor"
	"photo-gallery-consumer/pkg/service/image_processor"
)

const DefaultTimeout = 5

type Consumer struct {
	storage       types.ExternalStorage
	messageBroker types.MessageBroker
}

func New(storage types.ExternalStorage, messageBroker types.MessageBroker) *Consumer {
	return &Consumer{
		storage,
		messageBroker,
	}
}

func (c Consumer) Start() {
	for {
		err := c.process()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c Consumer) process() error {
	messageMap, handler, err := c.messageBroker.ReceiveMessage(DefaultTimeout)
	if err != nil {
		return err
	}

	if messageMap == nil {
		return nil
	}

	var parsedMessage types.ImageAttributes
	if err := helper.DecodeStringMap(messageMap, &parsedMessage); err != nil {
		return err
	}

	fmt.Println(parsedMessage)
	key := fmt.Sprintf("%s.%s", parsedMessage.Id, parsedMessage.Resolution)
	imgReader, err := c.storage.GetFile(key)
	if err != nil {
		return err
	}

	img, err := image_processor.Decode(imgReader)
	if err != nil {
		return err
	}
	// TODO
	compressedImages := compressor.Compress(img)

	if err := c.messageBroker.DeleteMessage(handler); err != nil {
		return err
	}

	return nil
}
