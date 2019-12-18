package subs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sofyan48/gempi/libs/connection"
)

type consumer struct{}

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BodyValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BodyStruct struct {
	Type  string    `json:"type"`
	Value BodyValue `json:"value"`
}

type Recovery struct {
	VisibleTimeot int `json:"visible_timeout"`
	Count         int `json:"try_counter"`
}

type StateFullModels struct {
	Topic     string      `json:"topic"`
	Status    string      `json:"status"`
	Recovery  Recovery    `json:"recovery"`
	Parameter []Parameter `json:"parameter"`
	Body      BodyStruct  `json:"body"`
}

type Publisher struct {
	DB     gorm.DB
	Broker *BrokerData
}

type BrokerData struct {
	MessageID string `gorm:"column:messageid;type:varchar(50)" json:"messageid"`
	Data      string `gorm:"column:data;type:string" json:"data"`
	Status    int    `gorm:"column:status;type:int unsigned" json:"status"`
}

func checkBroker(db gorm.DB, id string, broker *BrokerData) error {
	query := db.Table("broker")
	query = query.Where("messageid=?", id)
	query = query.First(broker)
	return query.Error
}

func updateBroker(db gorm.DB, id string, broker *BrokerData) error {
	query := db.Table("broker")
	query = query.Where("messageid=?", id)
	query = query.Updates(broker)
	query.Scan(&broker)
	return query.Error
}

func deleteBroker(db gorm.DB, id string, broker *BrokerData) error {
	query := db.Table("broker")
	query = query.Where("messageid=?", id)
	query = query.Delete(broker, query)
	return query.Error
}

func actionCallback() bool {
	return true
}

// DeleteMsg ...
func DeleteMsg(svc *sqs.SQS, handlerMsg *sqs.Message, url string) {
	dels := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(url),          // Required
		ReceiptHandle: handlerMsg.ReceiptHandle, // Required
	}
	_, err := svc.DeleteMessage(dels) // No response returned when successed.
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Message ID: %s has beed deleted", *handlerMsg.MessageId)
}

func saveToBroker(db gorm.DB, brokerData *BrokerData) error {
	query := db.Table("broker")
	query = query.Create(brokerData)
	query.Scan(&brokerData)
	return query.Error
}

func SendMsg(svc *sqs.SQS, url string, data *StateFullModels) string {
	dataSend, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	sendParams := &sqs.SendMessageInput{
		MessageBody:  aws.String(string(dataSend)), // Required
		QueueUrl:     aws.String(url),              // Required
		DelaySeconds: aws.Int64(3),                 // (optional) 傳進去的 message 延遲 n 秒才會被取出, 0 ~ 900s (15 minutes)
	}
	sendResponse, err := svc.SendMessage(sendParams)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data : ", sendResponse)
	return *sendResponse.MessageId
}
func ConsumeMessages() {
	godotenv.Load()
	pubs := &Publisher{}
	pubs.DB = *connection.GetConnection()
	// aws
	QueueURL := os.Getenv("SQS_URL")
	awsAccessKeyID := os.Getenv("ACCESS_KEY")
	awsSecretAccessKey := os.Getenv("SECRET_KEY")
	creds := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")

	creds.Get()

	cfg := aws.NewConfig().WithRegion("ap-southeast-1").WithCredentials(creds)
	svc := sqs.New(session.New(), cfg)
	// Receive message
	receiveparams := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(QueueURL),
		MaxNumberOfMessages: aws.Int64(3),
		VisibilityTimeout:   aws.Int64(30),
		WaitTimeSeconds:     aws.Int64(20),
	}
	var wgRot sync.WaitGroup
	wgRot.Add(10)
	go func() {
		for {

			output, err := svc.ReceiveMessage(receiveparams)
			if err != nil {
				continue
			}
			for _, message := range output.Messages {
				data := &StateFullModels{}
				brokers := &BrokerData{}
				json.Unmarshal([]byte(*message.Body), &data)
				fmt.Println("Msg ID: ", *message.MessageId)
				// checkBroker(pubs.DB, *message.MessageId, brokers)
				check := actionCallback()
				fmt.Println(check)
				if check {
					fmt.Println("Eksek Gagal ")
					brokers.Status = 1
					data.Status = "trying"
					data.Recovery.VisibleTimeot = 12
					data.Recovery.Count = 5
					DeleteMsg(svc, message, QueueURL)
					msgID := SendMsg(svc, QueueURL, data)
					brokers.MessageID = msgID
					updateBroker(pubs.DB, *message.MessageId, brokers)
					continue
				}
				fmt.Println("Eksek Berhasil ")
				brokers.Status = 2
				data.Recovery.Count = 0
				data.Status = "done"
				updateBroker(pubs.DB, *message.MessageId, brokers)
				DeleteMsg(svc, message, QueueURL)

				// next push data to paymentFail
			}
			time.After(time.Second * 10)
		}
	}()
	wgRot.Wait()
}
