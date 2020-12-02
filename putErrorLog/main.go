package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type SQSEvent events.SQSEvent

func Handler(ctx context.Context, sqsEvent SQSEvent) error {
	var project *string
	var messageBody *string
	var status *string
	for _, message := range sqsEvent.Records {
		project = message.MessageAttributes["project"].StringValue
		messageBody = message.MessageAttributes["message"].StringValue
		status = message.MessageAttributes["status"].StringValue
	}
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	))
	svc := dynamodb.New(mySession)
	inputItem := &dynamodb.PutItemInput{
		TableName: aws.String("faultline"),
		Item: map[string]*dynamodb.AttributeValue{
			"project": {
				S: project,
			},
			"message": {
				S: messageBody,
			},
			"status": {
				S: status,
			},
		},
	}
	svc.PutItem(inputItem)
	return nil
}

func main() {
	lambda.Start(Handler)
}
