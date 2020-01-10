package secrets

import (
	"fmt"
	"os"
	"strings"
	"testing"
)
var errMockSecretMissing = fmt.Errorf("unable to locate secret")

type MockSecretStoreOne struct{}
func (s *MockSecretStoreOne) get(path string) (string, error) {
	if path == "missing/secret/path" {
		return "", errMockSecretMissing
	}
	return fmt.Sprintf("%s-located-in-secret-store-one", path), nil
}

type MockSecretStoreTwo struct{}
func (s *MockSecretStoreTwo) get(path string) (string, error) {
	if path == "missing/secret/path" {
		return "", errMockSecretMissing
	}
	return fmt.Sprintf("%s-located-in-secret-store-two", path), nil
}

func TestGet(t *testing.T) {
	tests := map[string]struct {
		handlers     []Handler
		key          string
		pathContent  string
		valueContent string
		want         string
		wantErr      error
	}{
		"neither environment set": {
			[]Handler{
				{"mock-one://", &MockSecretStoreOne{}},
				{"mock-two://", &MockSecretStoreTwo{}},
			},
			"MISSING",
			"",
			"",
			"",
			fmt.Errorf("expected entry in the environment for"),
		},
		"both environments set, value trumps": {
			[]Handler{
				{"mock-one://", &MockSecretStoreOne{}},
				{"mock-two://", &MockSecretStoreTwo{}},
			},
			"BOTH",
			"mock://test/secret/path",
			"test",
			"test",
			nil,
		},
		"only value environment set": {
			[]Handler{
				{"mock-one://", &MockSecretStoreOne{}},
				{"mock-two://", &MockSecretStoreTwo{}},
			},
			"VALUEONLY",
			"",
			"testSecretValue",
			"testSecretValue",
			nil,
		},
		"only path environment set, matching mock handler one, success": {
			[]Handler{
				{"mock-one://", &MockSecretStoreOne{}},
				{"mock-two://", &MockSecretStoreTwo{}},
			},
			"PATHONLYMOCKONE",
			"mock-one://test/secret/path",
			"",
			"test/secret/path-located-in-secret-store-one",
			nil,
		},
		"only path environment set, matching mock handler two, success": {
			[]Handler{
				{"mock-one://", &MockSecretStoreOne{}},
				{"mock-two://", &MockSecretStoreTwo{}},
			},
			"PATHONLYMOCKTWO",
			"mock-two://test/secret/path",
			"",
			"test/secret/path-located-in-secret-store-two",
			nil,
		},
		"only path environment set, matching handler found, retrieval error": {
			[]Handler{
				{"mock-one://", &MockSecretStoreOne{}},
				{"mock-two://", &MockSecretStoreTwo{}},
			},
			"BADPATH",
			"mock-one://missing/secret/path",
			"",
			"",
			fmt.Errorf("unable to locate secret"),
		},
		"only path environment set, no protocol handler found": {
			[]Handler{
				{"mock-one://", &MockSecretStoreOne{}},
				{"mock-two://", &MockSecretStoreTwo{}},
			},
			"NOHANDLER",
			"secret://missing/handler",
			"",
			"",
			fmt.Errorf("unable to locate protocol handler for"),
		},
		"only path environment set, protocol handler found, no path specified": {
			[]Handler{
				{"mock-one://", &MockSecretStoreOne{}},
				{"mock-two://", &MockSecretStoreTwo{}},
			},
			"HASHANDLERNOPATH",
			"mock-one://",
			"",
			"",
			fmt.Errorf("no secret path specified for"),
		},
	}
	os.Clearenv()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer os.Clearenv()
			os.Setenv("APP_SECRET_PATH_"+test.key, test.pathContent)
			os.Setenv("APP_SECRET_VALUE_"+test.key, test.valueContent)
			got, gotErr := Get(test.key, test.handlers...)
			if test.wantErr != nil && gotErr == nil {
				t.Errorf("expected error with message containing: %v, got: none", test.wantErr)
			}
			if test.wantErr != nil && gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr.Error()) {
				t.Errorf("expected error with message containing: %#v, got: %#v", test.wantErr, gotErr)
			}
			if test.want != got {
				t.Errorf("expected %#v, got: %#v", test.want, got)
			}
		})
	}
}
