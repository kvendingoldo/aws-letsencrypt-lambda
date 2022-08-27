package types

import (
	"github.com/guregu/null"
)

type Response struct {
	Message string `json:"answer"`
}

type Event struct {
	AWSRegion string `json:"awsRegion"`

	ACMRegion     string `json:"acmRegion"`
	Route53Region string `json:"route53Region"`

	DomainName        string   `json:"domainName"`
	ReImportThreshold null.Int `json:"reImportThreshold"`
	AcmeURL           string   `json:"acmeUrl"`
	AcmeEmail         string   `json:"acmeEmail"`
	IssueType         string   `json:"issueType"`
}
