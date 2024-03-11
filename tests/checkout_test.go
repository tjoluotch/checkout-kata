package tests

import (
	"checkout-kata/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ()

func TestCheckout(t *testing.T) {

	t.Run("scan Item", func(t *testing.T) {
		checkout := pkg.NewShopper()
		_ = checkout.ScanItem("A")

		sh := getConcreteTypeShopper(t, checkout)

		if len(sh.Store) != 1 {
			t.Error("failed to scan item")
		}
	})

	t.Run("scan item that isn't in price list, error expected", func(t *testing.T) {
		scannedItem := "A33"
		checkout := pkg.NewShopper()
		err := checkout.ScanItem(scannedItem)

		// if error returned is nil
		if !assert.NotNil(t, err) {
			t.Errorf("expected error as SKU-%s isn't in price list", scannedItem)
		}
	})

	t.Run("scan item and get correct total", func(t *testing.T) {
		checkout := pkg.NewShopper()
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
		checkout := pkg.NewShopper()
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

}

func getConcreteTypeShopper(t testing.TB, checkout pkg.CheckoutIntf) (concrete *pkg.Shopper) {
	t.Helper()
	concrete, ok := checkout.(*pkg.Shopper)
	if !ok {
		t.Error("failed to get concrete type of checkout")
		return nil
	}
	return
}
