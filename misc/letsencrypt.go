package utils

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/cloud"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/config"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme"
	"os"
)

func getCertificate(ctx context.Context, client cloud.Client, config config.Config, domainName string) (string, error) {
	accountKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	acmeClient := &acme.Client{
		Key:          accountKey,
		DirectoryURL: config.AcmeUrl,
	}

	if _, err := acmeClient.Register(context.Background(), &acme.Account{},
		func(tos string) bool {
			log.Infof("Agreeing to ToS: %s", tos)
			return true
		}); err != nil {
		log.Fatal("Can't register an ACME account: ", err)
	}

	// Authorize a DNS name
	authz, err := acmeClient.Authorize(context.Background(), domainName)
	if err != nil {
		log.Fatal("Can't authorize: ", err)
	}

	// Find the DNS challenge for this authorization
	var challenge *acme.Challenge
	for _, c := range authz.Challenges {
		if c.Type == "dns-01" {
			challenge = c
			break
		}
	}
	if challenge == nil {
		log.Fatal("No DNS challenge was present")
	}




	// TODO
	//	err := cloud.ChangeRecord(ctx, *client, route53Types.ChangeActionDelete, *info.Certificate.DomainName , "_acme-challenge", "\"someToken\"")

	// Determine the TXT record values for the DNS challenge
	txtLabel := "_acme-challenge." + authz.Identifier.Value
	txtValue, _ := acmeClient.DNS01ChallengeRecord(chal.Token)
	log.Printf("Creating record %s with value %s", txtLabel, txtValue)

	// Then the usual: accept the challenge, wait for the authorization ...
	if _, err := acmeClient.Accept(context.Background(), chal); err != nil {
		log.Fatal("Can't accept challenge: ", err)
	}

	if _, err := acmeClient.WaitAuthorization(context.Background(), authz.URI); err != nil {
		log.Fatal("Failed authorization: ", err)
	}

	





	os.Exit(1)
	fmt.Println("STUB")
	return "", nil
}
