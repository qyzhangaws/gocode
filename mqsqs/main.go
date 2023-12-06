package main

import (
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/sqs"
)

// list sqs listqueue function
func listQueues(sqsClient *client.Client) {
	result, err := sqsClient.ListQueues(nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// create sqs function
func createQueue(sqsClient *client.Client, queueName string) (*sqs.CreateQueueOutput, error) {
	result, err := sqsClient.CreateQueue(&sqs.CreateQueueInput{
		QueueName: queueName,
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("60"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})
	// snippet-end:[sqs.go.create_queue.call]
	if err != nil {
		return nil, err
	}

	return result, nil
}

// delete sqs function
func deleteQueue(sqsClient *client.Client, queueUrl string) (*sqs.CreateQueueOutput, error) {
	_, err := sqsClient.DeleteQueue(&sqs.DeleteQueueInput{
		QueueUrl: aws.String(queueUrl),
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// send message to sqs function
func GetQueueURL(sqsClient *client.Client, queueName *string) (*sqs.GetQueueUrlOutput, error) {
	result, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queueName,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// sendMsg to queue
func SendMsg(sqsClient *client.Client, msgInput *sqs.SendMessageInput) error {
	_, err := sqsClient.SendMessage(msgInput)

	if err != nil {
		return err
	}

	return nil
}

// recieve message from queue
func RecieveMsg(sqsClient *client.Client, queueUrl *string, visibilityTimeout *int64) (*sqs.ReceiveMessageOutput, error) {
	result, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueUrl,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   visibilityTimeout,
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
