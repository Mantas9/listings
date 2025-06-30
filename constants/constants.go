package constants

import (
	"fmt"
	"os"
)

const helpMessage = `Usage: ./listings <parameters> <collection1> <collection2> ... <collectionX>` + "\nPossible parameters:\n\t--limit <intager>\tSets a limit to the amount of listings to fetch for each collection\n\t--min-price <number>\tFilters listings with a minimum price\n\t--max-price <number>\tFilters listings with a maximum price\n\t--desc\t\t\tCollection listings are automatically sorted by price ascendingly. This flag sorts the results descendingly"

// Prints a help message to terminal and exits the application
func HelpMessage() {
	fmt.Println(helpMessage)
	os.Exit(0)
}
