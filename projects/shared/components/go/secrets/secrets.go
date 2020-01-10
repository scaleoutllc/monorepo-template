package secrets

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Handler struct {
	Protocol string
	Store Store
}

// Store defines the interface required to fetch secrets.
type Store interface {
	get(string) (string, error)
}

// Get returns a secret either defined insecurely and directly in the
// environment or from a provided secret store at a user defined path in
// the environment. If both a path and explicit value are provided, the
// explicit value trumps.
func Get(name string, handlers ...Handler) (string, error) {
	protocols := make([]string, len(handlers))
	for i, item := range handlers {
		protocols[i] = item.Protocol
	}
	log.Printf("secret system initialized with support for protocol(s): %s", strings.Join(protocols, ", "))

	valueEnv := "APP_SECRET_VALUE_" + name
	pathEnv := "APP_SECRET_PATH_" + name
	secretValue := os.Getenv(valueEnv)
	secretPath := os.Getenv(pathEnv)

	// when neither path or value are provided, error our
	if secretValue == "" && secretPath == "" {
		return "", fmt.Errorf("expected entry in the environment for %s or %s", pathEnv, valueEnv)
	}

	// when both secret path and value are defined, value trumps
	if secretValue != "" && secretPath != "" {
		log.Printf("using secret value provided in the environment at %s, ignoring %s", valueEnv, pathEnv)
		return secretValue, nil
	}

	// when only value is defined, use it
	if secretValue != "" {
		log.Printf("using secret value explicitly provided in the environment at %s", valueEnv)
		return secretValue, nil
	}

	// if we made here, only the secretPath is set, let's go look up the value
	// find the appropriate handler for the secret being requested
	var handler Handler
	var secretLookup string
	for _, item := range handlers {
		if strings.HasPrefix(secretPath, item.Protocol) {
			handler = item
			secretLookup = strings.TrimPrefix(secretPath, item.Protocol)
			break;
		}
	}
	// if no handler was found for the specified protocol, bail out
	if handler == (Handler{}) {
		return "", fmt.Errorf("unable to locate protocol handler for %s=%s", pathEnv, secretPath)
	}

	// if the path was a bare protocol handler (e.g. APP_SECRET_PATH=aws-ssm:),
	// bail out, there is nothing to find
	if secretLookup == "" {
		return "", fmt.Errorf("no secret path specified for %s=%s", pathEnv, secretPath)
	}

	// if we made it here we're ready to actually request a secret from some
	// external store.
	value, err := handler.Store.get(secretLookup)
	if err != nil {
		return "", fmt.Errorf("%s secret at path \"%s\" retrieval failed using %s (%s): %s", pathEnv, secretLookup, handler.Protocol, handler.Store, err)
	}
	log.Printf("%s secret at path \"%s\" retrieved successfully using %s (%T)", pathEnv, secretLookup, handler.Protocol, handler.Store)
	return value, nil
}
