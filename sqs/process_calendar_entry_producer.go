package sqs

import (
	"calendar-service/config"
	"calendar-service/repositories"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Producer handles the process of fetching entries, writing to SQS, and updating the last processed time.
func ProcessCalendarEntryProducer() error {
	
	sqsConfig, err := config.LoadSQSConfig()
	if err != nil {
		return err
	}
	queueURL := sqsConfig.Queues["calendar-entry"]
	lastProcessedTime, err := repositories.FetchLastProcessedTime()
	if err != nil {
		return err
	}

	entries, err := repositories.FetchCalendarEntriesSince(lastProcessedTime)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		body, _ := json.Marshal(entry)
		_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
			QueueUrl:    aws.String(queueURL),
			MessageBody: aws.String(string(body)),
		})
		if err != nil {
			log.Printf("Failed to send message to SQS: %v", err)
			return err
		}
	}

	log.Println("Messages sent to SQS")

	if len(entries) > 0 {
		lastUpdatedTime := entries[len(entries)-1].UpdatedAt
		err = repositories.UpdateLastProcessedTime(lastUpdatedTime)
		if err != nil {
			return err
		}
	}

	return nil
}
