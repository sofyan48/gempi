package libs

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQSession get session sqs
// @url: string
// return *sqs.SQS
func (aw *Aws) SQSession() *sqs.SQS {
	awsAccessKeyID := os.Getenv("ACCESS_KEY")
	awsSecretAccessKey := os.Getenv("SECRET_KEY")
	creds := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")
	creds.Get()
	cfg := aws.NewConfig().WithRegion("ap-southeast-1").WithCredentials(creds)
	svc := sqs.New(session.New(), cfg)
	return svc
}

// Send messages to sqs
// @svc: *sqs.SQS
// @msgInput: *sqs.SendMessageInput
// return *sqs.SendMessageOutput, error
func (aw *Aws) Send(svc *sqs.SQS, msgInput *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	sendResponse, err := svc.SendMessage(msgInput)
	if err != nil {
		return sendResponse, err
	}
	return sendResponse, nil
}

// Recieved consume messages from sqs
// @svc: *sqs.SQS
// @receiveparams: *sqs.ReceiveMessageInput
// return *sqs.ReceiveMessageOutput, error
func (aw *Aws) Recieved(svc *sqs.SQS, receiveparams *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	messages, err := svc.ReceiveMessage(receiveparams)
	if err != nil {
		return messages, err
	}
	return messages, nil
}
