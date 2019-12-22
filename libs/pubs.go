package libs

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

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
