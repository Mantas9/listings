package httpfetcher

import "testing"

// TestFormURL runs formURL with all possible option variations and validates that the API URL was formed correctly
func TestFormURL(t *testing.T) {

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
			url, err := formURL(tt.options) // Run function

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
