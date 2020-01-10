package secrets

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// AWS holds an instance of a secretsmanager client.
type AWSParameterStore struct {
	client *ssm.SSM
}

// NewAWSSecretStore is a factory function that returns an instance of
// AWSParameterStore.
func NewAWSParameterStore() Store {
	return &AWSParameterStore{
		ssm.New(session.New()),
	}
}

func (s *AWSParameterStore) get(key string) (string, error) {
	result, err := s.client.GetParameters(&ssm.GetParametersInput{
		Names:[]*string{aws.String(key)},
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", err
	}
	if len(result.Parameters) != 1 {
		return "", fmt.Errorf("parameter not found")
	}
	return *result.Parameters[0].Value, nil
}
