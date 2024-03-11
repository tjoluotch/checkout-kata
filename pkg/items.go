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

type CheckoutIntf interface {
	ScanItem(sku string)
	GetTotal() int64
}

type Shopper struct {
	Store []Store
	mu    sync.Mutex
}

func (s *Shopper) ScanItem(sku string) {

}

func (s *Shopper) GetTotal() int64 {
	return 0
}
