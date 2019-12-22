package api

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sofyan48/gempi/entity"
	"github.com/sofyan48/gempi/libs"
)

// Callback ...
type Callback struct{}

// GetCallbackFunction ...
func GetCallbackFunction() *Callback {
	return &Callback{}
}

// Flush messages
// @msg: *sqs.Message
// @data: *entity.StateFullModels
func (clb *Callback) Flush(cl *entity.NewClient, msg *sqs.Message, data *entity.StateFullModels) {
	dataSend := &entity.StateFullModels{}
	dataSend.Status = "done"
	dataSend.Body = data.Body
	dataSend.Topic = data.Topic
	body, _ := json.Marshal(dataSend)
	awsLibs := &libs.Aws{}

	msgBroker := awsLibs.GetMessagesInput()
	msgBroker.QueueUrl = aws.String(cl.Configs.PathURL + "/broker")
	msgBroker.DelaySeconds = aws.Int64(3)
	msgBroker.MessageBody = aws.String(string(body))
	_, err := awsLibs.Send(cl.Sessions, msgBroker)
	if err != nil {
		log.Println("Error : ", err)
	}
	_, err = awsLibs.Delete(cl.Sessions, msg, cl.Configs.PathURL+"/broker")
	log.Println("Status Done : ", *msg.Body)
}

// Deleted messages
// @msg: *sqs.Message
// @data: *entity.StateFullModels
func (clb *Callback) Deleted(cl *entity.NewClient, msg *sqs.Message) {
	awsLibs := &libs.Aws{}
	_, err := awsLibs.Delete(cl.Sessions, msg, cl.Configs.PathURL+"/broker")
	if err != nil {
		log.Println("Error : ", err)
	}
	log.Println("Status Done : ", *msg.Body)
}
