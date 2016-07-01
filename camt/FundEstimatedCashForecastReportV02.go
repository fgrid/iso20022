package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04000102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.040.001.02 Document"`
	Message *FundEstimatedCashForecastReportV02 `xml:"camt.040.001.02"`
}

func (d *Document04000102) AddMessage() *FundEstimatedCashForecastReportV02 {
	d.Message = new(FundEstimatedCashForecastReportV02)
	return d.Message
}

// Scope
// The FundEstimatedCashForecastReport message is sent by a report provider, such as a transfer agent or a designated agent of the fund, to a report user, such as an investment manager, a fund accountant or any other interested party.
// This message is used to report the estimated cash incomings and outgoings of one or more investment funds, on one or more trade dates. These cash movements may result from, for example, redemption, subscription, switch transactions or dividends.
// Usage
// The FundEstimatedCashForecastReport message is used to report estimated cash movements, i.e., it is sent prior to the cutoff time and/or the price valuation of the fund.
type FundEstimatedCashForecastReportV02 struct {

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information related to the estimated cash-in and cash-out flows for a specific trade date as a result of investment fund transactions, for example,  subscriptions, redemptions or switch to/from a specified investment fund.
	EstimatedFundCashForecastDetails []*iso20022.EstimatedFundCashForecast1 `xml:"EstmtdFndCshFcstDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block. 
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (f *FundEstimatedCashForecastReportV02) AddPoolReference() *iso20022.AdditionalReference3 {
	f.PoolReference = new(iso20022.AdditionalReference3)
	return f.PoolReference
}

func (f *FundEstimatedCashForecastReportV02) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	f.PreviousReference = append(f.PreviousReference, newValue)
	return newValue
}

func (f *FundEstimatedCashForecastReportV02) AddRelatedReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundEstimatedCashForecastReportV02) AddEstimatedFundCashForecastDetails() *iso20022.EstimatedFundCashForecast1 {
	newValue := new (iso20022.EstimatedFundCashForecast1)
	f.EstimatedFundCashForecastDetails = append(f.EstimatedFundCashForecastDetails, newValue)
	return newValue
}

func (f *FundEstimatedCashForecastReportV02) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}

