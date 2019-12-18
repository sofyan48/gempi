package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sofyan48/gempi/api"
	"github.com/sofyan48/gempi/config"
	"github.com/sofyan48/gempi/entity"
)

func callbackData(results string) {
	obj := &entity.StateFullModels{}
	json.Unmarshal([]byte(results), &obj)
	fmt.Println("Data : ", obj)
}

func main() {
	// load dotenv
	godotenv.Load()
	// configure aws creds
	cfg := config.Configure()
	cfg.PathURL = os.Getenv("SQS_URL")
	cfg.AwsAccessKeyID = os.Getenv("ACCESS_KEY")
	cfg.AwsSecretAccessKey = os.Getenv("SECRET_KEY")
	cfg.APArea = "ap-southeast-1"
	// get sqs client
	client := config.NewConfig().Credential(cfg).New()

	// create sqs Producer
	producer := api.NewProducer(client)
	// Publish Messages
	message := producer.GetMessageInput()
	message.Topic = "send"
	message.Status = "progres"
	message.Body = "dataBody"
	message.Parameter = "dataParams"
	result, err := producer.Send(message)
	if err != nil {
		fmt.Println("ERROR : ", err)
	}
	fmt.Println("Messages Sending : ", result)

	// Create Consumer
	consumer := api.NewConsumer(client)
	// consumer get data with callback
	consumer.Consumer("send", callbackData, 1)

}
