package pricing

import "main/item"

// PricingRule defines the interface for pricing rules
type PricingRule interface {
	Apply(items []item.Item) float64
}
