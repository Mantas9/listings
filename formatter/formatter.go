package formatter

import (
	"encoding/json"
	"mantas9/listings/models"
	"os"

	"github.com/gocarina/gocsv"
)

// Unmarshals JSON data to struct
func UnmarshalJSON(input string) (models.Listing, error) {
	jsonStruct := models.ListingJSON{} // Json struct for seamless unmarshalling

	// Unmarshal into jsonStruct
	err := json.Unmarshal([]byte(input), &jsonStruct)

	if err != nil { // Error check
		return models.Listing{}, err
	}

	// Convert to CSV-compatible struct
	res := models.Listing{Collection: jsonStruct.TokenData.Collection, Seller: jsonStruct.Seller, Price: jsonStruct.Price, Mint: jsonStruct.TokenData.Mint}

	return res, nil
}

// Marshal struct data to a CSV file
func MarshalCSV(data []models.Listing) error {
	// Create CSV file if exists
	file, err := os.Create("listings.csv")
	if err != nil {
		return err
	}
	defer file.Close() // Close file at the end of writing

	// Marshal CSV
	if err := gocsv.MarshalFile(&data, file); err != nil {
		return err
	}

	return nil
}
