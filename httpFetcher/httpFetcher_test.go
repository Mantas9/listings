package httpfetcher

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestHttpRequest runs a moch http request and tests the function's responses
func TestHttpRequest(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping slow test in short mode")
	}

	// Test table
	var tests = []struct {
		name       string
		handler    http.HandlerFunc
		expectErr  bool
		wantStatus int
	}{
		{
			name: "successful request",
			handler: func(w http.ResponseWriter, r *http.Request) {
				// Validate request method (GET)
				if r.Method != http.MethodGet {
					t.Errorf("Expected GET")
				}

				// Check for JSON header
				if r.Header.Get("accept") != "application/json" {
					t.Errorf("Missing Accept header")
				}

				w.WriteHeader(http.StatusOK) // Return 200 if all is good
				w.Write([]byte(`{"message": "ok"}`))
			},
			expectErr:  false,
			wantStatus: http.StatusOK,
		},
		{
			name: "server error",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			expectErr:  false,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "Empty JSON return",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("[]"))
			},
			expectErr:  false,
			wantStatus: http.StatusOK,
		},
		{
			name: "Timeout",
			handler: func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(15 * time.Second)
				fmt.Fprint(w, "I am timed out!")
			},
			expectErr:  true,
			wantStatus: http.StatusGatewayTimeout,
		},
	}

	// Iterate through each scenario
	for _, tt := range tests {
		testName := tt.name

		t.Run(testName, func(t *testing.T) {

			// Mock HTTP server
			server := httptest.NewServer(tt.handler)
			defer server.Close() // Close server at the end of scope

			res, err := httpRequest(server.URL)

			// Error handling scenarios
			if tt.expectErr && err == nil {
				t.Errorf("Expected error, but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}

			// Result status code handling scenarios
			if res != nil && res.StatusCode != tt.wantStatus {
				t.Errorf("Expected status %d, got %d", tt.wantStatus, res.StatusCode)
			}
		})
	}
}

// TestFormListingsURL runs formURL with all possible option variations and validates that the API URL was formed correctly
func TestFormListingsURL(t *testing.T) {

	// Test table
	var tests = []struct {
		name      string          // Test name
		options   GetListingsOpts // Test options for URL generation
		want      string          // Wanted result
		expectErr bool            // Should an error return be expected
	}{
		{
			name: "Default",
			options: GetListingsOpts{
				Symbol: "degods",
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings",
			expectErr: false,
		},
		{
			name:      "Empty",
			options:   GetListingsOpts{},
			want:      "",
			expectErr: true,
		},
		{
			name: "Limit",
			options: GetListingsOpts{
				Symbol: "degods",
				Limit:  5,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?limit=5",
			expectErr: false,
		},
		{
			name: "Limit MinPrice",
			options: GetListingsOpts{
				Symbol:   "degods",
				Limit:    5,
				MinPrice: 2,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?limit=5&min_price=2.000000",
			expectErr: false,
		},
		{
			name: "Limit MaxPrice",
			options: GetListingsOpts{
				Symbol:   "degods",
				Limit:    5,
				MaxPrice: 10,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?limit=5&max_price=10.000000",
			expectErr: false,
		},
		{
			name: "Limit Desc",
			options: GetListingsOpts{
				Symbol: "degods",
				Limit:  5,
				Desc:   true,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?limit=5&sort_direction=desc",
			expectErr: false,
		},
		{
			name: "Limit MinPrice MaxPrice",
			options: GetListingsOpts{
				Symbol:   "degods",
				Limit:    5,
				MinPrice: 2,
				MaxPrice: 10,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?limit=5&min_price=2.000000&max_price=10.000000",
			expectErr: false,
		},
		{
			name: "Limit MinPrice MaxPrice Desc",
			options: GetListingsOpts{
				Symbol:   "degods",
				Limit:    5,
				MinPrice: 2,
				MaxPrice: 10,
				Desc:     true,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?limit=5&min_price=2.000000&max_price=10.000000&sort_direction=desc",
			expectErr: false,
		},
		{
			name: "MinPrice",
			options: GetListingsOpts{
				Symbol:   "degods",
				MinPrice: 2,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?min_price=2.000000",
			expectErr: false,
		},
		{
			name: "MinPrice MaxPrice",
			options: GetListingsOpts{
				Symbol:   "degods",
				MinPrice: 2,
				MaxPrice: 10,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?min_price=2.000000&max_price=10.000000",
			expectErr: false,
		},
		{
			name: "MinPrice Desc",
			options: GetListingsOpts{
				Symbol:   "degods",
				MinPrice: 2,
				Desc:     true,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?min_price=2.000000&sort_direction=desc",
			expectErr: false,
		},
		{
			name: "MinPrice MaxPrice Desc",
			options: GetListingsOpts{
				Symbol:   "degods",
				MinPrice: 2,
				MaxPrice: 10,
				Desc:     true,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?min_price=2.000000&max_price=10.000000&sort_direction=desc",
			expectErr: false,
		},
		{
			name: "MaxPrice",
			options: GetListingsOpts{
				Symbol:   "degods",
				MaxPrice: 10,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?max_price=10.000000",
			expectErr: false,
		},
		{
			name: "MaxPrice Desc",
			options: GetListingsOpts{
				Symbol:   "degods",
				MaxPrice: 10,
				Desc:     true,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?max_price=10.000000&sort_direction=desc",
			expectErr: false,
		},
		{
			name: "Desc",
			options: GetListingsOpts{
				Symbol: "degods",
				Desc:   true,
			},
			want:      "https://api-mainnet.magiceden.dev/v2/collections/degods/listings?sort_direction=desc",
			expectErr: false,
		},
	}

	// Iterate through tests and run them
	for _, tt := range tests {
		testname := tt.name // Declare test name

		// Run test
		t.Run(testname, func(t *testing.T) {
			url, err := formListingsURL(tt.options) // Run function

			// Check for error mismatches
			if tt.expectErr && err == nil {
				t.Error("error mismatch: expected error, but func returned no error.")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			// Validate generated URL
			if url != tt.want {
				t.Errorf("Got: %s\nWanted: %s", url, tt.want)
			}
		})
	}
}
