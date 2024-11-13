package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGreetingHandler_GetGreeting(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		queryParam     string
		expectedCode   int
		expectedError  bool
		expectedResult string
	}{
		{
			name:           "valid name",
			queryParam:     "Developer",
			expectedCode:   http.StatusOK,
			expectedError:  false,
			expectedResult: "Hello, Developer! Welcome to Go programming!",
		},
		{
			name:           "missing name",
			queryParam:     "",
			expectedCode:   http.StatusBadRequest,
			expectedError:  true,
			expectedResult: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup router and recorder
			router := gin.New()
			handler := NewGreetingHandler()
			router.GET("/greeting", handler.GetGreeting)

			// Create request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/greeting", nil)
			if tt.queryParam != "" {
				q := req.URL.Query()
				q.Add("name", tt.queryParam)
				req.URL.RawQuery = q.Encode()
			}

			// Perform request
			router.ServeHTTP(w, req)

			// Assert status code
			assert.Equal(t, tt.expectedCode, w.Code)

			// Parse response
			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			// Assert response
			if tt.expectedError {
				assert.Contains(t, response, "error")
			} else {
				assert.Equal(t, tt.expectedResult, response["message"])
			}
		})
	}
}
