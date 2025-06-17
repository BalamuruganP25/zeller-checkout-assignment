package checkout

type Checkout struct {
	pricingRules []PricingRule
	products     map[string]Product
	items        map[string]int
}

func NewCheckout(pricingRules []PricingRule, products map[string]Product) *Checkout {
	return &Checkout{
		pricingRules: pricingRules,
		products:     products,
		items:        make(map[string]int),
	}
}

func (c *Checkout) Scan(sku string) {
	c.items[sku]++
}

// Total calculates the total price of the items in the checkout, applying all applicable pricing rules.
// It processes each rule in the order they were added, applying them to the remaining items.
// If an item is processed by a rule, it is removed from the remaining items for subsequent rules.
// If there are any items left that were not processed by any rule, the default rule is applied to them.
// The total is the sum of all the prices calculated by the rules.
func (c *Checkout) Total() float64 {
	remaining := make(map[string]int)
	for k, v := range c.items {
		remaining[k] = v
	}
	total := 0.0
	for _, rule := range c.pricingRules {
		ruleTotal := rule.Apply(remaining, c.products)
		total += ruleTotal
		// Remove the items that were processed by the rule
		switch r := rule.(type) {
		case ThreeForTwoDeal:
			delete(remaining, r.SKU)
		case BulkDiscountRule:
			delete(remaining, r.SKU)
		}
	}
	// Apply the default rule to any remaining items that were not processed by any specific rule
	if len(remaining) > 0 {
		total += DefaultRule{}.Apply(remaining, c.products)
	}
	return total
}
