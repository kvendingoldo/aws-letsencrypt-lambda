package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	Region            string
	DomainName        string
	DomainOnly        bool
	DryRun            bool
	Mode              string
	ReImportThreshold int
	AcmeUrl           string
	AcmeEmail         string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func New() *Config {
	var config = Config{}

	if region := getEnv("AWS_REGION", "us-east-1"); region == "us-east-1" {
		log.Warnf("Required environment variable 'AWS_REGION' is empty. Default value %v will be used", region)
	} else {
		config.Region = region
	}

	if domainOnlyEnv := getEnv("DOMAIN_ONLY", ""); domainOnlyEnv == "" {
		log.Warnf("Environment variable 'DOMAIN_ONLY' is empty; Default value 'false' will be used")
		config.DomainOnly = false
	} else {
		domainOnly, err := strconv.ParseBool(domainOnlyEnv)
		if err != nil {
			log.Error(fmt.Sprintf("Could not parse DOMAIN_ONLY"), "error", err)
			os.Exit(1)
		}
		config.DomainOnly = domainOnly
	}

	if domain := getEnv("DOMAIN_NAME", ""); domain == "" {
		if config.DomainOnly {
			log.Error("Required environment variable 'DOMAIN' is empty. Please, specify")
			os.Exit(1)
		}
	} else {
		config.DomainName = domain
	}

	// TODO: delete it
	if dryRunEnv := getEnv("DRY_RUN", ""); dryRunEnv == "" {
		log.Warnf("Environment variable 'DRY_RUN' is empty; Default value 'false' will be used")
		config.DryRun = false
	} else {
		dryRun, err := strconv.ParseBool(dryRunEnv)
		if err != nil {
			log.Error(fmt.Sprintf("Could not parse DRY_RUN"), "error", err)
			os.Exit(1)
		}
		config.DryRun = dryRun
	}

	if mode := getEnv("MODE", "local"); mode == "local" {
		log.Infof("Environment variable 'MODE' is empty. Default value %v will be used", mode)
		config.Mode = mode
	} else if mode == "cloud" {
		config.Mode = mode
	} else {
		log.Errorf("Environment variable 'MODE' has unknown value '%v'. Value should be 'local' or 'cloud'", mode)
		os.Exit(1)
	}

	if useProdUrlEnv := getEnv("USE_PROD_URL", ""); useProdUrlEnv == "" {
		log.Warnf("Environment variable 'USE_PROD_URL' is empty; Default value 'false' will be used")
		config.AcmeUrl = "https://acme-staging-v02.api.letsencrypt.org/directory"
	} else {
		useProdUrl, err := strconv.ParseBool(useProdUrlEnv)
		if err != nil {
			log.Error(fmt.Sprintf("Could not parse USE_PROD_URL"), "error", err)
			os.Exit(1)
		}
		if useProdUrl {
			config.AcmeUrl = "https://acme-v02.api.letsencrypt.org/directory"
		} else {
			config.AcmeUrl = "https://acme-staging-v02.api.letsencrypt.org/directory"

		}
	}

	if acmeEmail := getEnv("ACME_EMAIL", ""); acmeEmail == "" {
		log.Error("Required environment variable 'ACME_EMAIL' is empty. Please, specify")
		os.Exit(1)
	} else {
		config.AcmeEmail = acmeEmail
	}

	config.ReImportThreshold = 10

	return &config
}
