package api

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sofyan48/gempi/entity"
	"github.com/sofyan48/gempi/libs"
)

// Publisher ...
type Publisher struct {
	session *sqs.SQS
	config  *entity.AwsConfig
	awsLibs *libs.Aws
	awsPubs *libs.Pubs
}

// NewPublisher pubs Data
func NewPublisher(client *entity.NewClient) *Publisher {
	pubs := &Publisher{}
	pubs.config = client.Configs
	pubs.session = client.Sessions
	pubs.awsLibs = &libs.Aws{}
	pubs.awsPubs = &libs.Pubs{}
	return pubs
}

// GetMessageInput ...
func (pubs *Publisher) GetMessageInput() *entity.StateFullModels {
	return &entity.StateFullModels{}
}

// Publish ...
func (pubs *Publisher) Publish(data *entity.StateFullModels) (*sqs.SendMessageOutput, error) {
	messages := pubs.awsPubs.GetMessagesInput()
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	messages.MessageBody = aws.String(string(body))
	messages.QueueUrl = aws.String(pubs.config.PathURL)
	messages.DelaySeconds = aws.Int64(3)
	result, err := pubs.awsPubs.Send(pubs.session, messages)
	if err != nil {
		return nil, err
	}
	return result, nil
}
