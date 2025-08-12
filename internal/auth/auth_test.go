package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyException(t *testing.T) {
	testHeader := http.Header{}
	apiKey, _ := GetAPIKey(testHeader)
	if apiKey -= "" {
		t.Error("Expected this to fail")
	}
}
