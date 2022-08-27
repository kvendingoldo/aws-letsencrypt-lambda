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
	AcmeURL           string
	IssueType         string
}

//nolint:unparam
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
			//nolint:stylecheck
			return nil, fmt.Errorf("Could not parse 'REIMPORT_THRESHOLD' variable'. Error: %w", err)
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

	// Process AcmeURL
	var acmeURL string
	if acmeURLEnv := getEnv("ACME_URL", ""); acmeURLEnv == "" {
		log.Warn("Environment variable 'ACME_URL' is empty")
	} else {
		acmeURL = acmeURLEnv
	}
	if getFromEvent {
		if event.AcmeURL != "" {
			acmeURL = event.AcmeURL
		} else {
			log.Warn("Event contains empty acmeUrl variable")
		}
	}

	switch acmeURL {
	case "prod":
		config.AcmeURL = "https://acme-v02.api.letsencrypt.org/directory"
		log.Info("Lambda will use PROD ACME URL")
	case "stage":
		config.AcmeURL = "https://acme-staging-v02.api.letsencrypt.org/directory"
		log.Info("Lambda will use STAGING ACME URL; If you need to use PROD URL specify it via 'ACME_URL' or pass in event body")
	default:
		//nolint:stylecheck
		return nil, fmt.Errorf("Unknown value '%v' for acmeUrl; Check env var 'ACME_URL' or event body; Valid values are: 'stage' or 'prod'", acmeURL)
	}

	// Process Force
	if issueType := getEnv("ISSUE_TYPE", ""); issueType == "" {
		log.Warnf("Environment variable 'ISSUE_TYPE' is empty")
		config.IssueType = types.IssueTypeDefault
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
		config.IssueType = types.IssueTypeDefault
		log.Infof("IssueType is empty; '%s' value will be used", types.IssueTypeDefault)
	}
	if !(config.IssueType == types.IssueTypeDefault || config.IssueType == types.IssueTypeForce) {
		//nolint:stylecheck
		return nil, fmt.Errorf("Bad IssueType value '%v'. It should be '%s' or '%s'", config.IssueType, types.IssueTypeDefault, types.IssueTypeForce)
	}

	return &config, nil
}
