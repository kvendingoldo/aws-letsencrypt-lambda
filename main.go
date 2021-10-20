package main

import (
	"context"
	"fmt"
	awsLambda "github.com/aws/aws-lambda-go/lambda"
	cfg "github.com/kvendingoldo/aws-letsencrypt-lambda/internal/config"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/lambda"
)

type Response struct {
	Message string `json:"Answer:"`
}

type Event struct {
	// TODO: add options here
	Name string `json:"name"`
}

func Handler(ctx context.Context, event Event) (Response, error) {
	fmt.Println("Hello from lambda")
	// TODO: find a way to pass config from main()
	config := cfg.New()
	lambda.Execute(*config)
	return Response{Message: fmt.Sprintf("Hello %v\n", event.Name)}, nil
}

func main() {
	config := cfg.New()

	if config.Mode == "local" {
		lambda.Execute(*config)
	} else if config.Mode == "cloud" {
		awsLambda.Start(Handler)
	}
}
