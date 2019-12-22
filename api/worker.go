package api

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sofyan48/gempi/entity"
	"github.com/sofyan48/gempi/libs"
)

// Worker ...
type Worker struct {
	session   *sqs.SQS
	config    *entity.AwsConfig
	awsLibs   *libs.Aws
	awsSubs   *libs.Subs
	MsgOutput *sqs.ReceiveMessageOutput
	Worker    int
}

// NewWorker pubs Data
// @client: *entity.NewClient
// return *Worker
func NewWorker(client *entity.NewClient) *Worker {
	Worker := &Worker{}
	Worker.config = client.Configs
	Worker.session = client.Sessions
	Worker.awsLibs = &libs.Aws{}
	Worker.awsSubs = &libs.Subs{}
	return Worker
}

// SetWorker ...
// @worker: int
func (wrk *Worker) SetWorker(worker int) *Worker {
	wrk.Worker = worker
	return wrk
}

// Start worker
// @topic: string
// @cb: func(string) // callback
// @delta: int
func (wrk *Worker) Start() {
	if wrk.Worker == 0 {
		log.Println("Setting Default Worker", 1)
		wrk.SetWorker(1)
	}
	messages := wrk.awsSubs.ReceiveMessagesInput()
	messages.QueueUrl = aws.String(wrk.config.PathURL + "/backend")
	messages.MaxNumberOfMessages = aws.Int64(3)
	messages.VisibilityTimeout = aws.Int64(30)
	messages.WaitTimeSeconds = aws.Int64(20)
	var wg sync.WaitGroup
	wg.Add(1)
	for w := 1; w <= wrk.Worker; w++ {
		go wrk.worker(w, messages)
	}
	wg.Wait()

}

func (wrk *Worker) worker(id int, messages *sqs.ReceiveMessageInput) {
	for {
		output, err := wrk.awsSubs.Recieved(wrk.session, messages)
		if err != nil {
			continue
		}
		var wg sync.WaitGroup
		for _, message := range output.Messages {
			wg.Add(1)
			go func(m *sqs.Message) {
				defer wg.Done()
				data := &entity.StateFullModels{}
				json.Unmarshal([]byte(*m.Body), &data)

				dataSend := &entity.StateFullModels{}
				dataSend.Status = "onBroker"
				dataSend.Body = data.Body
				dataSend.Topic = data.Topic
				dataSend.ID = *m.MessageId
				body, _ := json.Marshal(dataSend)
				msgBroker := wrk.awsSubs.GetMessagesInput()
				msgBroker.QueueUrl = aws.String(wrk.config.PathURL + "/" + wrk.config.Broker)
				msgBroker.DelaySeconds = aws.Int64(3)
				msgBroker.MessageBody = aws.String(string(body))
				_, err := wrk.awsSubs.Send(wrk.session, msgBroker)
				if err != nil {
					log.Println("Broker : ", err)
				} else {
					_, err = wrk.awsSubs.Delete(wrk.session, m, wrk.config.PathURL+"/"+wrk.config.Backend)
					if err != nil {
						log.Println("Backend Error : ", err)
					} else {
						log.Println("ID : ", *m.MessageId, " | Data :", dataSend.Body)
					}

				}

			}(message)

			wg.Wait()
		}
	}
}
