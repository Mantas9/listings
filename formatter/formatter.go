package formatter

import (
	"encoding/json"
	"mantas9/listings/models"
)

// Unmarshals JSON data to struct
func UnmarshalJSON(input []byte) ([]models.Listing, error) {
	jsonStruct := []models.ListingJSON{} // Json struct for seamless unmarshalling

	// Unmarshal into jsonStruct
	err := json.Unmarshal(input, &jsonStruct)

	if err != nil { // Error check
		return []models.Listing{}, err
	}

	// Convert to CSV-compatible struct
	res := []models.Listing{} // Result

	// Iterate through each listing
	for i := range jsonStruct {
		// Append converted jsonStruct value to result
		res = append(res, models.Listing{Collection: jsonStruct[i].TokenData.Collection, Seller: jsonStruct[i].Seller, Price: jsonStruct[i].Price, Mint: jsonStruct[i].TokenData.Mint})
	}

	return res, nil
}
