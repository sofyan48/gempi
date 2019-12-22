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

// ReceiveMessagesInput ...
func (sub *Subs) ReceiveMessagesInput() *sqs.ReceiveMessageInput {
	return &sqs.ReceiveMessageInput{}
}

// Recieved consume messages from sqs
// @svc: *sqs.SQS
// @receiveparams: *sqs.ReceiveMessageInput
// return *sqs.ReceiveMessageOutput, error
func (sub *Subs) Recieved(svc *sqs.SQS, receiveparams *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	messages, err := svc.ReceiveMessage(receiveparams)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// Delete queue
// @svc: *sqs.SQS
// @handlerMsg: *sqs.Message
// @topic string
// return *sqs.DeleteMessageOutput, error
func (sub *Subs) Delete(svc *sqs.SQS, handlerMsg *sqs.Message, topic string) (*sqs.DeleteMessageOutput, error) {
	dels := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(topic),
		ReceiptHandle: handlerMsg.ReceiptHandle,
	}
	data, err := svc.DeleteMessage(dels)
	if err != nil {
		return nil, err
	}
	return data, nil
}
