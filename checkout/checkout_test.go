package checkout_test

import (
	"main/checkout"
	"main/discount"
	"main/item"
	"main/pricing"
	"testing"
)

// Helper function to initialize the product catalog for tests
func initializeProductCatalog(t *testing.T) {
	t.Helper()

	item.ProductCatalog = map[string]item.Product{
		"ipd": {SKU: "ipd", Name: "Super iPad", Price: 549.99},
		"mbp": {SKU: "mbp", Name: "MacBook Pro", Price: 1399.99},
		"atv": {SKU: "atv", Name: "Apple TV", Price: 109.50},
		"vga": {SKU: "vga", Name: "VGA adapter", Price: 30.00},
	}
}

func TestCheckout_TotalWithoutRules(t *testing.T) {
	// Initialize product catalog
	initializeProductCatalog(t)

	// Initialize a checkout with no pricing rules
	co := checkout.NewCheckout([]pricing.PricingRule{})

	// Scan items
	err := co.Scan("ipd")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}
	err = co.Scan("mbp")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}
	err = co.Scan("vga")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}

	// Calculate total
	total := co.Total()
	expectedTotal := 549.99 + 1399.99 + 30.00

	if total != expectedTotal {
		t.Errorf("expected total %.2f but got %.2f", expectedTotal, total)
	}
}

func TestCheckout_TotalWithLargestXItemsDiscount(t *testing.T) {
	// Initialize product catalog
	initializeProductCatalog(t)

	// Initialize a checkout with a LargestXItemsDiscount rule
	rules := []pricing.PricingRule{
		discount.LargestXItemsDiscount{
			X:          2, // Largest 2 items get 30% off
			Percentage: 30.0,
		},
	}

	co := checkout.NewCheckout(rules)

	// Scan items
	err := co.Scan("ipd")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}
	err = co.Scan("mbp")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}
	err = co.Scan("vga")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}

	// Calculate total
	total := co.Total()
	expectedSubtotal := 549.99 + 1399.99 + 30.00
	expectedDiscount := (1399.99 * 0.30) + (549.99 * 0.30) // Largest 2 items discounted
	expectedTotal := expectedSubtotal - expectedDiscount

	if total != expectedTotal {
		t.Errorf("expected total %.2f but got %.2f", expectedTotal, total)
	}
}

func TestCheckout_TotalWithBuyXGetYFree(t *testing.T) {
	// Initialize product catalog
	initializeProductCatalog(t)

	// Initialize a checkout with a BuyXGetYFree rule for Apple TVs (Buy 2, Get 1 Free)
	rules := []pricing.PricingRule{
		discount.BuyXGetYFree{
			SKU: "atv", // Buy 2 Apple TVs, get 1 free
			X:   2,
			Y:   1,
		},
	}

	co := checkout.NewCheckout(rules)

	// Scan items
	err := co.Scan("atv")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}
	err = co.Scan("atv")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}
	err = co.Scan("atv")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}

	// Calculate total
	total := co.Total()
	expectedSubtotal := 109.50 * 3 // Price of 3 Apple TVs
	expectedDiscount := 109.50     // One Apple TV free
	expectedTotal := expectedSubtotal - expectedDiscount

	if total != expectedTotal {
		t.Errorf("expected total %.2f but got %.2f", expectedTotal, total)
	}
}

func TestCheckout_TotalWithMultipleRules(t *testing.T) {
	// Initialize product catalog
	initializeProductCatalog(t)

	// Initialize a checkout with multiple rules
	rules := []pricing.PricingRule{
		discount.BuyXGetYFree{
			SKU: "atv", // Buy 2 Apple TVs, get 1 free
			X:   2,
			Y:   1,
		},
		discount.LargestXItemsDiscount{
			X:          2, // Largest 2 items get 30% off
			Percentage: 30.0,
		},
	}

	co := checkout.NewCheckout(rules)

	// Scan items
	err := co.Scan("ipd")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}
	err = co.Scan("mbp")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}
	err = co.Scan("atv")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}
	err = co.Scan("atv")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}
	err = co.Scan("atv")
	if err != nil {
		t.Fatalf("unexpected error while scanning: %v", err)
	}

	// Calculate total
	subtotal := 549.99 + 1399.99 + (109.50 * 3)
	discountForAppleTV := 109.50                                  // One Apple TV free
	discountForLargestItems := (1399.99 * 0.30) + (549.99 * 0.30) // Largest items discounted

	expectedTotal := subtotal - discountForAppleTV - discountForLargestItems
	total := co.Total()

	if total != expectedTotal {
		t.Errorf("expected total %.2f but got %.2f", expectedTotal, total)
	}
}
