package tests

import (
	"checkout-kata/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ()

func TestCheckout(t *testing.T) {

	t.Run("scan Item", func(t *testing.T) {
		specialOffer, priceList := pkg.GetDefaultPriceOffers()
		checkout := pkg.NewShopper(pkg.NewPriceEngine(specialOffer, priceList))
		_ = checkout.ScanItem("A")

		sh := getConcreteTypeShopper(t, checkout)

		if len(sh.Store) != 1 {
			t.Error("failed to scan item")
		}
	})

	t.Run("scan item that isn't in price list, error expected", func(t *testing.T) {
		scannedItem := "A33"
		specialOffer, priceList := pkg.GetDefaultPriceOffers()
		checkout := pkg.NewShopper(pkg.NewPriceEngine(specialOffer, priceList))
		err := checkout.ScanItem(scannedItem)

		// if error returned is nil
		if !assert.NotNil(t, err) {
			t.Errorf("expected error as SKU-%s isn't in price list", scannedItem)
		}
	})

	t.Run("scan item and get correct total", func(t *testing.T) {
		specialOffer, priceList := pkg.GetDefaultPriceOffers()
		checkout := pkg.NewShopper(pkg.NewPriceEngine(specialOffer, priceList))
		_ = checkout.ScanItem("B")

		sh := getConcreteTypeShopper(t, checkout)

		if len(sh.Store) != 1 {
			t.Error("failed to scan item")
		}

		expect := int64(30)
		total := checkout.GetTotal()
		if total != expect {
			t.Errorf("failed to get the expected value, wanted %v gotten %v\n", expect, total)
		}
	})

	t.Run("scan 1 of A..D items and get correct total", func(t *testing.T) {
		specialOffer, priceList := pkg.GetDefaultPriceOffers()
		checkout := pkg.NewShopper(pkg.NewPriceEngine(specialOffer, priceList))
		_ = checkout.ScanItem("A")
		_ = checkout.ScanItem("B")
		_ = checkout.ScanItem("C")
		_ = checkout.ScanItem("D")

		expect := int64(115)
		total := checkout.GetTotal()
		if total != expect {
			t.Errorf("failed to get the expected value, wanted %v gotten %v\n", expect, total)
		}
	})

	t.Run("scan B, A, B items and get correct total", func(t *testing.T) {
		specialOffer, priceList := pkg.GetDefaultPriceOffers()
		checkout := pkg.NewShopper(pkg.NewPriceEngine(specialOffer, priceList))
		_ = checkout.ScanItem("B")
		_ = checkout.ScanItem("A")
		_ = checkout.ScanItem("B")

		expect := int64(95)
		total := checkout.GetTotal()
		if total != expect {
			t.Errorf("failed to get the expected value, wanted %v gotten %v\n", expect, total)
		}
	})

	t.Run("scan 8A's, 1D, and 5B's in a random order and get correct total", func(t *testing.T) {
		specialOffer, priceList := pkg.GetDefaultPriceOffers()
		checkout := pkg.NewShopper(pkg.NewPriceEngine(specialOffer, priceList))
		// test case demonstrates correct amount of offers applied with remaining items calculated at unit price
		_ = checkout.ScanItem("B")
		_ = checkout.ScanItem("A")
		_ = checkout.ScanItem("A")
		_ = checkout.ScanItem("D")
		_ = checkout.ScanItem("A")
		_ = checkout.ScanItem("B")
		_ = checkout.ScanItem("B")
		_ = checkout.ScanItem("A")
		_ = checkout.ScanItem("B")
		_ = checkout.ScanItem("A")
		_ = checkout.ScanItem("A")
		_ = checkout.ScanItem("A")
		_ = checkout.ScanItem("B")
		_ = checkout.ScanItem("A")

		expect := int64(495)
		total := checkout.GetTotal()
		if total != expect {
			t.Errorf("failed to get the expected value, wanted %v gotten %v\n", expect, total)
		}
	})

	t.Run("no items scanned, total should be 0, store should be empty", func(t *testing.T) {
		specialOffer, priceList := pkg.GetDefaultPriceOffers()
		checkout := pkg.NewShopper(pkg.NewPriceEngine(specialOffer, priceList))

		sh := getConcreteTypeShopper(t, checkout)

		if len(sh.Store) != 0 {
			t.Error("failed, store should not have any items")
		}

		expect := int64(0)
		total := checkout.GetTotal()
		if total != expect {
			t.Errorf("failed to get the expected value, wanted %v gotten %v\n", expect, total)
		}
	})

}

// getConcreteTypeShopper gets the concrete type *shopper of the checkout interface. This is a test helper
func getConcreteTypeShopper(t testing.TB, checkout pkg.CheckoutIntf) (concrete *pkg.Shopper) {
	t.Helper()
	concrete, ok := checkout.(*pkg.Shopper)
	if !ok {
		t.Error("failed to get concrete type of checkout")
		return nil
	}
	return
}
