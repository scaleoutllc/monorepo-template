package secrets

import "testing"

func TestAWSSecretStoreStoreIntegration(t *testing.T) {
	store := NewAWSSecretStore()
	expected := "iixDP55dhqiAREjz2b8ihBhcEYcdyw56"
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
