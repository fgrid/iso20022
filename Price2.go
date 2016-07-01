package iso20022

// Type and information about a price.
type Price2 struct {

	// Specification of the price type.
	Type *YieldedOrValueType1Choice `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountChoice `xml:"Val"`

}


func (p *Price2) AddType() *YieldedOrValueType1Choice {
	p.Type = new(YieldedOrValueType1Choice)
	return p.Type
}

func (p *Price2) AddValue() *PriceRateOrAmountChoice {
	p.Value = new(PriceRateOrAmountChoice)
	return p.Value
}

