# GEMPI
Implementing Publisher and Consumer with SQS
## Getting Started

### Installing
```
go get github.com/sofyan48/gempi
```
### Getting Client
```
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
```
### Publisher
```
    // get sqs publisher
	publisher := api.NewPublisher(client)
	// Publish Messages
	result, err := publisher.Publish("test", "TEST")
	if err != nil {
		fmt.Println("ERROR : ", err)
	}
	fmt.Println(result)
```
### Consumer
```
On Progress
```