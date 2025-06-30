package formatter

import (
	"mantas9/listings/models"
	"reflect"
	"testing"
)

// TestUnmarshalJSON calls formatter.UnmarshalJSON with a Valid, Empty and invalid JSON input, checking for a valid return value
func TestUnmarshalJSON(t *testing.T) {
	// Create test table
	var tests = []struct {
		name      string
		input     []byte
		want      []models.Listing
		expectErr bool // Should script be expecting an error?
	}{
		{ // Valid input expects valid output
			name:      "Valid",
			input:     []byte(`[{"pdaAddress":"HdddQqvN3Rri7ViPdXGdMTNAF2P5g7JQLvzVRE6yenbP","auctionHouse":"E8cU1WiRWjanGxmn96ewBgk9vPTcL6AEZ1t6F6fkgUWe","tokenAddress":"588ETYtsMZevJJAJ5YxyLUu9ZvytPB11SondgQczxPxn","tokenMint":"4yeWHnJ11vYYVJ9QWZiCxsiuXA2tQkj15AaczN6Z4hUt","seller":"hero2pbJAsW6iGxkiGBwbptJXeoR4eJuwJQFnp1qa3B","sellerReferral":"autMW8SgBkVYeBgqYiTuJZnkvDZMVU2MHJh9Jh7CSQ2","tokenSize":1,"price":5.3885,"priceInfo":{"solPrice":{"rawAmount":"5388500000","address":"So11111111111111111111111111111111111111112","decimals":9}},"rarity":{"meInstant":{"rank":3936}},"extra":{"img":"https://metadata.degods.com/g/3202-dead-rm.png"},"expiry":-1,"token":{"mintAddress":"4yeWHnJ11vYYVJ9QWZiCxsiuXA2tQkj15AaczN6Z4hUt","owner":"hero2pbJAsW6iGxkiGBwbptJXeoR4eJuwJQFnp1qa3B","supply":1,"collection":"degods","collectionName":"DeGods","name":"DeGod #3203","updateAuthority":"AxFuniPo7RaDgPH6Gizf4GZmLQFc4M5ipckeeZfkrPNn","primarySaleHappened":true,"sellerFeeBasisPoints":333,"image":"https://metadata.degods.com/g/3202-dead-rm.png","animationUrl":"https://animation-url.degods.com?tokenId=3202","attributes":[{"trait_type":"background","value":"Red"},{"trait_type":"skin","value":"Turquoise"},{"trait_type":"specialty","value":"God of War"},{"trait_type":"clothes","value":"Caesar Tunic"},{"trait_type":"neck","value":"None"},{"trait_type":"head","value":"Leaf Laurel"},{"trait_type":"eyes","value":"None"},{"trait_type":"mouth","value":"Hipster Beard"},{"trait_type":"version","value":"S3 - Male"},{"trait_type":"y00t","value":"Claimed"}],"properties":{"files":[{"uri":"https://metadata.degods.com/g/3202-dead-rm.png","type":"image/png"}],"category":"image","creators":[{"address":"AxFuniPo7RaDgPH6Gizf4GZmLQFc4M5ipckeeZfkrPNn","share":100}]},"price":5.3885,"listStatus":"listed","tokenAddress":"588ETYtsMZevJJAJ5YxyLUu9ZvytPB11SondgQczxPxn","priceInfo":{"solPrice":{"rawAmount":"5388500000","address":"So11111111111111111111111111111111111111112","decimals":9}}},"listingSource":"M2"}]`),
			want:      []models.Listing{{Collection: "degods", Seller: "hero2pbJAsW6iGxkiGBwbptJXeoR4eJuwJQFnp1qa3B", Price: 5.3885, Mint: "4yeWHnJ11vYYVJ9QWZiCxsiuXA2tQkj15AaczN6Z4hUt"}},
			expectErr: false,
		},
		{ // Empty input expects an empty array
			name:      "Empty",
			input:     []byte(""),
			want:      []models.Listing{},
			expectErr: true,
		},
		{ // Invalid input expects an error and an empty array
			name:      "Invalid",
			input:     []byte(`[{"pdaAddress":"HdddQqvN3Rri7ViPdXGdMTNAF2P5g7JQLvzVRE6yenbP","auctionHouse":"E8cU1WiRWjanGxmn96ewBgk9vPTcL6AEZ1t6F6fkgUWe","tokenAddress":"588ETYtsMZevJJAmount":"5388500000","address":"So11111111111111111111111111111111111111112","decimals":9}}},"listingSource":"M2"}]`),
			want:      []models.Listing{},
			expectErr: true,
		},
	}

	// Iterate through tests and run them
	for _, tt := range tests {

		testName := tt.name

		// Run test
		t.Run(testName, func(t *testing.T) {
			ans, err := UnmarshalJSON(tt.input) // Unmarshal testing variable

			// Compare answer with wanted data
			if !reflect.DeepEqual(ans, tt.want) { // If tests don't match, throw error
				t.Errorf("Got %v, wanted %v", ans, tt.want)
			}

			// Check for faulty error cases
			if tt.expectErr && err == nil { // Error is expected, but returned as nil
				t.Errorf("Expected error, got nil.")
			}
			if !tt.expectErr && err != nil { // Error is unexpected, but returned as non-nil
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
