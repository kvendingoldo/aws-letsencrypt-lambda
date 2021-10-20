package main

import (
	"context"
	awsLambda "github.com/aws/aws-lambda-go/lambda"
	cfg "github.com/kvendingoldo/aws-letsencrypt-lambda/internal/config"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/lambda"
)

type MyResponse struct {
	Message string `json:"Answer:"`
}

func Handler(ctx context.Context) (string, error) {
	// TODO
	//config := cfg.New()
	//lambda.Execute()
	//return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
	return "test", nil
}

func main() {
	config := cfg.New()

	if config.Mode == "local" {
		lambda.Execute(*config)
	} else if config.Mode == "cloud" {
		awsLambda.Start(Handler)
	}
}
