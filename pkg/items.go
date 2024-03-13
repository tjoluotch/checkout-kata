package pkg

import (
	"errors"
	"fmt"
	"sync"
)

type ItemDetails struct {
	SKU          string
	UnitPrice    int64
	SpecialOffer Offer
}

type Offer map[string]string

type Store []ItemDetails

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

	// get unit price of sku
	price := s.PEngine.PList[found]

	s.mu.Lock()
	s.Store = append(s.Store, ItemDetails{SKU: found, UnitPrice: price})
	s.mu.Unlock()

	return nil
}

func (s *Shopper) GetTotal() int64 {
	total := int64(0)
	// if store has no items
	if len(s.Store) < 1 {
		return 0
	}

	for _, item := range s.Store {
		total += item.UnitPrice
	}
	return total
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
