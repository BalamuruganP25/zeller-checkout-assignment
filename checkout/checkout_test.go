package checkout_test

import (
	"testing"

	"zeller-checkout-assignment/checkout"
)

func testProducts() map[string]checkout.Product {
	return map[string]checkout.Product{
		"ipd": {SKU: "ipd", Name: "Super iPad", Price: 549.99},
		"mbp": {SKU: "mbp", Name: "MacBook Pro", Price: 1399.99},
		"atv": {SKU: "atv", Name: "Apple TV", Price: 109.50},
		"vga": {SKU: "vga", Name: "VGA adapter", Price: 30.00},
	}
}

func testPricingRules() []checkout.PricingRule {
	return []checkout.PricingRule{
		checkout.ThreeForTwoDeal{SKU: "atv"},
		checkout.BulkDiscountRule{SKU: "ipd", MinQuantity: 4, NewPrice: 499.99},
	}
}

func TestCheckout_DefaultRuleProduct(t *testing.T) {
	co := checkout.NewCheckout(testPricingRules(), testProducts())
	co.Scan("mbp") // MacBook Pro has no special pricing rule
	expected := 1399.99
	got := co.Total()
	if got != expected {
		t.Errorf("Default rule: expected %.2f, got %.2f", expected, got)
	}
}

func TestCheckout_ThreeForTwoDeal(t *testing.T) {
	co := checkout.NewCheckout(testPricingRules(), testProducts())
	co.Scan("atv")
	co.Scan("atv")
	co.Scan("atv")
	// Only pay for 2 Apple TVs (3 for 2 deal)
	expected := 2 * 109.50
	got := co.Total()
	if got != expected {
		t.Errorf("ThreeForTwoDeal: expected %.2f, got %.2f", expected, got)
	}
}

func TestCheckout_ThreeForTwoAndBulkDiscount(t *testing.T) {
	co := checkout.NewCheckout(testPricingRules(), testProducts())
	// 3 Apple TVs (3-for-2 deal) and 5 iPads (bulk discount)
	co.Scan("atv")
	co.Scan("atv")
	co.Scan("atv")
	co.Scan("ipd")
	co.Scan("ipd")
	co.Scan("ipd")
	co.Scan("ipd")
	co.Scan("ipd")

	expected := (2 * 109.50) + (5 * 499.99) // 3-for-2 on atv, bulk price for ipd
	got := co.Total()
	if got != expected {
		t.Errorf("ThreeForTwoAndBulkDiscount: expected %.2f, got %.2f", expected, got)
	}
}


func TestCheckout_BulkDiscount(t *testing.T) {
	co := checkout.NewCheckout(testPricingRules(), testProducts())
	// Scan 5 iPads to trigger the bulk discount
	co.Scan("ipd")
	co.Scan("ipd")
	co.Scan("ipd")
	co.Scan("ipd")
	co.Scan("ipd")

	expected := 5 * 499.99 // Bulk discount price
	got := co.Total()
	if got != expected {
		t.Errorf("BulkDiscount: expected %.2f, got %.2f", expected, got)
	}
}

func TestCheckout_InvalidSKU(t *testing.T) {
	co := checkout.NewCheckout(testPricingRules(), testProducts())
	co.Scan("invalid") // SKU does not exist in products map
	expected := 0.0
	got := co.Total()
	if got != expected {
		t.Errorf("Invalid SKU: expected %.2f, got %.2f", expected, got)
	}
}