package types

type Response struct {
	Message string `json:"answer"`
}

type Event struct {
	AWSRegion string `json:"awsRegion"`

	ACMRegion     string `json:"acmRegion"`
	Route53Region string `json:"route53Region"`

	DomainName        string `json:"domainName"`
	ReImportThreshold int    `json:"reimportThreshold"`
	AcmeUrl           string `json:"acmeUrl"`
	AcmeEmail         string `json:"acmeEmail"`
	IssueType         string `json:"issueType"`
}
