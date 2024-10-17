# interface definition for Pricing rule

```
type PricingRule interface {
    Apply(items []item.Item) float64
}
```

- We can now easily create new discount by simply implementing the Apply function. And at checkout time we can decide which discount to apply.