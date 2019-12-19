package entity

import "github.com/aws/aws-sdk-go/service/sqs"

// NewClient ...
type NewClient struct {
	Sessions *sqs.SQS
	Configs  *AwsConfig
}

// StateFullModels data modeling for input
type StateFullModels struct {
	Recovery Recovery `json:"recovery"`
	Body     string   `json:"body"`
}

// Recovery status data
type Recovery struct {
	VisibleTimeot int `json:"visible_timeout"`
	Count         int `json:"try_counter"`
}
