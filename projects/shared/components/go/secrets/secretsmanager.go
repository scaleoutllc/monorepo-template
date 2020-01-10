package secrets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// AWS holds an instance of a secretsmanager client.
type AWSSecretStore struct {
	client *secretsmanager.SecretsManager
}

// NewAWSSecretStore is a factory function that returns an instance of
// AWSSecretStore.
func NewAWSSecretStore() Store {
	return &AWSSecretStore{
		secretsmanager.New(session.New()),
	}
}

func (s *AWSSecretStore) get(key string) (string, error) {
	result, err := s.client.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(key),
	})
	if err != nil {
		return "", err
	}
	return *result.SecretString, nil
}
