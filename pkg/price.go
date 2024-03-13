package pkg

// PriceList is a mapping of the SKU to its unit price.
type PriceList map[string]int64

// Offer represent a mapping of the SKU and the particular offer given in a string format
// e.g ["A"]"3x150" represents item A having a special offer of 3 for 150.
type Offer map[string]string

// PriceEngine is an abstraction to denote the items for sale and their unit prices and special offers if any.
// This sticks to the brief requirment of pricing being independent of the checkout system.
type PriceEngine struct {
	PList        PriceList
	SpecialOffer Offer
}

// NewPriceEngine returns a pricing engine given a PriceList and Offer thus allowing its implementation to be dynamic.
func NewPriceEngine(specialOffer Offer, priceList PriceList) PriceEngine {
	return PriceEngine{
		PList:        priceList,
		SpecialOffer: specialOffer,
	}
}
