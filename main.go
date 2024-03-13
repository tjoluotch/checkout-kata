package main

import (
	"checkout-kata/pkg"
	"fmt"
)

func main() {
	specialOffer, priceList := pkg.GetDefaultPriceOffers()
	checkout := pkg.NewShopper(pkg.NewPriceEngine(specialOffer, priceList))
	_ = checkout.ScanItem("B")
	_ = checkout.ScanItem("A")
	_ = checkout.ScanItem("B")

	expect := int64(95)
	total := checkout.GetTotal()
	fmt.Printf("gotten the correct total for checkout: %t", expect == total)
}
