package tests

import (
	"checkout-kata/pkg"
	"testing"
)

func TestCheckout(t *testing.T) {

	t.Run("Scan Item", func(t *testing.T) {
		sh := pkg.NewShopper()
		sh.ScanItem("A")

		if len(sh.Store) != 1 {
			t.Error("failed to scan item")
		}
	})

	t.Run("test adding single item to checkout", func(t *testing.T) {
		//s := new(pkg.Shopper)
	})
}
