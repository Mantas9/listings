package models

// ========= Nested Structs (for JSON unmarshaling) ===========
// Structure of the TokenJSON field in an NFT listing
type TokenJSON struct {
	Mint       string `json:"mintAddress"` // NFT mint address
	Collection string `json:"collection"`  // Collection name/symbol
}

// Structure of a single NFT listing
type ListingJSON struct {
	Seller    string    `json:"seller"` // Seller address
	Price     float64   `json:"price"`  // NFT price in SOL
	TokenData TokenJSON `json:"token"`  // Token data
}

// ========= Flat Struct for CSV data ==========
type Listing struct {
	Collection string  `csv:"collection" json:"collection"`
	Seller     string  `csv:"seller" json:"seller"`
	Price      float64 `csv:"price" json:"price"`
	Mint       string  `csv:"mintAddress" json:"mintAddress"`
}
