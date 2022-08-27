## Environment variables

* FORMATTER_TYPE
    * Description: formatter type for Lambda's logs
    * Possible values: JSON | TEXT
    * Required: no

* MODE
    * Description: application mode
    * Possible values: cloud | local
    * Required: yes

* LOG_LEVEL
    * Description: Lambda's log level
    * Possible values: panic|fatal|error|warn|info|debug|trace
    * Required: no

* AWS_REGION
    * Description: Default AWS Region. Inside of Lambda it's setting automatically by AWS
    * Possible values: <any valid AWS region>
    * Required: yes

* DOMAIN_NAME
    * Description: Name of domain for which certificate will be issued/renewed
    * Possible values: *any valid domain name*

* ACME_URL
    * Description: If prod then *production* LE URL will be used, otherwise *stage* URL will be used
    * Possible values: prod | stage

* ACME_EMAIL
    * Description: Email that will be associated with LE certificate
    * Possible values: *any valid email*

* REIMPORT_THRESHOLD
    * Description: If TTL of cert == REIMPORT_THRESHOLD then cert will be renewed
    * Possible values: *any int > 0*
