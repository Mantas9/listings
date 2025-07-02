package constants

import (
	"fmt"
	"os"
)

const helpMessage = `Usage: ./listings <parameters> <collection1> <collection2> ... <collectionX>` + "\nPossible parameters:\n\t--limit <integer>\tSets a limit to the amount of listings to fetch for each collection\n\t--min-price <number>\tFilters listings with a minimum price\n\t--max-price <number>\tFilters listings with a maximum price\n\t--desc\t\t\tSort by price in Descending order (default - by price in Ascending order)\n\t--json\t Export data in JSON format"

// Prints a help message to terminal and exits the application
func HelpMessage() {
	fmt.Println(helpMessage)
	os.Exit(0)
}
