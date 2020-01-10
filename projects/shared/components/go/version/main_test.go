package version

import "testing"

func TestNumber(t *testing.T) {
	expected := "unknown"
	if NUMBER != expected {
		t.Errorf("NUMBER: got %v wanted %v", NUMBER, expected)
	}
}

func TestCommit(t *testing.T) {
	expected := "unknown"
	if COMMIT != expected {
		t.Errorf("COMMIT: got %v wanted %v", COMMIT, expected)
	}
}

func TestIdentifier(t *testing.T) {
	expected := "unknown [commit: unknown]"
	actual := Identifier()
	if actual != expected {
		t.Errorf("COMMIT: got %v wanted %v", actual, expected)
	}
}
