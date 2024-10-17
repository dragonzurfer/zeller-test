package main

import (
	"fmt"
	"main/checkout"
	"main/discount"
	"main/pricing"
)

func main() {
	// Define dynamic pricing rules
	pricingRules := []pricing.PricingRule{
		// Example: Largest 3 items get 30% off
		discount.LargestXItemsDiscount{
			X:          3,
			Percentage: 30.0,
		},
		// Example: Buy 2 Apple TVs get 1 free
		discount.BuyXGetYFree{
			SKU: "atv",
			X:   2,
			Y:   1,
		},
		// Example: Bulk discount on Super iPad when buying 4 or more
		discount.BulkDiscount{
			SKU:         "ipd",
			MinQuantity: 4,
			NewPrice:    499.99,
		},
	}

	// Initialize checkout with pricing rules
	co := checkout.NewCheckout(pricingRules)

	// Scan items
	co.Scan("ipd")
	co.Scan("mbp")
	co.Scan("atv")
	co.Scan("atv")
	co.Scan("atv")
	co.Scan("vga")

	// Calculate total
	total := co.Total()

	fmt.Printf("Total: $%.2f\n", total)
}
