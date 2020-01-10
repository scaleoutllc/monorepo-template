package main

import (
	"log"
	"org/shared/components/go/secrets"
)

func main() {
	secretsmanager := secrets.Handler{
		"aws-secretsmanager:",
		secrets.NewAWSSecretStore(),
	}
	ssm := secrets.Handler{
		"aws-ssm:",
		secrets.NewAWSParameterStore(),
	}
	_, err := secrets.Get("TEST", secretsmanager, ssm)
	if err != nil {
		log.Fatal(err)
	}
	// do something with the secret
}
