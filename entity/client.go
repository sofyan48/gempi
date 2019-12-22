package entity

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

// NewClient ...
type NewClient struct {
	Sessions *sqs.SQS
	Configs  *AwsConfig
}

// Context ...
type Context struct {
	Message   *sqs.Message
	StateFull *StateFullModels
}

// StateFullModels data modeling for input
type StateFullModels struct {
	Status   string   `json:"status"`
	Recovery Recovery `json:"recovery"`
	Body     string   `json:"body"`
}

// Recovery status data
type Recovery struct {
	VisibleTimeot int `json:"visible_timeout"`
	Count         int `json:"try_counter"`
}

// ListOutput ...
type ListOutput struct {
	*sqs.ListQueuesOutput
}

// ListInput ...
type ListInput struct {
	*sqs.ListQueuesInput
}
