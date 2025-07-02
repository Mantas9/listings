package writer

import (
	"mantas9/listings/models"
	"os"
	"testing"
)

// TestWriteJSON calls WriteJSON inputting Valid, Empty and Partial/Invalid struct arrays
func TestWriteJSON(t *testing.T) {
	var tests = []struct {
		name      string
		filename  string
		input     []models.Listing
		want      string
		expectErr bool
	}{
		{
			name:     "Valid input",
			filename: "valid.json",
			input: []models.Listing{
				{Collection: "degods", Seller: "9taD9QshRxnMzsPcnYpcamu66pwQurfyQ29tbkZVdrS6", Price: 5.2084, Mint: "DNLRw1cWR8nbfgBuxsdiP6KAHMoEFvAVCRgc4PjQbEoY"},
				{Collection: "degods", Seller: "skyi3dryB3a3M5Sh5CTDYput2ET7BjTvXY8SMxZkheA", Price: 5.2094, Mint: "BJh3bS9gxfae6TNuwdorJ7pztVUeXhw2DnP4KVwgfeNV"},
			},
			expectErr: false,
			want:      `[{"collection":"degods","seller":"9taD9QshRxnMzsPcnYpcamu66pwQurfyQ29tbkZVdrS6","price":5.2084,"mintAddress":"DNLRw1cWR8nbfgBuxsdiP6KAHMoEFvAVCRgc4PjQbEoY"},{"collection":"degods","seller":"skyi3dryB3a3M5Sh5CTDYput2ET7BjTvXY8SMxZkheA","price":5.2094,"mintAddress":"BJh3bS9gxfae6TNuwdorJ7pztVUeXhw2DnP4KVwgfeNV"}]`,
		},
		{
			name:      "Empty input",
			filename:  "empty.json",
			input:     []models.Listing{},
			expectErr: false,
			want:      "[]",
		},
		{
			name:     "Invalid input",
			filename: "invalid.json",
			input: []models.Listing{
				{Seller: "9taD9QshRxnMzsPcnYpcamu66pwQurfyQ29tbkZVdrS6"},
				{Price: 5.2094, Mint: "BJh3bS9gxfae6TNuwdorJ7pztVUeXhw2DnP4KVwgfeNV"},
				{},
			},
			expectErr: false,
			want:      `[{"collection":"","seller":"9taD9QshRxnMzsPcnYpcamu66pwQurfyQ29tbkZVdrS6","price":0,"mintAddress":""},{"collection":"","seller":"","price":5.2094,"mintAddress":"BJh3bS9gxfae6TNuwdorJ7pztVUeXhw2DnP4KVwgfeNV"},{"collection":"","seller":"","price":0,"mintAddress":""}]`,
		},
	}

	// Iterate tests and run
	for _, tt := range tests {
		testName := tt.name

		// Run test
		t.Run(testName, func(t *testing.T) {
			err := WriteJSON(tt.input, tt.filename) // Run WriteJSON

			// Error check
			if tt.expectErr && err == nil {
				t.Errorf("Expected error, but method returned no error.")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Expected no error, but got:\n\t%v", err)
			}

			if err == nil { // Check if info exists and read its contents
				info, err := os.Stat(tt.filename)

				if err != nil { // File does not exist
					t.Errorf("File %s should exist, but doesn't:\n\t%v", tt.filename, err)
				} else { // File exists
					// Compare file's content with expected value
					file, err := os.ReadFile(info.Name())

					if err != nil { // Error reading file
						t.Errorf("Reading file %s errored:\n\t%v", info.Name(), err)
					}

					// Compare file contents
					if string(file) != tt.want {
						t.Errorf("Got %v,\nWant %v", string(file), tt.want) // Fail if contents do not match
					}

				}
			}
		})

		// Cleanup test files afterwards
		t.Cleanup(func() {
			os.Remove(tt.filename)
		})
	}
}

// TestWriteCSV calls WriteCSV inputting valid, empty and partial/invalid struct data
func TestWriteCSV(t *testing.T) {

	// Test table
	var tests = []struct {
		name      string
		filename  string
		input     []models.Listing
		want      string
		expectErr bool
	}{
		{
			name:     "Valid input",
			filename: "valid.csv",
			input: []models.Listing{
				{
					Collection: "degods",
					Seller:     "skyi3dryB3a3M5Sh5CTDYput2ET7BjTvXY8SMxZkheA",
					Price:      5.2361,
					Mint:       "BJh3bS9gxfae6TNuwdorJ7pztVUeXhw2DnP4KVwgfeNV",
				},
				{
					Collection: "degods",
					Seller:     "8Gwdguqu9B96eSGFWJbz49PRuKRT5nZNLBDttm4mDQrh",
					Price:      5.2362,
					Mint:       "3TkKMw9BAfd8FQTw352UrbVWKzzBJQFpeMzPGaj2MnVP",
				},
			},
			want:      "collection,seller,price,mintAddress\ndegods,skyi3dryB3a3M5Sh5CTDYput2ET7BjTvXY8SMxZkheA,5.2361,BJh3bS9gxfae6TNuwdorJ7pztVUeXhw2DnP4KVwgfeNV\ndegods,8Gwdguqu9B96eSGFWJbz49PRuKRT5nZNLBDttm4mDQrh,5.2362,3TkKMw9BAfd8FQTw352UrbVWKzzBJQFpeMzPGaj2MnVP\n",
			expectErr: false,
		},
		{
			name:      "Empty input",
			filename:  "empty.csv",
			input:     []models.Listing{},
			want:      "collection,seller,price,mintAddress\n",
			expectErr: false,
		},
		{
			name:     "Invalid input",
			filename: "invalid.csv",
			input: []models.Listing{
				{
					Collection: "degods",
					Mint:       "BJh3bS9gxfae6TNuwdorJ7pztVUeXhw2DnP4KVwgfeNV",
				},
				{
					Seller: "8Gwdguqu9B96eSGFWJbz49PRuKRT5nZNLBDttm4mDQrh",
					Price:  5.2362,
				},
			},
			want:      "collection,seller,price,mintAddress\ndegods,,0,BJh3bS9gxfae6TNuwdorJ7pztVUeXhw2DnP4KVwgfeNV\n,8Gwdguqu9B96eSGFWJbz49PRuKRT5nZNLBDttm4mDQrh,5.2362,\n",
			expectErr: false,
		},
	}

	// Iterate through tests table
	for _, tt := range tests {
		testName := tt.name // Get test name

		// Run test
		t.Run(testName, func(t *testing.T) {
			err := WriteCSV(tt.input, tt.filename) // Run WriteCSV

			// Error check
			if tt.expectErr && err == nil {
				t.Errorf("Expected error, but no error was returned")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if err == nil { // If no error returned, check written file contents
				info, err := os.Stat(tt.filename)

				if err != nil { // File not found
					t.Errorf("Error accessing file %s: %v", tt.filename, err)
				} else { // If exists, compare contents
					file, err := os.ReadFile(info.Name()) // Read contents

					// Error check
					if err != nil {
						t.Errorf("Error reading file %s: %v", info.Name(), err)
					} else if string(file) != tt.want {
						t.Errorf("Got %v, wanted %s", string(file), tt.want)
					}
				}
			}
		})

		// Cleanup test files afterwards
		t.Cleanup(func() {
			os.Remove(tt.filename)
		})
	}
}
