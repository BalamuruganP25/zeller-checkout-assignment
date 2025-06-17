package checkout

type PricingRule interface {
	Apply(items map[string]int, products map[string]Product) float64
}

type ThreeForTwoDeal struct {
	SKU string
}

func (r ThreeForTwoDeal) Apply(items map[string]int, products map[string]Product) float64 {
	count := items[r.SKU]
	if count == 0 {
		return 0
	}
	price := products[r.SKU].Price
	setsOfThree := count / 3
	remaining := count % 3
	return float64(setsOfThree*2+remaining) * price
}

type BulkDiscountRule struct {
	SKU         string
	MinQuantity int
	NewPrice    float64
}

func (r BulkDiscountRule) Apply(items map[string]int, products map[string]Product) float64 {
	count := items[r.SKU]
	if count == 0 {
		return 0
	}
	price := products[r.SKU].Price
	if count > r.MinQuantity {
		price = r.NewPrice
	}
	return float64(count) * price
}

type DefaultRule struct{}

func (r DefaultRule) Apply(items map[string]int, products map[string]Product) float64 {
	total := 0.0
	for sku, count := range items {
		total += float64(count) * products[sku].Price
	}
	return total
}
