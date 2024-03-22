package auth

import (
	"bytes"
	"encoding/json"
	"github.com/patrickkabwe/go-api-starter/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHandler_handleRegister(t *testing.T) {
	testCases := []struct {
		name           string
		payload        types.RegisterUserPayload
		expectedStatus int
	}{
		{
			name: "Should create a new user",
			payload: types.RegisterUserPayload{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "test@gmail.com",
				Password:  "password",
			},
			expectedStatus: 201,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock store
			mockStore := &MockStore{}
			h := NewHandler(mockStore)

			payload, err := json.Marshal(tc.payload)
			if err != nil {
				t.Errorf("Error marshalling payload: %v", err)
			}

			// Create a new request
			req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")

			// Create a new response recorder
			rr := httptest.NewRecorder()

			// Serve the request
			mux := http.NewServeMux()

			mux.HandleFunc("/api/v1/register", h.handleRegister)

			// Register the routes
			mux.ServeHTTP(rr, req)

			// Check the status code
			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, rr.Code)
			}

		})
	}
}

type MockStore struct {
}

func (m *MockStore) CreateUser(user *types.User) error {
	return nil
}

func (m *MockStore) GetUser(id string) (*types.User, error) {
	return nil, nil
}
