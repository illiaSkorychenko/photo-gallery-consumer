package main

import (
	"photo-gallery-consumer/consumer"
	"photo-gallery-consumer/consumer/types"
	"photo-gallery-consumer/pkg/client/aws"
	_ "photo-gallery-consumer/pkg/config"
)

func main() {
	var externalStorage types.ExternalStorage = aws.GetExternalStorage()
	var messageBroker types.MessageBroker = aws.GetMessageBroker()

	cons := consumer.New(externalStorage, messageBroker)
	cons.Start()
}
