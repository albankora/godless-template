package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	if len(sqsEvent.Records) == 0 {
		return errors.New("No SQS message passed to function")
	}

	for err, message := range sqsEvent.Records {

		if err != nil {
			return err
		}

		fmt.Printf(message)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
