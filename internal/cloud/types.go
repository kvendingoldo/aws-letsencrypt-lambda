package cloud


type Certificate struct {
	Arn                 string
	Region              string
	DomainName          string
	Type                string
	Status              string
	FailureReason       string
	ValidationMethod    string
}
