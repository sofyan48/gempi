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

// ListQueue list all queue
// @svc: *sqs.SQS
// @inQueue: *sqs.ListQueuesInput
func (aw *Aws) ListQueue(svc *sqs.SQS, inQueue *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	result, err := svc.ListQueues(inQueue)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChangeVisibile ...
func (aw *Aws) ChangeVisibile() {}

// Delete ...
func (aw *Aws) Delete(svc *sqs.SQS, handlerMsg *sqs.Message, topic string) (*sqs.DeleteMessageOutput, error) {
	dels := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(topic),
		ReceiptHandle: handlerMsg.ReceiptHandle,
	}
	data, err := svc.DeleteMessage(dels)
	return data, err
}

// GetMessagesInput ...
func (aw *Aws) GetMessagesInput() *sqs.SendMessageInput {
	return &sqs.SendMessageInput{}
}

// Send ...
func (aw *Aws) Send(session *sqs.SQS, msgInput *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	sendResponse, err := session.SendMessage(msgInput)
	if err != nil {
		return sendResponse, err
	}
	return sendResponse, nil
}
