package main

import (
	httpfetcher "mantas9/listings/httpFetcher"
	"os"
	"strconv"
)

func main() {
	// Get arguments (exclude program call)
	args := os.Args[1:]

	// Handle parameters
	params := httpfetcher.GetListingsOpts{}
	argsToDrop := 0    // Counter for how many arguments to drop from args list after parsing parameters
	valueFlag := false // Flag to parse next value as a parameter argument

	// iterate through each argument and parse its value
	for i, arg := range args {
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
		} else { // If not a parameter and no valueflag, break loop
			break
		}
	}

	// Drop parameter arguments
	args = args[argsToDrop:]
}
