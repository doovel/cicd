package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		headers          http.Header
		expectedAPIKey   string
		expectedErrorMsg string
	}{
		{
			headers:          http.Header{},
			expectedAPIKey:   "",
			expectedErrorMsg: "no authorization header included",
		},
		{
			headers:          http.Header{"Authorization": []string{"Bearer abc123"}},
			expectedAPIKey:   "",
			expectedErrorMsg: "malformed authorization header",
		},
		{
			headers:          http.Header{"Authorization": []string{"ApiKey"}},
			expectedAPIKey:   "",
			expectedErrorMsg: "malformed authorization header",
		},
		{
			headers:          http.Header{"Authorization": []string{"ApiKey validKey123"}},
			expectedAPIKey:   "validKey123",
			expectedErrorMsg: "",
		},
	}

	for _, tt := range tests {
		apiKey, err := GetAPIKey(tt.headers)

		if err != nil && err.Error() != tt.expectedErrorMsg {
			t.Errorf("unexpected error: got %v, want %v", err, tt.expectedErrorMsg)
		}

		if apiKey != tt.expectedAPIKey {
			t.Errorf("unexpected API key: got %v, want %v", apiKey, tt.expectedAPIKey)
		}
	}
}
