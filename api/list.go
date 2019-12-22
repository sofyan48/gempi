package api

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sofyan48/gempi/entity"
	"github.com/sofyan48/gempi/libs"
)

// ListQueue ...
func ListQueue(cl *entity.NewClient, inQueue *entity.ListInput) (*entity.ListOutput, error) {
	var input *sqs.ListQueuesInput
	if inQueue != nil {
		input = inQueue.ListQueuesInput
	}
	awsLibs := &libs.Aws{}
	list := &entity.ListOutput{}
	result, err := awsLibs.ListQueue(cl.Sessions, input)
	list.ListQueuesOutput = result
	return list, err
}
