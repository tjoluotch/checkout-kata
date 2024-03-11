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
	//Quantity     int
}

type Offer map[int]int64

type Store []ItemDetails

type PriceList map[string]int64

type CheckoutIntf interface {
	ScanItem(sku string) error
	GetTotal() int64
}

type Shopper struct {
	Store Store
	mu    sync.Mutex
	PList PriceList
}

func NewShopper() CheckoutIntf {
	return &Shopper{
		Store: Store{},
		mu:    sync.Mutex{},
		PList: getPriceList(),
	}
}

func (s *Shopper) ScanItem(sku string) error {
	// check sku is in priceList
	var found string
	for key, _ := range s.PList {
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
	price := s.PList[found]

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
