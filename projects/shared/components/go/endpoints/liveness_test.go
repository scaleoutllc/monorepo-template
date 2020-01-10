package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLivenessHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LivenessHandler)
	handler.ServeHTTP(rr, req)
	expectedCode := http.StatusOK
	if status := rr.Code; status != expectedCode {
		t.Errorf("status code: got %v want %v", status, expectedCode)
	}
	expectedBody := http.StatusText(expectedCode)
	if rr.Body.String() != expectedBody {
		t.Errorf("body: got %v want %v", rr.Body.String(), expectedBody)
	}
}
