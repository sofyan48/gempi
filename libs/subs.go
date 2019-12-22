package libs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// GetMessagesInput ...
func (sub *Subs) GetMessagesInput() *sqs.SendMessageInput {
	return &sqs.SendMessageInput{}
}

// Send messages to sqs
// @svc: *sqs.SQS
// @msgInput: *sqs.SendMessageInput
// return *sqs.SendMessageOutput, error
func (sub *Subs) Send(session *sqs.SQS, msgInput *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
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
	return data, err
}
