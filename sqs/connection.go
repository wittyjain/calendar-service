package sqs

import (
	"calendar-service/config"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var sqsClient *sqs.SQS

// InitSQSClient initializes the SQS client with the provided configuration
func InitSQSClient() {
	sqsConfig, err := config.LoadSQSConfig()
	if err != nil {
		log.Fatalf("Failed to load SQS config: %v", err)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(sqsConfig.Region),
		Endpoint:    aws.String(sqsConfig.Endpoint),
		Credentials: credentials.NewStaticCredentials(sqsConfig.Credentials.AccessKey, sqsConfig.Credentials.SecretKey, ""),
	}))
	sqsClient = sqs.New(sess)
	err = createQueue("calendar-entry")
	if err != nil {
		log.Fatalf("Failed to create SQS queue: %v", err)
	}
}

// GetClient returns the initialized SQS client
func GetClient() *sqs.SQS {
	return sqsClient
}

func createQueue(queueURL string) error {
	_, err := sqsClient.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(queueURL), // Use the queue URL as the queue name
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == sqs.ErrCodeQueueNameExists {
			log.Println("Queue already exists")
			return nil
		}
		return err
	}

	log.Println("Queue created successfully")
	return nil
}
