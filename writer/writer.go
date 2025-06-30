package writer

import (
	"encoding/json"
	"fmt"
	"mantas9/listings/models"
	"os"

	"github.com/gocarina/gocsv"
)

// Marshals Listing data to JSON and writes output to file
func WriteJSON(data []models.Listing) error {
	json, err := json.Marshal(data) // Marshal JSON

	if err != nil { // Error check
		return err
	}

	// Write to file
	if err := os.WriteFile("listings.json", json, 0644); err != nil {
		return err
	}

	// Print success message
	fmt.Println("Your selected NFT Listings' data has been written to listings.json successfully.")

	return nil
}

// Write struct data to CSV file
func WriteCSV(data []models.Listing) error {
	// Create CSV file if exists
	file, err := os.Create("listings.csv")
	if err != nil {
		return err
	}
	defer file.Close() // Close file at the end of writing

	// Write to CSV
	if err := gocsv.MarshalFile(&data, file); err != nil {
		return err
	}

	fmt.Println("Your selected NFT Listings' data has been written to listings.csv successfully.")

	return nil
}
