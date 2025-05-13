package jwt

import (
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestCreateJWToken(t *testing.T) {
	os.Setenv("APP_SECRET", "secret")
	defer os.Unsetenv("APP_SECRET")

	j := JWT{}
	token, err := j.Create(1, "127.0.0.1")
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Error("expected token to be not empty")
	}
}

func TestNewJWT(t *testing.T) {
	os.Setenv("APP_SECRET", "secret")
	defer os.Unsetenv("APP_SECRET")

	j := JWT{}
	token, err := j.Create(1, "127.0.0.1")
	if err != nil {
		t.Fatalf("error creating JWT: %v", err)
	}

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	jwt, err := New(req)
	if err != nil {
		t.Errorf("error creating new JWT: %v", err)
	}

	if jwt == nil {
		t.Error("expected JWT to be not nil")
	}

	if jwt.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", jwt.ID)
	}

	if jwt.ClientIp != "127.0.0.1" {
		t.Errorf("expected ClientIp to be 127.0.0.1, got %s", jwt.ClientIp)
	}
}

func TestNewJWTWithoutBearerToken(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	_, err := New(req)
	if err == nil || !strings.Contains(err.Error(), "no bearer token") {
		t.Error("expected error for missing bearer token")
	}
}

func TestValidateInvalidToken(t *testing.T) {
	os.Setenv("APP_SECRET", "secret")
	defer os.Unsetenv("APP_SECRET")

	invalidToken := "invalid.token.value"
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+invalidToken)

	_, err := New(req)
	if err == nil || !strings.Contains(err.Error(), "invalid token") {
		t.Error("expected error for invalid token")
	}
}
