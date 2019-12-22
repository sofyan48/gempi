package api

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sofyan48/gempi/entity"
	"github.com/sofyan48/gempi/libs"
)

// Consumer ...
type Consumer struct {
	session   *sqs.SQS
	config    *entity.AwsConfig
	awsLibs   *libs.Aws
	awsSubs   *libs.Subs
	MsgOutput *sqs.ReceiveMessageOutput
	Worker    int
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

// SetWorker ...
// @worker: int
func (csm *Consumer) SetWorker(worker int) *Consumer {
	csm.Worker = worker
	return csm
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
func (csm *Consumer) Subscribe(topic string, cb func(*entity.Context, *entity.NewClient)) {
	if csm.Worker == 0 {
		log.Println("Setting Default Worker", 1)
		csm.SetWorker(1)
	}
	messages := csm.awsSubs.ReceiveMessagesInput()
	messages.QueueUrl = aws.String(csm.config.PathURL + "/broker")
	messages.MaxNumberOfMessages = aws.Int64(3)
	messages.VisibilityTimeout = aws.Int64(30)
	messages.WaitTimeSeconds = aws.Int64(20)
	var wg sync.WaitGroup
	wg.Add(1)
	for w := 1; w <= csm.Worker; w++ {
		go csm.worker(w, messages, topic, cb)
	}
	wg.Wait()

}

func (csm *Consumer) worker(id int, messages *sqs.ReceiveMessageInput, topic string, cb func(*entity.Context, *entity.NewClient)) {
	callbackContext := &entity.Context{}
	callbackDataContext := &entity.StateFullModels{}
	for {
		output, err := csm.awsSubs.Recieved(csm.session, messages)
		if err != nil {
			continue
		}
		var wg sync.WaitGroup
		for _, message := range output.Messages {
			data := &entity.StateFullModels{}
			json.Unmarshal([]byte(*message.Body), &data)
			if topic == data.Topic && data.Status == "onBroker" {
				wg.Add(1)
				go func(m *sqs.Message, topic string, cb func(*entity.Context, *entity.NewClient), data *entity.StateFullModels) {
					defer wg.Done()
					callbackContext.Message = m
					callbackContext.StateFull = callbackDataContext
					callbackContext.Data = data
					clients := &entity.NewClient{}
					clients.Sessions = csm.session
					clients.Configs = csm.config
					callback(cb, callbackContext, clients)
				}(message, topic, cb, data)
				wg.Wait()
			}
		}
	}
}

// callback processing function
// @fn: interface{}
// @data: interface{}
func callback(fn interface{}, context interface{}, client interface{}) {
	switch fn.(type) {
	case func(string):
		fn.(func(string))(context.(string))
	case func(*entity.Context, *entity.NewClient):
		fn.(func(*entity.Context, *entity.NewClient))(context.(*entity.Context), client.(*entity.NewClient))
	default:
		fmt.Println("default")
	}
}
