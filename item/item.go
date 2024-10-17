package item

// Item represents a scanned item
type Item struct {
	Product Product
}

// Product represents a product in the catalog
type Product struct {
	SKU   string
	Name  string
	Price float64
}

// Initialize the product catalog
var ProductCatalog = map[string]Product{
	"ipd": {SKU: "ipd", Name: "Super iPad", Price: 549.99},
	"mbp": {SKU: "mbp", Name: "MacBook Pro", Price: 1399.99},
	"atv": {SKU: "atv", Name: "Apple TV", Price: 109.50},
	"vga": {SKU: "vga", Name: "VGA adapter", Price: 30.00},
}
