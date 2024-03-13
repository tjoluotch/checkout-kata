package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Offer map[string]string

type Store map[string]int

type PriceList map[string]int64

type PriceEngine struct {
	PList        PriceList
	SpecialOffer Offer
}

type CheckoutIntf interface {
	ScanItem(sku string) error
	GetTotal() int64
}

type Shopper struct {
	Store   Store
	mu      sync.Mutex
	PEngine PriceEngine
}

func NewPriceEngine(specialOffer Offer, priceList PriceList) PriceEngine {
	return PriceEngine{
		PList:        priceList,
		SpecialOffer: specialOffer,
	}
}

func NewShopper(pe PriceEngine) CheckoutIntf {

	return &Shopper{
		Store:   Store{},
		mu:      sync.Mutex{},
		PEngine: pe,
	}
}

func (s *Shopper) ScanItem(sku string) error {
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

func decodeSpecialOffer(offer string) (quantity, offerPrice int) {
	results := strings.Split(offer, "x")
	quantity, _ = strconv.Atoi(results[0])
	offerPrice, _ = strconv.Atoi(results[1])
	return
}

func getPriceList() PriceList {
	return map[string]int64{
		"A": 50,
		"B": 30,
		"C": 20,
		"D": 15,
	}
}

func getSpecialOffers() Offer {
	return map[string]string{
		"A": "3x130",
		"B": "2x45",
	}
}

func GetDefaultPriceOffers() (Offer, PriceList) {
	return getSpecialOffers(), getPriceList()
}
