package iso20022

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms18 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Amount following a foreign exchange conversion.
	ConvertedAmount *ActiveCurrencyAndAmount `xml:"ConvtdAmt"`
}

func (f *ForeignExchangeTerms18) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms18) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms18) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms18) SetConvertedAmount(value, currency string) {
	f.ConvertedAmount = NewActiveCurrencyAndAmount(value, currency)
}
