package pkg

import "sync"

// CheckoutIntf is an interface representing the checkout
type CheckoutIntf interface {
	ScanItem(sku string) error
	GetTotal() int64
}

// Store holds all items scanned during a checkout.
type Store map[string]int

// Shopper is an abstraction that implements CheckoutIntf
type Shopper struct {
	Store Store
	// mu is a Mutex for managing state and in a more complex domain preventing race conditions
	mu      sync.Mutex
	PEngine PriceEngine
}

// NewShopper returns an implementation of the Checkout interface with a concrete type *Shopper
func NewShopper(pe PriceEngine) CheckoutIntf {
	return &Shopper{
		Store:   Store{},
		mu:      sync.Mutex{},
		PEngine: pe,
	}
}
