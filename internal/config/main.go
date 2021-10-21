package config

import (
	"fmt"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/types"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	Region            string
	DomainName        string
	ReImportThreshold int
	AcmeEmail         string
	AcmeUrl           string
	IssueType         string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func New(eventRaw interface{}) *Config {
	var config = Config{}
	var getFromEvent bool
	var event types.Event

	switch value := eventRaw.(type) {
	case types.Event:
		getFromEvent = true
		event = value
	default:
		getFromEvent = false
	}

	// Process Region
	if region := getEnv("AWS_REGION", ""); region == "" {
		log.Errorf("Required environment variable 'AWS_REGION' is empty. Please, specify", region)
		os.Exit(1)
	} else {
		config.Region = region
	}

	// Process DomainName
	domain := getEnv("DOMAIN_NAME", "")
	if domain == "" {
		log.Warnf("Environment variable 'DOMAIN_NAME' is empty")
	} else {
		config.DomainName = domain
	}

	if getFromEvent {
		if event.DomainName == "" {
			log.Warnf("Event contains empty DomainName")
			if domain == "" {
				log.Error("DomainName is empty; Configure it via 'DOMAIN_NAME' env variable OR pass in event body")
				os.Exit(1)
			}
		} else {
			config.DomainName = event.DomainName
		}
	}

	// Process ReImportThreshold
	if reimportThreshold := getEnv("REIMPORT_THRESHOLD", ""); reimportThreshold == "" {
		log.Warnf("Environment variable 'REIMPORT_THRESHOLD' is empty")
	} else {
		value, err := strconv.Atoi(reimportThreshold)
		if err != nil {
			log.Error(fmt.Sprintf("Could not parse 'REIMPORT_THRESHOLD' variable'"), "error", err)
			os.Exit(1)
		}
		config.ReImportThreshold = value
	}

	if getFromEvent && event.ReImportThreshold != 0 {
		config.ReImportThreshold = event.ReImportThreshold
	}

	if config.ReImportThreshold == 0 {
		log.Error("ReImportThreshold == 0 ; Configure non-zero value via 'ReImportThreshold' env variable OR pass in event body")
		os.Exit(1)
	}

	// Process AcmeEmail
	if acmeEmail := getEnv("ACME_EMAIL", ""); acmeEmail == "" {
		log.Warnf("Environment variable 'ACME_EMAIL' is empty")
	} else {
		config.AcmeEmail = acmeEmail
	}

	if getFromEvent && event.AcmeEmail != "" {
		config.AcmeEmail = event.AcmeEmail
	}

	if config.AcmeEmail == "" {
		log.Errorf("AcmeEmail is empty; Configure it via 'ACME_EMAIL' env variable OR pass in event body")
		os.Exit(1)
	}

	// Process AcmeUrl
	var acmeUrl string

	if acmeUrlEnv := getEnv("ACME_URL", ""); acmeUrlEnv == "" {
		log.Warnf("Environment variable 'ACME_URL' is empty")
	} else {
		acmeUrl = acmeUrlEnv
	}

	if getFromEvent && event.AcmeUrl != "" {
		acmeUrl = event.AcmeUrl
	}

	switch acmeUrl {
	case "prod":
		config.AcmeUrl = "https://acme-v02.api.letsencrypt.org/directory"
		log.Info("Lambda will use PROD ACME URL")
	case "stage":
		config.AcmeUrl = "https://acme-staging-v02.api.letsencrypt.org/directory"
		log.Info("Lambda will use STAGING ACME URL; If you need to use PROD URL specify it via 'ACME_URL' or pass in event body")
	default:
		log.Errorf("Unkown value '%v' for acmeUrl; Check env var 'ACME_URL' or event body", acmeUrl)
		os.Exit(1)
	}

	// Process Force
	if issueType := getEnv("ISSUE_TYPE", ""); issueType == "" {
		log.Warnf("Environment variable 'ISSUE_TYPE' is empty; 'default' value will be used")
		config.IssueType = "default"
	} else {
		config.IssueType = issueType
	}

	if getFromEvent && event.IssueType != "" {
		config.IssueType = event.IssueType
	}

	if !(config.IssueType == "default" || config.IssueType == "force") {
		log.Errorf("Bad IssueType value (%v). It should be 'default' or 'force'", config.IssueType)
		os.Exit(1)
	}

	return &config
}
