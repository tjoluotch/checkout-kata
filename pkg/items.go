package pkg

import "sync"

type ItemDetails struct {
	SKU          string
	UnitPrice    int64
	SpecialOffer Offer
	Quantity     int
}

type Offer map[int]int64

type Store []ItemDetails

type PriceList map[string]int64

type CheckoutIntf interface {
	ScanItem(sku string)
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

func (s *Shopper) ScanItem(sku string) {
	s.mu.Lock()
	s.Store = append(s.Store, ItemDetails{SKU: sku})
	s.mu.Unlock()
}

func (s *Shopper) GetTotal() int64 {
	return 0
}

func getPriceList() PriceList {
	return map[string]int64{
		"A": 50,
		"B": 30,
		"C": 20,
		"D": 15,
	}
}
