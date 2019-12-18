package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sofyan48/gempi/api"
	"github.com/sofyan48/gempi/config"
)

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
	// get sqs publisher
	publisher := api.NewPublisher(client)
	// Publish Messages
	message := publisher.GetMessageInput()
	message.Topic = "send"
	message.Status = "progres"
	message.Body = "data"
	message.Parameter = "data"
	result, err := publisher.Publish(message)
	if err != nil {
		fmt.Println("ERROR : ", err)
	}
	fmt.Println(result)

}
