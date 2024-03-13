package pkg

import (
	"errors"
	"fmt"
	"unicode"
)

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
