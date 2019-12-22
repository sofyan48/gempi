package api

import (
	"fmt"
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
// @client: *entity.NewClient
// return *Consumer
func NewConsumer(client *entity.NewClient) *Consumer {
	consumer := &Consumer{}
	consumer.config = client.Configs
	consumer.session = client.Sessions
	consumer.awsLibs = &libs.Aws{}
	consumer.awsSubs = &libs.Subs{}
	return consumer
}

// GetMessageOutput statefull
// return *entity.StateFullModels
func (csm *Consumer) GetMessageOutput() *entity.StateFullModels {
	return &entity.StateFullModels{}
}

// Delete ...
func (csm *Consumer) Delete(handler *sqs.Message, topic string) (*sqs.DeleteMessageOutput, error) {
	result, err := csm.awsSubs.Delete(csm.session, handler, topic)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Subscribe topic
// @topic: string
// @cb: func(string) // callback
// @delta: int
func (csm *Consumer) Subscribe(topic string, cb func(*entity.Context), delta int) {
	messages := csm.awsSubs.ReceiveMessagesInput()
	messages.QueueUrl = aws.String(csm.config.PathURL + "/" + topic)
	messages.MaxNumberOfMessages = aws.Int64(3)
	messages.VisibilityTimeout = aws.Int64(30)
	messages.WaitTimeSeconds = aws.Int64(20)
	callbackContext := &entity.Context{}
	callbackDataContext := &entity.StateFullModels{}
	var wg sync.WaitGroup
	wg.Add(delta)
	go func() {
		for {
			msg, err := csm.awsSubs.Recieved(csm.session, messages)
			if err != nil {
				continue
			}
			for _, message := range msg.Messages {
				callbackContext.Message = message
				callbackContext.StateFull = callbackDataContext
				callback(cb, callbackContext)
			}
		}
	}()
	wg.Wait()
}

// callback processing function
// @fn: interface{}
// @data: interface{}
func callback(fn interface{}, context interface{}) {
	switch fn.(type) {
	case func(string):
		fn.(func(string))(context.(string))
	case func(*entity.Context):
		fn.(func(*entity.Context))(context.(*entity.Context))
	default:
		fmt.Println("default")
	}
}
