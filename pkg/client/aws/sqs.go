package aws

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSBroker struct {
	query    string
	client   *sqs.Client
	queryUrl *string
}

func configSQS(cfg aws.Config, query string) (*SQSBroker, error) {
	brokerClient := newSQS(cfg, query)

	if err := brokerClient.setQueryUrl(); err != nil {
		return nil, err
	}

	return brokerClient, nil
}

func newSQS(cfg aws.Config, query string) *SQSBroker {
	client := sqs.NewFromConfig(cfg)

	return &SQSBroker{
		query:  query,
		client: client,
	}
}

func (b *SQSBroker) setQueryUrl() error {
	var err error
	b.queryUrl, err = b.getQueryUrl()
	if err != nil {
		return err
	}

	return nil
}

func (b *SQSBroker) checkQueryUrl() error {
	if b.queryUrl == nil {
		return errors.New("query URL was not set")
	}

	return nil
}

func (b *SQSBroker) getQueryUrl() (*string, error) {
	input := &sqs.GetQueueUrlInput{
		QueueName: &b.query,
	}

	output, err := b.client.GetQueueUrl(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output.QueueUrl, err
}

func (b *SQSBroker) ReceiveMessage(timeout int32) (res map[string]string, handler *string, err error) {
	if err = b.checkQueryUrl(); err != nil {
		return
	}

	input := &sqs.ReceiveMessageInput{
		QueueUrl:              b.queryUrl,
		MessageAttributeNames: []string{".*"},
		VisibilityTimeout:     timeout,
	}
	receiveMessageOutput, err := b.client.ReceiveMessage(context.TODO(), input)

	if len(receiveMessageOutput.Messages) == 0 {
		return
	}

	message := receiveMessageOutput.Messages[0]
	handler = message.ReceiptHandle

	res = make(map[string]string)
	for key, attribute := range message.MessageAttributes {
		res[key] = *attribute.StringValue
	}

	return
}

func (b *SQSBroker) DeleteMessage(handler *string) (err error) {
	if err = b.checkQueryUrl(); err != nil {
		return
	}

	input := &sqs.DeleteMessageInput{
		QueueUrl:      b.queryUrl,
		ReceiptHandle: handler,
	}
	_, err = b.client.DeleteMessage(context.TODO(), input)

	return
}
