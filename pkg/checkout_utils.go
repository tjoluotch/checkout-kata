package pkg

import (
	"strconv"
	"strings"
)

// decodeSpecialOffer decodes the offer string given [intxint] format and return the quantity of the offer and the offer price
func decodeSpecialOffer(offer string) (quantity, offerPrice int) {
	results := strings.Split(offer, "x")
	quantity, _ = strconv.Atoi(results[0])
	offerPrice, _ = strconv.Atoi(results[1])
	return
}

// getPriceList gets the default price List
func getPriceList() PriceList {
	return map[string]int64{
		"A": 50,
		"B": 30,
		"C": 20,
		"D": 15,
	}
}

// getSpecialOffers returns the default special offers
func getSpecialOffers() Offer {
	return map[string]string{
		"A": "3x130",
		"B": "2x45",
	}
}

// GetDefaultPriceOffers returns the default special offers and items given unit prices.
func GetDefaultPriceOffers() (Offer, PriceList) {
	return getSpecialOffers(), getPriceList()
}
