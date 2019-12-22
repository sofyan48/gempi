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

	// Create Consumer
	consumer := api.NewConsumer(client)
	// consumer.SetWorker()
	consumer.Subscribe(callback)

}

func callback(context *entity.Context) {
	obj := &entity.StateFullModels{}
	json.Unmarshal([]byte(*context.Message.Body), &obj)
	fmt.Println("Message Raw From Context", obj)
}
