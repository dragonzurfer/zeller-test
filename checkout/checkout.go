package checkout

import (
	"fmt"
	"main/item"
	"main/pricing"
)

// Checkout holds the scanned items and pricing rules
type Checkout struct {
	items        []item.Item
	pricingRules []pricing.PricingRule
}

// NewCheckout initializes a new Checkout with given pricing rules
func NewCheckout(pricingRules []pricing.PricingRule) *Checkout {
	return &Checkout{
		pricingRules: pricingRules,
		items:        []item.Item{},
	}
}

// Scan adds an item to the checkout
func (co *Checkout) Scan(sku string) error {
	product, exists := item.ProductCatalog[sku]
	if !exists {
		return fmt.Errorf("product with SKU '%s' not found", sku)
	}

	co.items = append(co.items, item.Item{Product: product})

	return nil
}

// Total calculates the total price after applying pricing rules
func (co *Checkout) Total() float64 {
	subtotal := 0.0
	for _, item := range co.items {
		subtotal += item.Product.Price
	}

	totalDiscount := 0.0
	for _, rule := range co.pricingRules {
		discount := rule.Apply(co.items)
		totalDiscount += discount
	}

	total := subtotal - totalDiscount
	if total < 0 {
		total = 0
	}
	return total
}
