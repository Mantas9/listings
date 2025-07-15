package httpfetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// GetListings call parameters
type GetListingsOpts struct {
	Symbol   string  // Collection symbol
	Limit    int64   // Limit of total collection listings to be fetched
	MinPrice float64 // Minimum price of listings
	MaxPrice float64 // Maximum price of listings
	Desc     bool    // Sort results by price descending
}

func GetListings(opts GetListingsOpts) ([]byte, error) { // Base GetListings function call

	// URL To API
	url, err := formListingsURL(opts)

	if err != nil { // Error check
		return nil, err
	}

	// Execute HTTP request
	res, err := httpRequest(url)

	if err != nil { // Error check
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request did not return OK (200)")
	}

	// Close body reader when done
	defer res.Body.Close()
	// Read fetched data
	body, err := io.ReadAll(res.Body)

	// Return result
	return body, err
}

func formListingsURL(opts GetListingsOpts) (string, error) { // Forms the magicEden API URL according to input parameters
	// Handle empty symbol
	if opts.Symbol == "" {
		return "", fmt.Errorf("cannot form URL to API: NFT Symbol is invalid: "+`"`+"%v"+`"`, opts.Symbol)
	}

	url := fmt.Sprintf("https://api-mainnet.magiceden.dev/v2/collections/%s/listings", opts.Symbol)

	// Handle extra parameters
	paramCnt := 0        // Amount of parameters added
	if opts.Limit != 0 { // Listing count limit
		url += fmt.Sprintf("?limit=%d", opts.Limit) // Add limit option
		paramCnt++
	}
	if opts.MinPrice != 0 { // Min price of listing
		// Check if there are already options listed
		if paramCnt != 0 {
			url += "&" // Add & for next param
		} else {
			url += "?" // Add ? because this is the first param
		}

		// Add MinPrice argument
		url += fmt.Sprintf("min_price=%f", opts.MinPrice)

		// Param counter
		paramCnt++
	}
	if opts.MaxPrice != 0 { // Max price of listing
		// Check if there are already options listed
		if paramCnt != 0 {
			url += "&" // Add & for next param
		} else {
			url += "?" // Add ? because this is the first param
		}

		// Add MaxPrice argument
		url += fmt.Sprintf("max_price=%f", opts.MaxPrice)

		// Param counter
		paramCnt++
	}
	if opts.Desc { // Sort descending
		// Check if there are already options listed
		if paramCnt != 0 {
			url += "&" // Add & for next param
		} else {
			url += "?" // Add ? because this is the first param
		}

		// Add Descending order argument
		url += "sort_direction=desc"

		// Param counter
		paramCnt++
	}

	return url, nil
}

func httpRequest(url string) (*http.Response, error) { // Performs a HTTP request and returns the output
	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil { // Error check
		return nil, err
	}

	// Add JSON header to request
	req.Header.Add("accept", "application/json")

	// Http client with timeout handling
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Execute HTTP request
	res, err := client.Do(req)

	if err != nil { // Error check
		return nil, err
	}

	// Success
	return res, nil
}
