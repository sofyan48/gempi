package api

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sofyan48/gempi/entity"
	"github.com/sofyan48/gempi/libs"
)

// Producer ...
type Producer struct {
	session *sqs.SQS
	config  *entity.AwsConfig
	awsLibs *libs.Aws
	awsPubs *libs.Pubs
}

// NewProducer pubs Data
// @client: *entity.NewClient
// return *Producer
func NewProducer(client *entity.NewClient) *Producer {
	pubs := &Producer{}
	pubs.config = client.Configs
	pubs.session = client.Sessions
	pubs.awsLibs = &libs.Aws{}
	pubs.awsPubs = &libs.Pubs{}
	return pubs
}

// GetMessageInput statefull
// return *entity.StateFullModels
func (pubs *Producer) GetMessageInput() *entity.StateFullModels {
	return &entity.StateFullModels{}
}

// Send sending to sqs
// @topic: string
// @data: *entity.StateFullModels
// return *sqs.SendMessageOutput, error
func (pubs *Producer) Send(data *entity.StateFullModels) (*sqs.SendMessageOutput, error) {
	messages := pubs.awsPubs.GetMessagesInput()
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	messages.MessageBody = aws.String(string(body))
	messages.QueueUrl = aws.String(pubs.config.PathURL + "/" + pubs.config.Backend)
	messages.DelaySeconds = aws.Int64(3)
	result, err := pubs.awsPubs.Send(pubs.session, messages)
	return result, err
}
