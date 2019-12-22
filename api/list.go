package api

import (
	"github.com/sofyan48/gempi/entity"
	"github.com/sofyan48/gempi/libs"
)

// ListQueue ...
func ListQueue(cl *entity.NewClient) (*entity.ListOutput, error) {
	awsLibs := &libs.Aws{}
	list := &entity.ListOutput{}
	result, err := awsLibs.ListQueue(cl.Sessions, nil)
	if err != nil {
		return list, err
	}
	list.ListQueuesOutput = result
	return list, nil
}
