package tests

import (
	"checkout-kata/pkg"
	"testing"
)

func TestCheckout(t *testing.T) {

	t.Run("Scan Item", func(t *testing.T) {
		checkout := pkg.NewShopper()
		checkout.ScanItem("A")

		sh := getConcreteTypeShopper(t, checkout)

		if len(sh.Store) != 1 {
			t.Error("failed to scan item")
		}
	})

	t.Run("test adding single item to checkout", func(t *testing.T) {
		//s := new(pkg.Shopper)
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
