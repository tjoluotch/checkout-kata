package pkg

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
