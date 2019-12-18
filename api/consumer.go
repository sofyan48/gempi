package api

import (
	"encoding/json"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sofyan48/gempi/entity"
	"github.com/sofyan48/gempi/libs"
)

// Consumer ...
type Consumer struct {
	session *sqs.SQS
	config  *entity.AwsConfig
	awsLibs *libs.Aws
	awsSubs *libs.Subs
}

// NewConsumer pubs Data
func NewConsumer(client *entity.NewClient) *Consumer {
	consumer := &Consumer{}
	consumer.config = client.Configs
	consumer.session = client.Sessions
	consumer.awsLibs = &libs.Aws{}
	consumer.awsSubs = &libs.Subs{}
	return consumer
}

// GetMessageOutput ...
func (csm *Consumer) GetMessageOutput() *entity.StateFullModels {
	return &entity.StateFullModels{}
}

// Consumer ...
func (csm *Consumer) Consumer(topic string, cb func(string), delta int) {
	messages := csm.awsSubs.ReceiveMessagesInput()
	messages.QueueUrl = aws.String(csm.config.PathURL)
	messages.MaxNumberOfMessages = aws.Int64(3)
	messages.VisibilityTimeout = aws.Int64(30)
	messages.WaitTimeSeconds = aws.Int64(20)
	var wg sync.WaitGroup
	wg.Add(delta)
	go func() {
		for {
			msg, err := csm.awsSubs.Recieved(csm.session, messages)
			if err != nil {
				continue
			}
			for _, message := range msg.Messages {
				msgData := &entity.StateFullModels{}
				json.Unmarshal([]byte(*message.Body), &msgData)
				if msgData.Topic != topic {
					continue
				}
				callback(cb, *message.Body)
			}
		}
	}()
	wg.Wait()
}

func callback(fn interface{}, data interface{}) {
	switch fn.(type) {
	case func(string):
		fn.(func(string))(data.(string))
	case func(int):
		fn.(func(int))(data.(int))
	}
}
