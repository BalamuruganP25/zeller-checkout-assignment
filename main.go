package main

import (
	"fmt"
	"zeller-checkout-assignment/checkout"
)

func main() {
	products := map[string]checkout.Product{
		"ipd": {SKU: "ipd", Name: "Super iPad", Price: 549.99},
		"mbp": {SKU: "mbp", Name: "MacBook Pro", Price: 1399.99},
		"atv": {SKU: "atv", Name: "Apple TV", Price: 109.50},
		"vga": {SKU: "vga", Name: "VGA adapter", Price: 30.00},
	}

	// Define pricing rules
	pricingRules := []checkout.PricingRule{
		checkout.ThreeForTwoDeal{SKU: "atv"},
		checkout.BulkDiscountRule{SKU: "ipd", MinQuantity: 4, NewPrice: 499.99},
	}

	co := checkout.NewCheckout(pricingRules, products)
	co.Scan("atv")
	co.Scan("atv")
	co.Scan("atv")
	co.Scan("vga")
	fmt.Printf("Total: $%.2f\n", co.Total()) // Expected: $249.00

	co2 := checkout.NewCheckout(pricingRules, products)
	co2.Scan("atv")
	co2.Scan("ipd")
	co2.Scan("ipd")
	co2.Scan("atv")
	co2.Scan("ipd")
	co2.Scan("ipd")
	co2.Scan("ipd")
	fmt.Printf("Total: $%.2f\n", co2.Total()) // Expected: $2718.95
}
