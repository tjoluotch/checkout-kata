package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

// Offer represent a mapping of the SKU and the particular offer given in a string format
// e.g ["A"]"3x150" represents item A having a special offer of 3 for 150.
type Offer map[string]string

// Store holds all items scanned during a checkout.
type Store map[string]int

// PriceList is a mapping of the SKU to its unit price.
type PriceList map[string]int64

// PriceEngine is an abstraction to denote the items for sale and their unit prices and special offers if any.
// This sticks to the brief requirment of pricing being independent of the checkout system.
type PriceEngine struct {
	PList        PriceList
	SpecialOffer Offer
}

// CheckoutIntf is an intereface representing the checkout
type CheckoutIntf interface {
	ScanItem(sku string) error
	GetTotal() int64
}

// Shopper is an abstraction that implements CheckoutIntf
type Shopper struct {
	Store Store
	// mu is a Mutex for managing state and in a more complex domain preventing race conditions
	mu      sync.Mutex
	PEngine PriceEngine
}

// NewPriceEngine returns a pricing engine given a PriceList and Offer thus allowing its implementation to be dynamic.
func NewPriceEngine(specialOffer Offer, priceList PriceList) PriceEngine {
	return PriceEngine{
		PList:        priceList,
		SpecialOffer: specialOffer,
	}
}

// NewShopper returns an implementation of the Checkout interface with a concrete type *Shopper
func NewShopper(pe PriceEngine) CheckoutIntf {

	return &Shopper{
		Store:   Store{},
		mu:      sync.Mutex{},
		PEngine: pe,
	}
}

// ScanItem scans an SKU and reports any errors
func (s *Shopper) ScanItem(sku string) error {
	// validate format of scanned item stock keeping unit string passed
	checkStr := []rune(sku)
	for _, val := range checkStr {
		if !unicode.IsUpper(val) {
			return errors.New("incorrect item stock keeping unit, SKU should only be uppercase letter(s)")
		}
	}

	// check sku is in priceList
	var found string
	for key, _ := range s.PEngine.PList {
		if key != sku {
			continue
		} else {
			found = sku
			break
		}
	}
	// error reporting if sku not found in price list
	if found == "" {
		return errors.New(fmt.Sprintf("the given SKU %s is not in the store's price list, try again", sku))
	}

	s.mu.Lock()
	s.Store[found] += 1
	s.mu.Unlock()

	return nil
}

// GetTotal gets the total of the items scanned in the checkout with special offers and unit prices applied.
func (s *Shopper) GetTotal() int64 {
	total := int64(0)
	// if store has no items
	if len(s.Store) < 1 {
		return 0
	}

	total = s.calculateCheckout()
	return total
}

// calculateCheckout calculates the total value of the items in the checkout given both unit prices and special offer
// prices returning back the final amount.
func (s *Shopper) calculateCheckout() int64 {
	final := int64(0)

	for key, amount := range s.Store {
		// singular item, simply add unit price
		if amount == 1 {
			final += s.PEngine.PList[key]
			continue
		}

		// no special offer to check and possibly apply so next iteration
		if s.PEngine.SpecialOffer[key] == "" {
			continue
		}

		quantityOffer, offerPrice := decodeSpecialOffer(s.PEngine.SpecialOffer[key])
		mod := amount % quantityOffer
		if mod == 0 {
			apply := amount / quantityOffer
			temp := int64(apply * offerPrice)
			final += temp
			continue
		} else {
			// remainder after special offer - so apply regular unit price
			partA := int64(mod) * s.PEngine.PList[key]

			// applied offer
			spOfferMultiplier := amount - mod
			spOfferMultiplier = spOfferMultiplier / quantityOffer
			partB := int64(spOfferMultiplier * offerPrice)

			total := partA + partB
			final += total
			continue
		}
	}

	return final
}

// decodeSpecialOffer decodes the offer string given [intxint] format and retursn the quantity of the offer and the offer price
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
