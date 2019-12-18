package libs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sofyan48/gempi/entity"
)

// SQSession get session sqs
// @url: string
// return *sqs.SQS
func (aw *Aws) SQSession(cfg *entity.AwsConfig) *sqs.SQS {
	creds := credentials.NewStaticCredentials(cfg.AwsAccessKeyID, cfg.AwsSecretAccessKey, "")
	creds.Get()
	cfgAws := aws.NewConfig().WithRegion(cfg.APArea).WithCredentials(creds)
	svc := sqs.New(session.New(), cfgAws)
	return svc
}

// GetMessagesInput ...
func (pubs *Pubs) GetMessagesInput() *sqs.SendMessageInput {
	return &sqs.SendMessageInput{}
}

// Send messages to sqs
// @svc: *sqs.SQS
// @msgInput: *sqs.SendMessageInput
// return *sqs.SendMessageOutput, error
func (pubs *Pubs) Send(session *sqs.SQS, msgInput *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	sendResponse, err := session.SendMessage(msgInput)
	if err != nil {
		return sendResponse, err
	}
	return sendResponse, nil
}

// GetMessagesOutput ...
func (sub *Subs) GetMessagesOutput() *sqs.SendMessageOutput {
	return &sqs.SendMessageOutput{}
}

// Recieved consume messages from sqs
// @svc: *sqs.SQS
// @receiveparams: *sqs.ReceiveMessageInput
// return *sqs.ReceiveMessageOutput, error
func (sub *Subs) Recieved(svc *sqs.SQS, receiveparams *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	messages, err := svc.ReceiveMessage(receiveparams)
	if err != nil {
		return messages, err
	}
	return messages, nil
}
