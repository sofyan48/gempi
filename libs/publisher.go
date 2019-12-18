package libs

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sofyan48/gempi/libs/connection"
)

type Publisher struct {
	DB     gorm.DB
	Broker *BrokerData
}

type BrokerData struct {
	MessageID string `gorm:"column:messageid;type:varchar(50)" json:"messageid"`
	Data      string `gorm:"column:data;type:string" json:"data"`
	Status    int    `gorm:"column:status;type:int unsigned" json:"status"`
}

func saveToBroker(db gorm.DB, brokerData *BrokerData) error {
	query := db.Table("broker")
	query = query.Create(brokerData)
	query.Scan(&brokerData)
	return query.Error
}

func Publis() {
	godotenv.Load()
	pubs := &Publisher{}
	pubs.DB = *connection.GetConnection()
	// broker

	// aws sqs
	QueueURL := os.Getenv("SQS_URL")
	awsAccessKeyID := os.Getenv("ACCESS_KEY")
	awsSecretAccessKey := os.Getenv("SECRET_KEY")
	creds := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")

	_, err := creds.Get()

	cfg := aws.NewConfig().WithRegion("ap-southeast-1").WithCredentials(creds)
	svc := sqs.New(session.New(), cfg)
	// Send message
	data := `{
		"action": "bca-va-payment",
		"topic": "payment",
		"parameter": [{
			"name": "va_number",
			"value": "123456"
		}],
		"body": {
			"type": "json",
			"value": {
				
			}
		}
	}`

	sendParams := &sqs.SendMessageInput{
		MessageBody:  aws.String(data),     // Required
		QueueUrl:     aws.String(QueueURL), // Required
		DelaySeconds: aws.Int64(3),         // (optional) 傳進去的 message 延遲 n 秒才會被取出, 0 ~ 900s (15 minutes)
	}
	sendResponse, err := svc.SendMessage(sendParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	broker := &BrokerData{}
	broker.MessageID = *sendResponse.MessageId
	broker.Data = data
	// err = saveToBroker(pubs.DB, broker)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }
	fmt.Println("Data : ", sendResponse)
}
