# GEMPI
Implementing Publisher and Consumer with SQS
## Getting Started

### Installing
```
go get github.com/sofyan48/gempi
```
### Getting Client
```golang
// configure aws creds
cfg := config.Configure()
cfg.PathURL = os.Getenv("SQS_URL")
cfg.AwsAccessKeyID = os.Getenv("ACCESS_KEY")
cfg.AwsSecretAccessKey = os.Getenv("SECRET_KEY")
cfg.APArea = "ap-southeast-1"
// get sqs client
client := config.NewConfig().Credential(cfg).New()
```
### Producer
```golang
// create sqs Producer
producer := api.NewProducer(client)
// Publish Messages
message := producer.GetMessageInput()
message.Topic = "send1"
message.Status = "progres"
message.Body = "dataBody"
message.Parameter = "dataParams"
result, err := producer.Send(message)
if err != nil {
	fmt.Println("ERROR : ", err)
}
fmt.Println("Messages Sending : ", result)
```
### Consumer
***Callback***
```golang
func callbackData(results string) {
	obj := &entity.StateFullModels{}
	json.Unmarshal([]byte(results), &obj)
	fmt.Println("Data : ", obj)
}
```
now get consume
```golang
// Create Consumer
consumer := api.NewConsumer(client)
// consumer get data with callback
consumer.Consumer("send", callbackData, 1)
```