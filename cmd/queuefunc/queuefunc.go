package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.SQSEvent) error {

	if len(event.Records) == 0 {
		return errors.New("No SQS message passed to function")
	}

	for _, message := range event.Records {
		fmt.Printf(message.Body)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
