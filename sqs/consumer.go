package sqs

import (
	"calendar-service/models"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func StartPollingQueue(queueURL string) {
    for {
        output, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
            QueueUrl:            aws.String(queueURL),
            MaxNumberOfMessages: aws.Int64(10),
            WaitTimeSeconds:     aws.Int64(20), // Long polling for 20 seconds
        })
        if err != nil {
            log.Printf("Failed to receive messages from SQS: %v", err)
            continue
        }

        if len(output.Messages) == 0 {
            log.Println("No messages received in this poll, continuing...")
            continue
        }

        for _, message := range output.Messages {
            var entry models.CalendarEntry
            err = json.Unmarshal([]byte(*message.Body), &entry)
            if err != nil {
                log.Printf("Failed to unmarshal message: %v", err)
                continue
            }

            log.Printf("Received message: %+v\n", entry)

            // Delete the message after processing
            _, err = sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
                QueueUrl:      aws.String(queueURL),
                ReceiptHandle: message.ReceiptHandle,
            })
            if err != nil {
                log.Printf("Failed to delete message from SQS: %v", err)
            }
        }
    }
}