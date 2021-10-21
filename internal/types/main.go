package types

type Response struct {
	Message string `json:"Answer:"`
}

type Event struct {
	ID                string `json:"id"`
	DomainName        string `json:"domain_name"`
	ReImportThreshold int    `json:"reimport_threshold"`
	AcmeUrl           string `json:"acme_url"`
	AcmeEmail         string `json:"acme_email"`
	IssueType         string `json:"issue_type"`
}
