package writer

import (
	"encoding/json"
	"fmt"
	"mantas9/listings/models"
	"os"

	"github.com/gocarina/gocsv"
)

// Marshals Listing data to JSON and writes output to file
func WriteJSON(data []models.Listing, filename string) error {
	json, err := json.Marshal(data) // Marshal JSON

	if err != nil { // Error check
		return err
	}

	// Write to file
	if err := os.WriteFile(filename, json, 0644); err != nil {
		return err
	}

	// Print success message
	fmt.Printf("Your selected NFT Listings' data has been written to %s successfully.\n", filename)

	return nil
}

// Write struct data to CSV file
func WriteCSV(data []models.Listing, filename string) error {
	// Create CSV file if exists
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Close file at the end of writing

	// Write to CSV
	if err := gocsv.MarshalFile(&data, file); err != nil {
		return err
	}

	fmt.Printf("Your selected NFT Listings' data has been written to %s successfully.\n", filename)

	return nil
}
