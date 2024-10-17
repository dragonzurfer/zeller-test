package discount

import (
	"main/item"
	"sort"
)

// LargestXItemsDiscount applies a percentage discount to the largest X items
type LargestXItemsDiscount struct {
	X          int
	Percentage float64
}

func (r LargestXItemsDiscount) Apply(items []item.Item) float64 {
	if r.X <= 0 || r.Percentage <= 0 {
		return 0
	}

	// Sort items by price descending
	sortedItems := make([]item.Item, len(items))
	copy(sortedItems, items)
	sort.Slice(sortedItems, func(i, j int) bool {
		return sortedItems[i].Product.Price > sortedItems[j].Product.Price
	})

	discount := 0.0
	for i := 0; i < r.X && i < len(sortedItems); i++ {
		discount += sortedItems[i].Product.Price * (r.Percentage / 100)
	}
	return discount
}

// BuyXGetYFree applies a buy X get Y free rule for a specific SKU
type BuyXGetYFree struct {
	SKU string
	X   int
	Y   int
}

func (r BuyXGetYFree) Apply(items []item.Item) float64 {
	count := 0
	for _, item := range items {
		if item.Product.SKU == r.SKU {
			count++
		}
	}
	freeItems := (count / (r.X + r.Y)) * r.Y
	// Assuming the cheapest items are free
	totalDiscount := 0.0
	if freeItems > 0 {
		// Collect prices of the specific SKU
		prices := []float64{}
		for _, item := range items {
			if item.Product.SKU == r.SKU {
				prices = append(prices, item.Product.Price)
			}
		}
		sort.Float64s(prices) // Sort ascending
		for i := 0; i < freeItems && i < len(prices); i++ {
			totalDiscount += prices[i]
		}
	}
	return totalDiscount
}

// BulkDiscount applies a discounted price when a minimum quantity is purchased
type BulkDiscount struct {
	SKU         string
	MinQuantity int
	NewPrice    float64
}

func (r BulkDiscount) Apply(items []item.Item) float64 {
	count := 0
	for _, item := range items {
		if item.Product.SKU == r.SKU {
			count++
		}
	}
	if count >= r.MinQuantity {
		originalTotal := 0.0
		discountedTotal := 0.0
		for _, item := range items {
			if item.Product.SKU == r.SKU {
				originalTotal += item.Product.Price
				discountedTotal += r.NewPrice
			}
		}
		return originalTotal - discountedTotal
	}
	return 0
}
