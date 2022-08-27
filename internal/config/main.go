package config

import (
	"fmt"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/types"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	AWSRegion string

	ACMRegion     string
	Route53Region string

	DomainName        string
	ReImportThreshold int64
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

func New(eventRaw interface{}) (*Config, error) {
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

	// Process AWSRegion
	if awsRegion := getEnv("AWS_REGION", ""); awsRegion != "" {
		config.AWSRegion = awsRegion
	} else {
		log.Warn("Environment variable AWS_REGION is empty")
	}
	if getFromEvent {
		if event.AWSRegion != "" {
			config.AWSRegion = event.AWSRegion
		} else {
			log.Warn("Event contains empty awsRegion variable")
		}
	}
	if config.AWSRegion == "" {
		return nil, fmt.Errorf("AWSRegion is empty; Configure it via 'AWS_REGION' env variable OR pass in event body")
	}

	// Process DomainName
	if domain := getEnv("DOMAIN_NAME", ""); domain != "" {
		config.DomainName = domain
	} else {
		log.Warn("Environment variable 'DOMAIN_NAME' is empty")
	}
	if getFromEvent {
		if event.DomainName == "" {
			log.Warnf("Event contains empty domainName variable")
		} else {
			config.DomainName = event.DomainName
		}
	}
	if event.DomainName == "" {
		return nil, fmt.Errorf("DomainName is empty; Configure it via 'DOMAIN_NAME' env variable OR pass in event body")
	}

	// Process ReImportThreshold
	if reimportThreshold := getEnv("REIMPORT_THRESHOLD", ""); reimportThreshold == "" {
		log.Warn("Environment variable 'REIMPORT_THRESHOLD' is empty")
	} else {
		reimportThresholdValue, err := strconv.ParseInt(reimportThreshold, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Could not parse 'REIMPORT_THRESHOLD' variable'. Error: %s", err)
		}
		config.ReImportThreshold = reimportThresholdValue
	}
	if getFromEvent {
		if event.ReImportThreshold.Valid {
			config.ReImportThreshold = event.ReImportThreshold.Int64
		} else {
			log.Warnf("Event contains empty OR invalid reImportThreshold")
		}
	}
	if config.ReImportThreshold == 0 {
		return nil, fmt.Errorf("ReImportThreshold == 0; Configure non-zero value via 'REIMPORT_THRESHOLD' env variable OR pass in event body")
	}

	// Process AcmeEmail
	if acmeEmail := getEnv("ACME_EMAIL", ""); acmeEmail == "" {
		log.Warn("Environment variable 'ACME_EMAIL' is empty")
	} else {
		config.AcmeEmail = acmeEmail
	}
	if getFromEvent {
		if event.AcmeEmail != "" {
			config.AcmeEmail = event.AcmeEmail
		} else {
			log.Warn("Event contains empty acmeEmail")
		}
	}
	if config.AcmeEmail == "" {
		return nil, fmt.Errorf("AcmeEmail is empty; Configure it via 'ACME_EMAIL' env variable OR pass in event body")
	}

	// Process AcmeUrl
	var acmeUrl string
	if acmeUrlEnv := getEnv("ACME_URL", ""); acmeUrlEnv == "" {
		log.Warn("Environment variable 'ACME_URL' is empty")
	} else {
		acmeUrl = acmeUrlEnv
	}
	if getFromEvent {
		if event.AcmeUrl != "" {
			acmeUrl = event.AcmeUrl
		} else {
			log.Warn("Event contains empty acmeUrl variable")
		}
	}

	switch acmeUrl {
	case "prod":
		config.AcmeUrl = "https://acme-v02.api.letsencrypt.org/directory"
		log.Info("Lambda will use PROD ACME URL")
	case "stage":
		config.AcmeUrl = "https://acme-staging-v02.api.letsencrypt.org/directory"
		log.Info("Lambda will use STAGING ACME URL; If you need to use PROD URL specify it via 'ACME_URL' or pass in event body")
	default:
		return nil, fmt.Errorf("Unkown value '%v' for acmeUrl; Check env var 'ACME_URL' or event body; Valid values are: 'stage' or 'prod'", acmeUrl)
	}

	// Process Force
	if issueType := getEnv("ISSUE_TYPE", ""); issueType == "" {
		log.Warnf("Environment variable 'ISSUE_TYPE' is empty")
		config.IssueType = "default"
	} else {
		config.IssueType = issueType
	}
	if getFromEvent {
		if event.IssueType != "" {
			config.IssueType = event.IssueType
		} else {
			log.Warn("Event contains empty issueType")
		}
	}
	if config.IssueType == "" {
		config.IssueType = "default"
		log.Info("IssueType is empty; 'default' value will be used")
	}
	if !(config.IssueType == "default" || config.IssueType == "force") {
		return nil, fmt.Errorf("Bad IssueType value (%v). It should be 'default' or 'force'", config.IssueType)
	}

	return &config, nil
}
