package services

import (
	"testing"
)

func TestConnection(t *testing.T) {
	result := Connection()
	if result.Server.Port == 0 {
		t.Fatal("Invalid port")
	}
	if len(result.Apikey) == 0 {
		t.Fatal("no Apikey found")
	}
	if len(result.Url) == 0 {
		t.Fatal("no Url found")
	}
}
