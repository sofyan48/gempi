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
	// Queue List
	list, _ := api.ListQueue(client, nil)
	for _, i := range list.QueueUrls {
		fmt.Println(*i)
	}
}
