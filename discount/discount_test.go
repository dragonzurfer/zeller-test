package discount_test

import (
	"main/discount"
	"main/item"
	"testing"
)

// Helper function to initialize the product catalog for tests
func initializeProductCatalog() map[string]item.Product {
	return map[string]item.Product{
		"ipd": {SKU: "ipd", Name: "Super iPad", Price: 549.99},
		"mbp": {SKU: "mbp", Name: "MacBook Pro", Price: 1399.99},
		"atv": {SKU: "atv", Name: "Apple TV", Price: 109.50},
		"vga": {SKU: "vga", Name: "VGA adapter", Price: 30.00},
	}
}

func TestLargestXItemsDiscount(t *testing.T) {
	productCatalog := initializeProductCatalog()

	items := []item.Item{
		{Product: productCatalog["ipd"]},
		{Product: productCatalog["mbp"]},
		{Product: productCatalog["atv"]},
		{Product: productCatalog["vga"]},
	}

	// Create LargestXItemsDiscount rule (30% off largest 2 items)
	rule := discount.LargestXItemsDiscount{
		X:          2,
		Percentage: 30.0,
	}

	discount := rule.Apply(items)

	expectedDiscount := (1399.99 * 0.30) + (549.99 * 0.30) // 30% off MacBook Pro and iPad

	if discount != expectedDiscount {
		t.Errorf("expected discount %.2f but got %.2f", expectedDiscount, discount)
	}
}

func TestBuyXGetYFree(t *testing.T) {
	productCatalog := initializeProductCatalog()

	// Buy 3 Apple TVs, 1 should be free
	items := []item.Item{
		{Product: productCatalog["atv"]},
		{Product: productCatalog["atv"]},
		{Product: productCatalog["atv"]},
	}

	// Create BuyXGetYFree rule (Buy 2 get 1 free)
	rule := discount.BuyXGetYFree{
		SKU: "atv",
		X:   2,
		Y:   1,
	}

	discount := rule.Apply(items)
	expectedDiscount := 109.50 // 1 Apple TV free

	if discount != expectedDiscount {
		t.Errorf("expected discount %.2f but got %.2f", expectedDiscount, discount)
	}
}

func TestBuyXGetYFree_NotEnoughItems(t *testing.T) {
	productCatalog := initializeProductCatalog()

	// Buy 2 Apple TVs, should not get any free
	items := []item.Item{
		{Product: productCatalog["atv"]},
		{Product: productCatalog["atv"]},
	}

	// Create BuyXGetYFree rule (Buy 2 get 1 free)
	rule := discount.BuyXGetYFree{
		SKU: "atv",
		X:   2,
		Y:   1,
	}

	discount := rule.Apply(items)
	expectedDiscount := 0.0 // No discount since we didn't buy enough for the rule

	if discount != expectedDiscount {
		t.Errorf("expected discount %.2f but got %.2f", expectedDiscount, discount)
	}
}

func TestBulkDiscount(t *testing.T) {
	productCatalog := initializeProductCatalog()

	// Buy 5 Super iPads, bulk discount should apply
	items := []item.Item{
		{Product: productCatalog["ipd"]},
		{Product: productCatalog["ipd"]},
		{Product: productCatalog["ipd"]},
		{Product: productCatalog["ipd"]},
		{Product: productCatalog["ipd"]},
	}

	// Create BulkDiscount rule (Super iPads at $499.99 if buying 4 or more)
	rule := discount.BulkDiscount{
		SKU:         "ipd",
		MinQuantity: 4,
		NewPrice:    499.99,
	}

	discount := rule.Apply(items)
	originalTotal := 5 * 549.99
	expectedDiscount := originalTotal - (5 * 499.99)

	if discount != expectedDiscount {
		t.Errorf("expected discount %.2f but got %.2f", expectedDiscount, discount)
	}
}

func TestBulkDiscount_NotEnoughItems(t *testing.T) {
	productCatalog := initializeProductCatalog()

	// Buy 3 Super iPads, bulk discount should not apply
	items := []item.Item{
		{Product: productCatalog["ipd"]},
		{Product: productCatalog["ipd"]},
		{Product: productCatalog["ipd"]},
	}

	// Create BulkDiscount rule (Super iPads at $499.99 if buying 4 or more)
	rule := discount.BulkDiscount{
		SKU:         "ipd",
		MinQuantity: 4,
		NewPrice:    499.99,
	}

	discount := rule.Apply(items)
	expectedDiscount := 0.0 // No discount since we didn't meet the minimum quantity

	if discount != expectedDiscount {
		t.Errorf("expected discount %.2f but got %.2f", expectedDiscount, discount)
	}
}
