package secrets

import "testing"

func TestAWSParameterStoreIntegration(t *testing.T) {
	store := NewAWSParameterStore()
	expected := "test-value-in-parameter-store"
	got, err := store.get("test")
	if err != nil {
		t.Errorf("did not expect error: #%v", err)
	}
	if expected != got {
		t.Errorf("expected %#v, got: %#v", expected, got)
	}
	_, err = store.get("bad")
	if err == nil {
		t.Error("expected error")
	}
}
