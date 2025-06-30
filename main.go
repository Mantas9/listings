package main

import (
	"fmt"
	"mantas9/listings/constants"
	"mantas9/listings/formatter"
	httpfetcher "mantas9/listings/httpFetcher"
	"mantas9/listings/models"
	"mantas9/listings/writer"
	"os"
	"strconv"
	"sync"
)

func main() {
	// Get arguments (exclude program call)
	args := os.Args[1:]

	// Handle parameters
	params := httpfetcher.GetListingsOpts{}
	argsToDrop := 0     // Counter for how many arguments to drop from args list after parsing parameters
	valueFlag := false  // Flag to parse next value as a parameter argument
	exportJSON := false // Flag to export to JSON instead of CSV

	// If no arguments were passed, print Help message and exit
	if len(args) <= 0 {
		constants.HelpMessage()
	}

	// iterate through each argument and parse its value
	for i, arg := range args {

		// If help was specified, print Help message
		if arg == "--help" || arg == "-h" {
			constants.HelpMessage()
		}

		if valueFlag { // Skip if value flag and reset
			valueFlag = false
			continue
		}

		if arg == "--limit" && i+1 < len(args) { // Limit param
			// Set value flag
			valueFlag = true

			// Get limit value
			limit, err := strconv.ParseInt(args[i+1], 10, 64)

			if err != nil { // Error check
				panic(err)
			}

			// Set parameter
			params.Limit = limit

			// Add to drop value
			argsToDrop += 2 // param name and value
		} else if arg == "--min-price" && i+1 < len(args) { // Minprice param
			// Set value flag
			valueFlag = true

			// Get minPrice value
			minPrice, err := strconv.ParseFloat(args[i+1], 64)

			if err != nil { // Error check
				panic(err)
			}

			// Set parameter
			params.MinPrice = minPrice

			// Add to drop value
			argsToDrop += 2 // param name and value
		} else if arg == "--max-price" && i+1 < len(args) { // Maxprice param
			// Set value flag
			valueFlag = true

			// Get minPrice value
			maxPrice, err := strconv.ParseFloat(args[i+1], 64)

			if err != nil { // Error check
				panic(err)
			}

			// Set parameter
			params.MaxPrice = maxPrice

			// Add to drop value
			argsToDrop += 2 // param name and value
		} else if arg == "--desc" { // Descending order
			params.Desc = true // Set descending order

			argsToDrop++ // Add one parameter to drop
		} else if arg == "--json" {
			exportJSON = true // Flag export JSON to true

			argsToDrop++ // +1 parameter to drop
		} else if (len(arg) >= 2 && arg[:2] == "--") || arg[0] == '-' { // If invalid parameter, print help message
			constants.HelpMessage()
		} else { // If not a parameter and no valueflag, break loop
			break
		}
	}

	// Drop parameter arguments
	args = args[argsToDrop:]

	var wg sync.WaitGroup             // Waitgroup to prevent code from exiting prematurely
	ch := make(chan []models.Listing) // Channel for concurrent data fetching

	// Start parsing NFT data
	for _, arg := range args {
		wg.Add(1)

		go func(symbol string) {
			defer wg.Done()

			params.Symbol = symbol // Set collection symbol in params

			// Call getListings
			listings, err := getListings(params)

			// Error check
			if err != nil {
				fmt.Printf("Error in fetching stats:\n%s", err)
				os.Exit(1)
			}

			if len(listings) <= 0 { // Warn user about no matches for his collection
				fmt.Printf(`There are no matching Listings for the collection "%v" on the MagicEden Marketplace according to your parameters.`+"\nThis collection will be skipped.\n\n", symbol) // Warn user
			}

			// Push to channel
			ch <- listings

		}(arg)

	}

	// Wait for all tasks to finish
	go func() {
		wg.Wait()
		close(ch) // Close channel
	}()

	var allListings []models.Listing // Combined list of all fetched listings

	// Iterate through channel
	for listing := range ch {
		allListings = append(allListings, listing...) // Append all listings data to main list
	}

	// Export everything in specified format
	if exportJSON { // JSON
		if err := writer.WriteJSON(allListings); err != nil {
			panic(err)
		}
	} else if err := writer.WriteCSV(allListings); err != nil { // Else, CSV
		panic(err)
	}

}

func getListings(options httpfetcher.GetListingsOpts) ([]models.Listing, error) {
	ch := make(chan []byte, 1) // Channel for communicating with http

	// Call httpFetcher to get collection listings
	if err := httpfetcher.GetListings(options, ch); err != nil {
		return nil, err // Error check
	}

	json := <-ch // Fetch channel data

	// Unmarshal JSON
	res, err := formatter.UnmarshalJSON(json)

	// Error check
	if err != nil {
		return nil, err
	}

	return res, nil
}
