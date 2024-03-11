package pkg

type ItemDetails struct {
	Name         string
	UnitPrice    float64
	SpecialOffer Offer
}

type Offer map[int]float64
