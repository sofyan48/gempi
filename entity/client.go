package entity

import "github.com/aws/aws-sdk-go/service/sqs"

// NewClient ...
type NewClient struct {
	Sessions *sqs.SQS
	Configs  *AwsConfig
}
