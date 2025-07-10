package controller

import (
	"bytes"
	"context"
	"crud-api-go/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// helper function: creates a context for testing
func createTestContext(method, path string, body []byte) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	return c, w
}

// Test successful signup
func TestSignup_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	input := model.SignupInput{
		Name:     "Test User",
		Email:    "testsignup@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(input)

	c, w := createTestContext("POST", "/signup", body)

	// Cleanup user if it already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = getUserCollection().FindOneAndDelete(ctx, bson.M{"email": input.Email})

	Signup(c)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

// Cleanup user if it already exists
func TestSignup_DuplicateEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	input := model.SignupInput{
		Name:     "Test User",
		Email:    "testsignup@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(input)

	c, w := createTestContext("POST", "/signup", body)

	// This time, signup should fail
	Signup(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}
