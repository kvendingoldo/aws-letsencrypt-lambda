package main

import (
	"context"
	"fmt"
	awsLambda "github.com/aws/aws-lambda-go/lambda"
	cfg "github.com/kvendingoldo/aws-letsencrypt-lambda/internal/config"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/lambda"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/types"
	log "github.com/sirupsen/logrus"
	"os"
)

func Handler(ctx context.Context, event types.Event) (types.Response, error) {
	log.Infof("Handling labmda for event: %v", event)
	config := cfg.New(event)
	lambda.Execute(*config)
	return types.Response{Message: fmt.Sprintf("Labmda has been completed for %v\n", event.ID)}, nil
}

func main() {
	log.Info("Starting lambda execution ...")
	if mode, ok := os.LookupEnv("MODE"); ok {
		if !(mode == "local" || mode == "cloud") {
			log.Errorf("Environment variable 'MODE' has unknown value '%v'. Value should be 'local' or 'cloud'", mode)
			os.Exit(1)
		}

		if mode == "local" {
			config := cfg.New(nil)
			lambda.Execute(*config)
		} else if mode == "cloud" {
			awsLambda.Start(Handler)
		}
	}
	log.Info("Lambda has completed")
}
