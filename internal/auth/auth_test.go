package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {

	wantKey := "myKey"
	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("ApiKey %s", wantKey))

	gotKey, err := GetAPIKey(header)
	if err != nil {
		t.Fatal(err)
	}
	if gotKey != wantKey {
		t.Errorf("expected api key to be %s, got %s", wantKey, gotKey)
	}

	header.Del("Authorization")

	_, err = GetAPIKey(header)
	if err == nil {
		t.Error("expected error due to missing authorization header, instead got none")
	}
}
