package utils

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/go-acme/lego/providers/dns/route53"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/config"
	"log"
	"time"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

type user struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *user) GetEmail() string {
	return u.Email
}
func (u user) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *user) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func GetCertificates(config config.Config, domainName string) (*certificate.Resource, error) {
	fmt.Println("+++")
	fmt.Println(domainName)


	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return &certificate.Resource{}, err
	}

	acmeUser := user{
		Email: config.AcmeEmail,
		key:   privateKey,
	}

	acmeConfig := lego.NewConfig(&acmeUser)

	acmeConfig.CADirURL = config.AcmeUrl
	acmeConfig.Certificate.KeyType = certcrypto.RSA2048

	// NOTE: A client facilitates communication with the CA server.
	acmeClient, err := lego.NewClient(acmeConfig)
	if err != nil {
		log.Fatal(err)
	}

	// NOTE: Set up route53 dns-01 challenge provider
	route53Config := route53.NewDefaultConfig()
	route53Config.PropagationTimeout = time.Second * 300
	route53Provider, err := route53.NewDNSProviderConfig(route53Config)
	if err != nil {
		log.Fatal(err)
	}
	err = acmeClient.Challenge.SetDNS01Provider(route53Provider)
	if err != nil {
		log.Fatal(err)
	}

	// NOTE: New users will need to register
	reg, err := acmeClient.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		log.Fatal(err)
	}
	acmeUser.Registration = reg

	request := certificate.ObtainRequest{
		Domains: []string{
			domainName,
			fmt.Sprintf("www.%v", domainName),
		},
		Bundle: true,
	}

	// NOTE: Each cert comes back with the cert bytes, the bytes of the client's private key, and a certificate URL
	certificates, err := acmeClient.Certificate.Obtain(request)
	if err != nil {
		log.Fatal(err)
	}

	return certificates, nil
}
