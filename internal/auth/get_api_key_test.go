package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://nowhere.mil.gov.sp.it.fu", nil)
	req.Header.Set("Authorization", "ApiKey 123456")
	actual, _ := GetAPIKey(req.Header)
	if actual != "12345" {
		t.Fail()
	}
}
