package utils

import (
	"os"
	"testing"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	testEmail := "test@example.com"

	os.Setenv("JWT_SECRET", "testsecret")
	jwtKey = []byte(os.Getenv("JWT_SECRET"))

	token, err := GenerateJWT(testEmail)
	if err != nil {
		t.Fatalf("Error generating token: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Error validating token: %v", err)
	}

	if claims.Email != testEmail {
		t.Errorf("Incorrect email. Expected %s but got %s", testEmail, claims.Email)
	}
}
