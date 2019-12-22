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
	cfg.Backend = "backend"
	cfg.Broker = "broker"
	// get sqs client
	client := config.NewConfig().Credential(cfg).New()

	// create sqs Producer
	producer := api.NewProducer(client)
	// Publish Messages
	message := producer.GetMessageInput()
	message.Topic = "payment"
	message.Body = "GEMPI 3"
	result, err := producer.Send(message)
	if err != nil {
		fmt.Println("ERROR : ", err)
	}
	fmt.Println("Messages Sending : ", result)

}
