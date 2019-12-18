package api

import (
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

// Publish ...
func (pubs *Publisher) Publish(topic, data string) (*sqs.SendMessageOutput, error) {
	messages := pubs.awsPubs.GetMessagesInput()
	messages.MessageBody = aws.String(data)
	messages.QueueUrl = aws.String(pubs.config.PathURL)
	messages.DelaySeconds = aws.Int64(3)
	result, err := pubs.awsPubs.Send(pubs.session, messages)
	if err != nil {
		return result, err
	}
	return result, nil
}
